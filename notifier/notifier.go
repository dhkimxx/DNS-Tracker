package notifier

import (
	"fmt"
	"strings"
	"sync"
	"tracker/config"
	implements "tracker/notifier/impl"
)

type Notifier interface {
	SendMessage(message string) error
}

var notifiers []Notifier

func init() {
	notifierType := config.AppConfig.Notifier.NotifierType
	switch notifierType {
	case "slack":
		notifiers = append(notifiers, implements.GetSlackNotifierImpl())
	case "lark":
		notifiers = append(notifiers, implements.GetLarkkNotifierImpl())
	case "both":
		notifiers = append(notifiers, implements.GetSlackNotifierImpl())
		notifiers = append(notifiers, implements.GetLarkkNotifierImpl())
	default:
		panic(fmt.Errorf("unsupported notifier type: [%s]", notifierType))
	}
}

func Notify(a ...any) error {
	message := fmt.Sprint(a...)

	errChan := make(chan error, len(notifiers))
	var wg sync.WaitGroup

	for _, notifier := range notifiers {
		wg.Add(1)
		go func(notifier Notifier) {
			defer wg.Done()
			err := notifier.SendMessage(message)
			if err != nil {
				errChan <- err
			}
		}(notifier)
	}

	wg.Wait()
	close(errChan)

	var errors []string
	for errMsg := range errChan {
		errors = append(errors, errMsg.Error())
	}

	if len(errors) > 0 {
		return fmt.Errorf("%s", strings.Join(errors, "\n"))
	}
	return nil
}

func Notifyf(format string, a ...any) error {
	message := fmt.Sprintf(format, a...)
	err := Notify(message)
	if err != nil {
		return err
	}
	return nil
}
