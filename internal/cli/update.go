package cli

import (
	"context"

	"github.com/Lucino772/envelop/internal/install"
	"github.com/spf13/cobra"
)

func updateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update the install manifests",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpdate()
		},
		DisableFlagsInUseLine: true,
	}
	return cmd
}

func runUpdate() (err error) {
	installer, err := install.NewInstaller()
	if err != nil {
		return err
	}
	return installer.UpdateManifests(context.Background())
}
