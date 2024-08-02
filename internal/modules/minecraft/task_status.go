package minecraft

import (
	"context"
	"fmt"
	"regexp"

	"github.com/Lucino772/envelop/internal/wrapper"
)

type checkMinecraftStatusTask struct{}

func NewCheckMinecraftStatusTask() *checkMinecraftStatusTask {
	return &checkMinecraftStatusTask{}
}

func (task *checkMinecraftStatusTask) Run(ctx context.Context) error {
	wp, err := wrapper.FromContext(ctx)
	if err != nil {
		return err
	}

	sub := wp.SubscribeLogs()
	defer sub.Close()
	messages := sub.Receive()

	for {
		select {
		case value := <-messages:
			task.processValue(wp, value)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (task *checkMinecraftStatusTask) processSubexpNames(regex *regexp.Regexp, matches []string) map[string]string {
	result := make(map[string]string)
	for ix, name := range regex.SubexpNames() {
		if ix != 0 && name != "" {
			result[name] = matches[ix]
		}
	}
	return result
}

func (task *checkMinecraftStatusTask) processValue(wp wrapper.WrapperContext, value string) {
	if matches := serverStartingRegex.FindStringSubmatch(value); matches != nil {
		wp.PublishState(wrapper.ProcessStatusState{
			Description: "Starting",
		})
	} else if matches := serverPreparingRegex.FindStringSubmatch(value); matches != nil {
		groups := task.processSubexpNames(serverPreparingRegex, matches)
		wp.PublishState(wrapper.ProcessStatusState{
			Description: fmt.Sprintf("Preparing (%s%%)", groups["progress"]),
		})
	} else if matches := serverReadyRegex.FindStringSubmatch(value); matches != nil {
		wp.PublishState(wrapper.ProcessStatusState{
			Description: "Ready",
		})
	} else if matches := serverStoppingRegex.FindStringSubmatch(value); matches != nil {
		wp.PublishState(wrapper.ProcessStatusState{
			Description: "Stopping",
		})
	}
}
