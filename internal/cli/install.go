package cli

import (
	"context"
	"os"
	"path/filepath"

	"github.com/Lucino772/envelop/internal/install"
	"github.com/spf13/cobra"
)

type installOptions struct {
	Game      string
	Directory string
}

func installCommand() *cobra.Command {
	options := &installOptions{}
	cmd := &cobra.Command{
		Use:   "install GAME [DIRECTORY]",
		Short: "Install game server",
		Args: cobra.MatchAll(
			cobra.MinimumNArgs(1),
			cobra.MaximumNArgs(2),
		),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			options.Game = args[0]
			if len(args) > 1 {
				options.Directory = args[1]
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
			return runInstall(options)
		},
		DisableFlagsInUseLine: true,
	}
	return cmd
}

func runInstall(opts *installOptions) (err error) {
	installer, err := install.NewInstaller()
	if err != nil {
		return err
	}

	if err := installer.CheckManifestsAvailable(); err != nil {
		return err
	}

	manifest, err := installer.GetManifest(opts.Game)
	if err != nil {
		return err
	}
	return installer.Install(
		context.Background(),
		manifest,
		install.DownloadConfig{
			InstallDir: opts.Directory,
		},
	)
}
