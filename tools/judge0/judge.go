package judge0

import (
	"SYSUCODER/boot/configuration"
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
)

// InitJudge 初始化与 Judge0 服务器的连接，主要是检测连通性。如果网络连接正常，先通过 Token 设置好登录信息
func InitJudge() error {
	config = configuration.Conf.Judge
	preUrl = config.Host + ":" + config.Port
	log.Println("Connecting to judge server: " + preUrl)
	response, err := About()
	if err != nil || response.StatusCode != http.StatusOK {
		log.Println("------------------------Judge server is not available!------------------------")
		return err
	}

	log.Println("------------------------Judge server is available.------------------------")
	return nil
}

// httpInteraction 是通用的、与 Judge0 服务器交互的方法。通过指定 URL 以及请求方法获取对应的响应，最后返回回送的报文
func httpInteraction(route string, httpMethod string, reader *bytes.Reader) (string, error) {
	url := preUrl + route
	var req *http.Request
	var err error
	if route == "/submissions" && httpMethod == "POST" {
		log.Println("Wait for judge0 server to finish checking...")
		url = url + "?wait=true"
	}
	if reader == nil {
		req, err = http.NewRequest(httpMethod, url, nil)
	} else {
		req, err = http.NewRequest(httpMethod, url, reader)
	}
	if err != nil {
		return "", err
	}

	req.Header.Set("x-rapidapi-key", config.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	bodystr := string(body)
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		return "", errors.New(bodystr)
	}
	return bodystr, nil
}
