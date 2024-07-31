package cli

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/Lucino772/envelop/internal/modules/core"
	"github.com/Lucino772/envelop/internal/modules/minecraft"
	"github.com/Lucino772/envelop/internal/wrapper"
	"github.com/google/shlex"
	"github.com/spf13/cobra"
)

type wrapperOptions struct {
	Directory string
	Config    string
}

func runCommand() *cobra.Command {
	options := &wrapperOptions{}
	cmd := &cobra.Command{
		Use:   "run [FLAGS] [DIRECTORY]",
		Short: "Run the envelop wrapper",
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				options.Directory = args[0]
			} else {
				directory, err := os.Getwd()
				if err != nil {
					return err
				}
				options.Directory = directory
			}
			absDirectory, err := filepath.Abs(options.Directory)
			if err != nil {
				return err
			}
			options.Directory = absDirectory
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRun(options)
		},
		DisableFlagsInUseLine: true,
	}
	flags := cmd.Flags()
	flags.SetInterspersed(false)
	flags.StringVarP(&options.Config, "config", "", "", "Path to envelop config")
	return cmd
}

func runRun(opts *wrapperOptions) (err error) {
	if opts.Config == "" {
		opts.Config = filepath.Join(opts.Directory, "envelop.yaml")
	}

	conf, err := loadConfig(opts.Config)
	if err != nil {
		return err
	}

	command, err := shlex.Split(conf.Process.Command)
	if err != nil {
		return err
	}

	var options []wrapper.WrapperOptFunc
	options = append(
		options,
		wrapper.WithWorkingDirectory(opts.Directory),
		wrapper.WithGracefulTimeout(time.Duration(conf.Process.Graceful.Timeout)*time.Second),
		wrapper.WithHooks(conf.Hooks),
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

func loadConfig(configPath string) (*wrapper.Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	conf, err := wrapper.LoadConfig(data)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
