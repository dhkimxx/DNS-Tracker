package notifier

import (
	"fmt"
	"strings"
	"sync"
	"tracker/config"

	"github.com/slack-go/slack"
)

type SlackNotifierClient struct {
	isInitialized bool
	api           *slack.Client
	channelIds    []string
}

var client SlackNotifierClient

func init() {
	if config.AppConfig.Notifier.Slack.Enable {
		client.isInitialized = true
		client.api = slack.New(config.AppConfig.Notifier.Slack.Token)
		client.channelIds = config.AppConfig.Notifier.Slack.ChannelIds
	}
}

func Notify(a ...any) error {
	message := fmt.Sprint(a...)
	err := SendMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func Notifyf(format string, a ...any) error {
	message := fmt.Sprintf(format, a...)
	err := SendMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func SendMessage(message string) error {
	if !client.isInitialized {
		return fmt.Errorf("[notifier-slack] slack notifier client is not enabled")
	}

	errChan := make(chan string, len(client.channelIds))
	var wg sync.WaitGroup

	for _, channelId := range client.channelIds {
		wg.Add(1)
		go func(channelId string) {
			defer wg.Done()
			_, _, err := client.api.PostMessage(channelId, slack.MsgOptionText(message, false))
			if err != nil {
				errChan <- fmt.Sprintf("[notifier-slack] channel %s: %v", channelId, err)
			}
		}(channelId)
	}

	wg.Wait()
	close(errChan)

	var errors []string
	for errMsg := range errChan {
		errors = append(errors, errMsg)
	}

	if len(errors) > 0 {
		return fmt.Errorf("%s", strings.Join(errors, "\n"))
	}
	return nil
}
