package utils

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/smtp"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

// 使用邮件通知相关的服务需要利用能够提供对应 API 的邮箱服务，比如网易邮箱
var (
	EmailHost        string
	EmailPort        string
	FromEmail        string
	FromEmailSmtpPwd string
)

var (
	// 缓存中的验证代码在创建后 5 分钟内有效，每个 10 分钟进行一次整理
	verificationCodeCache = cache.New(5*time.Minute, 10*time.Minute)
)

// SendVerificationCode 向用户的邮箱发送验证码
func SendVerificationCode(to string) error {
	// 首先检查验证码是否在 60s 内发送给这个用户
	if timestamp, found := verificationCodeCache.Get(to + "_timestamp"); found {
		if time.Since(timestamp.(time.Time)) < 60*time.Second {
			return fmt.Errorf("请在 60s 后重试")
		}
	}

	code := generateVerificationCode()

	err := sendVerificationCode(to, code)

	if err != nil {
		return err
	}

	// 缓存验证码以及时间戳供后续验证使用
	verificationCodeCache.Set(to, code, cache.DefaultExpiration)
	verificationCodeCache.Set(to+"_timestamp", time.Now(), cache.DefaultExpiration)

	return nil
}

// sendVerificationCode 发送验证码到指定的邮箱
// to: 邮件接收者的邮箱地址
// code: 需要发送的验证码
// error: 发送过程中的错误
func sendVerificationCode(to string, code string) error {
	server := fmt.Sprintf("%s:%s", EmailHost, EmailPort)

	header := make(map[string]string)

	header["From"] = "SYSUCODER" + "<" + FromEmail + ">"
	header["To"] = to
	header["Subject"] = "SYSUCODER 邮箱验证"
	header["Content-Type"] = "text/html;charset=UTF-8"

	body := fmt.Sprintf("[博学 审问 慎思 明辨 笃行]\n\n欢迎使用 SYSUCODER！\n\n您的验证码是： %s", code)

	message := ""

	for k, v := range header {
		message += fmt.Sprintf("%s:%s\r\n", k, v)
	}

	message += "\r\n" + body

	auth := smtp.PlainAuth(
		"",
		FromEmail,
		FromEmailSmtpPwd,
		EmailHost,
	)

	err := SendMailUsingTLS(
		server,
		auth,
		FromEmail,
		[]string{to},
		[]byte(message),
	)

	return err
}

// generateVerificationCode 随机生成一个 6 位验证码
func generateVerificationCode() string {
	// 使用 crypto/rand 生成一个 0 到 999999 之间的随机数
	max := big.NewInt(999999)
	num, err := rand.Int(rand.Reader, max.Add(max, big.NewInt(1)))
	if err != nil {
		// 处理随机数生成失败的情况
		fmt.Println("Error generating random number:", err)
		return "000000" // 返回默认值或采取其他措施
	}

	// 将大整数转换为字符串并格式化为 6 位数
	code := fmt.Sprintf("%06d", num.Int64())
	return code
}

// VerifyVerificationCode 检查验证码是否被正确发送至用户
func VerifyVerificationCode(email string, code string) bool {
	// 从缓存中获取用户邮箱地址对应的验证码
	cachedCode, found := verificationCodeCache.Get(email)
	if !found {
		return false
	}

	// 检查验证码
	if cachedCode != code {
		return false
	}

	// 走到这一步就是通过检查了，删除缓存的验证码
	verificationCodeCache.Delete(email)

	return true
}

// Dial 返回 smtp 客户端
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Panicln("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

// SendMailUsingTLS
// 参考net/smtp的func SendMail()
// 使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
// len(to)>1时,to[1]开始提示是密送
func SendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {
	// smtp 客户端与服务器连接
	c, err := Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = c.Quit()
	if err != nil && !strings.Contains(err.Error(), "250") {
		return err
	}
	return nil
}
