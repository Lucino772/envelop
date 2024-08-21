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

func (task *checkMinecraftStatusTask) Name() string {
	return "watch-minecraft-status"
}

func (task *checkMinecraftStatusTask) Run(ctx context.Context, wp wrapper.Wrapper) error {
	sub := wp.SubscribeLogs()
	defer sub.Close()

	for value := range sub.Receive() {
		task.processValue(wp, value)
	}
	return nil
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

func (task *checkMinecraftStatusTask) processValue(wp wrapper.Wrapper, value string) {
	if matches := serverStartingRegex.FindStringSubmatch(value); matches != nil {
		wp.UpdateState(wrapper.ProcessStatusState{
			Description: "Starting",
		})
	} else if matches := serverPreparingRegex.FindStringSubmatch(value); matches != nil {
		groups := task.processSubexpNames(serverPreparingRegex, matches)
		wp.UpdateState(wrapper.ProcessStatusState{
			Description: fmt.Sprintf("Preparing (%s%%)", groups["progress"]),
		})
	} else if matches := serverReadyRegex.FindStringSubmatch(value); matches != nil {
		wp.UpdateState(wrapper.ProcessStatusState{
			Description: "Ready",
		})
	} else if matches := serverStoppingRegex.FindStringSubmatch(value); matches != nil {
		wp.UpdateState(wrapper.ProcessStatusState{
			Description: "Stopping",
		})
	}
}
