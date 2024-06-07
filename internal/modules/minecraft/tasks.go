package minecraft

import (
	"context"
	"fmt"
	"regexp"

	"github.com/Lucino772/envelop/internal/wrapper"
)

type checkMinecraftStatusTask struct {
	serverStartingRegex  *regexp.Regexp
	serverPreparingRegex *regexp.Regexp
	serverReadyRegex     *regexp.Regexp
	serverStoppingRegex  *regexp.Regexp
}

func NewCheckMinecraftStatusTask() *checkMinecraftStatusTask {
	return &checkMinecraftStatusTask{
		serverStartingRegex:  regexp.MustCompile(`\[Server thread\/INFO\]\: Starting Minecraft server on \*:(?P<port>[0-9]+)`),
		serverPreparingRegex: regexp.MustCompile(`\[(.*?)\]: Preparing spawn area: (?P<progress>[0-9]+)%`),
		serverReadyRegex:     regexp.MustCompile(`\[Server thread\/INFO\]\: Done \((.*?)\)! For help, type \"help\"`),
		serverStoppingRegex:  regexp.MustCompile(`\[Server thread\/INFO\]\: Stopping server`),
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

func (task *checkMinecraftStatusTask) processValue(wp wrapper.Wrapper, value string) {
	if matches := task.serverStartingRegex.FindStringSubmatch(value); matches != nil {
		wp.GetProcessStatusState().Set(wrapper.ProcessStatusState{
			Description: "Starting",
		})
	} else if matches := task.serverPreparingRegex.FindStringSubmatch(value); matches != nil {
		groups := task.processSubexpNames(task.serverPreparingRegex, matches)
		wp.GetProcessStatusState().Set(wrapper.ProcessStatusState{
			Description: fmt.Sprintf("Preparing (%s%%)", groups["progress"]),
		})
	} else if matches := task.serverReadyRegex.FindStringSubmatch(value); matches != nil {
		wp.GetProcessStatusState().Set(wrapper.ProcessStatusState{
			Description: "Ready",
		})
	} else if matches := task.serverStoppingRegex.FindStringSubmatch(value); matches != nil {
		wp.GetProcessStatusState().Set(wrapper.ProcessStatusState{
			Description: "Stopping",
		})
	}
}

func (task *checkMinecraftStatusTask) Run(ctx context.Context) {
	wp, ok := wrapper.FromIncomingContext(ctx)
	if !ok {
		return
	}

	sub := wp.SubscribeLogs()
	defer sub.Unsubscribe()
	messages := sub.Messages()

	for {
		select {
		case value := <-messages:
			task.processValue(wp, value)
		case <-ctx.Done():
			return
		}
	}
}
