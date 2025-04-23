package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ajukan pertanyaan ke AI",
	Run: func(cmd *cobra.Command, args []string) {
		question, _ := cmd.Flags().GetString("q")
		fmt.Printf("[ASK] Pertanyaan: %s\n", question)
		// TODO: Query memory dan generate jawaban
	},
}

func init() {
	askCmd.Flags().String("q", "", "Pertanyaan untuk AI")
	rootCmd.AddCommand(askCmd)
}
