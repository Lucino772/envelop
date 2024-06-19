package cli

import (
	"context"
	"log"
	"os"

	"github.com/Lucino772/envelop/internal/install"
	"github.com/spf13/cobra"
)

type installOptions struct {
	gameId     string
	workingDir string
}

func installCommand() *cobra.Command {
	options := &installOptions{}
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install game server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInstall(options)
		},
	}
	cmd.Flags().StringVarP(&options.workingDir, "working-dir", "c", "", "Working directory")
	cmd.Flags().StringVarP(&options.gameId, "game-id", "g", "", "Game Identifier")
	cmd.MarkFlagRequired("game-id")
	return cmd
}

func runInstall(opts *installOptions) (err error) {
	if opts.workingDir == "" {
		opts.workingDir, err = os.Getwd()
		if err != nil {
			log.Println("Failed to get working directory")
			return err
		}
	}

	manifest, err := install.GetManifest(opts.gameId)
	if err != nil {
		log.Println(err)
		return err
	}
	installer := install.NewInstaller()
	return installer.Install(context.Background(), manifest, opts.workingDir)
}
