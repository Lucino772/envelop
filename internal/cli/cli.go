package cli

import "github.com/spf13/cobra"

func RootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "envelop",
	}
	rootCmd.AddCommand(
		runWrapperCommand(),
		installCommand(),
		updateCommand(),
	)
	return rootCmd
}
