package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Input pengetahuan baru",
	Run: func(cmd *cobra.Command, args []string) {
		topic, _ := cmd.Flags().GetString("topic")
		content, _ := cmd.Flags().GetString("content")
		fmt.Printf("[INPUT] Topik: %s\n%s\n", topic, content)
		// TODO: Simpan ke DB
	},
}

func init() {
	inputCmd.Flags().String("topic", "", "Topik pengetahuan")
	inputCmd.Flags().String("content", "", "Isi pengetahuan")
	rootCmd.AddCommand(inputCmd)
}
