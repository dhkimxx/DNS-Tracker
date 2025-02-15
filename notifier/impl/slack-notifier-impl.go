package implements

import (
	"fmt"
	"strings"
	"sync"
	"tracker/config"

	"github.com/slack-go/slack"
)

type SlackNotifierImpl struct {
	api        *slack.Client
	channelIds []string
}

var slackNotifier SlackNotifierImpl

func GetSlackNotifierImpl() *SlackNotifierImpl {
	slackNotifier.api = slack.New(config.AppConfig.Notifier.Slack.Token)
	slackNotifier.channelIds = config.AppConfig.Notifier.Slack.ChannelIds
	return &slackNotifier
}

func (n *SlackNotifierImpl) SendMessage(message string) error {
	errChan := make(chan string, len(n.channelIds))
	var wg sync.WaitGroup

	for _, channelId := range n.channelIds {
		wg.Add(1)
		go func(channelId string) {
			defer wg.Done()
			_, _, err := n.api.PostMessage(channelId, slack.MsgOptionText(message, false))
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
