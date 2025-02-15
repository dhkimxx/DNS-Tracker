package implements

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"tracker/config"
)

type LarkNotifierImpl struct {
	webhookUrl string
}

type LarkMessage struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

var larkNotifier LarkNotifierImpl

func GetLarkkNotifierImpl() *LarkNotifierImpl {
	larkNotifier.webhookUrl = config.AppConfig.Notifier.Lark.WebhookUrl
	return &larkNotifier
}

func (n *LarkNotifierImpl) SendMessage(message string) error {
	msg := LarkMessage{
		MsgType: "text",
	}
	msg.Content.Text = message

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("[notifier-lark] %v", err)
	}

	resp, err := http.Post(n.webhookUrl, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		return fmt.Errorf("[notifier-lark] %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[notifier-lark] %v", fmt.Errorf("failed to send message, status code: %d", resp.StatusCode))
	}

	return nil
}
