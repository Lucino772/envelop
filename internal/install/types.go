package install

type DownloadConfig struct {
	TargetOs    string
	TargetArch  string
	TargetLang  string
	InstallDir  string
	LowViolence bool
}

type Waiter interface {
	Done() <-chan struct{}
	Wait() error
}
