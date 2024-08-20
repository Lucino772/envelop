package cli

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/Lucino772/envelop/internal/wrapper"
	wrapperconf "github.com/Lucino772/envelop/internal/wrapper/conf"
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

	conf, err := wrapperconf.LoadFile(opts.Config)
	if err != nil {
		return err
	}

	conf.Options = append(
		conf.Options,
		wrapper.WithWorkingDirectory(opts.Directory),
		wrapper.WithForwardProcessLogsToLogger(),
	)

	run, err := wrapper.New(conf.Program, conf.Args, conf.Options...)
	if err != nil {
		log.Println("Error while creating wrapper")
		return err
	}
	err = run(context.Background())
	if errors.Is(err, context.Canceled) {
		return nil
	}
	return err
}
