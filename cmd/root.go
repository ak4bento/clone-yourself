package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cyourself",
	Short: "AI clone of Ben",
	Long:  `AI yang dibangun berdasarkan karakter dan pengetahuan Ben.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
