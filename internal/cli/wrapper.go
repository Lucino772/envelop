package cli

import (
	"context"
	"errors"
	"log"
	"os"
	"path"
	"syscall"
	"time"

	"github.com/Lucino772/envelop/internal/config"
	"github.com/Lucino772/envelop/internal/modules/core"
	"github.com/Lucino772/envelop/internal/modules/minecraft"
	"github.com/Lucino772/envelop/internal/wrapper"
	"github.com/google/shlex"
	"github.com/spf13/cobra"
)

type wrapperOptions struct {
	workingDir string
}

func runWrapperCommand() *cobra.Command {
	options := &wrapperOptions{}
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run the envelop wrapper",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runWrapper(options)
		},
	}
	cmd.Flags().StringVarP(&options.workingDir, "working-dir", "c", "", "Working directory")
	return cmd
}

func runWrapper(opts *wrapperOptions) (err error) {
	if opts.workingDir == "" {
		opts.workingDir, err = os.Getwd()
		if err != nil {
			log.Println("Failed to get working directory")
			return err
		}
	}

	conf, err := loadConfig(path.Join(opts.workingDir, "envelop.yaml"))
	if err != nil {
		log.Println("Failed to load config")
		return err
	}
	command, err := shlex.Split(conf.Process.Command)
	if err != nil {
		log.Println("Failed to parse command")
		return err
	}

	var options []wrapper.WrapperOptFunc
	options = append(
		options,
		wrapper.WithWorkingDirectory(opts.workingDir),
		wrapper.WithGracefulTimeout(time.Duration(conf.Process.Graceful.Timeout)*time.Second),
		wrapper.WithForwardLogToEvent(),
		wrapper.WithForwardLogToStdout(),
	)
	if conf.Process.Graceful.Type == "cmd" {
		options = append(
			options,
			wrapper.WithGracefulStopCommand(conf.Process.Graceful.Options["cmd"].(string)),
		)
	} else if conf.Process.Graceful.Type == "signal" {
		options = append(
			options,
			wrapper.WithGracefulStopSignal(conf.Process.Graceful.Options["signal"].(syscall.Signal)),
		)
	}

	modules := map[string]wrapper.WrapperModule{
		"envelop.core":      core.NewCoreModule(),
		"envelop.minecraft": minecraft.NewMinecraftModule(),
	}
	for _, mod := range conf.Modules {
		if module, ok := modules[mod.Uses]; ok {
			options = append(options, wrapper.WithModule(module))
		} else {
			log.Printf("Failed to load module '%s'\n", mod.Uses)
		}
	}

	wp, err := wrapper.NewWrapper(command[0], command[1:], options...)
	if err != nil {
		log.Println("Error while creating wrapper")
		return err
	}
	err = wp.Run(context.Background())
	if errors.Is(err, context.Canceled) {
		return nil
	}
	return err
}

func loadConfig(configPath string) (*config.Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	conf, err := config.LoadConfig(data)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
