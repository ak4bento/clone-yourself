package cmd

import (
	"fmt"
	"github.com/ak4bento/clone-yourself/internal/core"
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Input pengetahuan baru",
	Run: func(cmd *cobra.Command, args []string) {
		topic, _ := cmd.Flags().GetString("topic")
		content, _ := cmd.Flags().GetString("content")
		err := core.SaveKnowledge(topic, content)
		if err != nil {
			fmt.Println("Gagal menyimpan pengetahuan:", err)
			return
		}
		fmt.Printf("[INPUT] Topik '%s' berhasil disimpan!\n", topic)
	},
}

func init() {
	inputCmd.Flags().String("topic", "", "Topik pengetahuan")
	inputCmd.Flags().String("content", "", "Isi pengetahuan")
	rootCmd.AddCommand(inputCmd)
}
