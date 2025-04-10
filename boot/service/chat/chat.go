package chat

import (
	"SYSUCODER/boot/model"
	"SYSUCODER/tools/open_ai"
	"SYSUCODER/tools/prompt"
	"log"
	"strings"
)

// 助手
func Assistant(msg model.ChatMsg) (string, error) {
	msg.Content = strings.TrimSpace(msg.Content)
	if msg.Content == "" {
		return "", nil
	}
	log.Println("请求对话，内容长度:", len(msg.Content))

	// 请求模型
	resp, err := open_ai.Chat(prompt.ChatAssistant, msg.Content)
	if err != nil {
		return "", err
	}
	log.Println("请求成功，结果长度:", len(resp))

	return resp, nil
}
