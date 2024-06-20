package minecraft

import "regexp"

var (
	serverStartingRegex   = regexp.MustCompile(`\[Server thread\/INFO\]\: Starting Minecraft server on \*:(?P<port>[0-9]+)`)
	serverPreparingRegex  = regexp.MustCompile(`\[(.*?)\]: Preparing spawn area: (?P<progress>[0-9]+)%`)
	serverReadyRegex      = regexp.MustCompile(`\[Server thread\/INFO\]\: Done \((.*?)\)! For help, type \"help\"`)
	serverStoppingRegex   = regexp.MustCompile(`\[Server thread\/INFO\]\: Stopping server`)
	serverRegexQueryReady = regexp.MustCompile(`\[Query Listener #1\/INFO\]\: Query running on (.*?):(?P<port>[0-9]+)`)
)
