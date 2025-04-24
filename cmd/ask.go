package cmd

import (
	"fmt"
	"github.com/ak4bento/clone-yourself/internal/core"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ajukan pertanyaan ke AI",
	Run: func(cmd *cobra.Command, args []string) {
		question, _ := cmd.Flags().GetString("q")
		fmt.Printf("[ASK] Pertanyaan: %s\n", question)

		answer, err := core.FindRelevantKnowledge(question)
		if err != nil {
			fmt.Println("Gagal mencari jawaban:", err)
			return
		}

		fmt.Println("[JAWABAN]:", answer)

		core.LearnFromInteraction(question, answer)
	},
}

func init() {
	askCmd.Flags().String("q", "", "Pertanyaan untuk AI")
	rootCmd.AddCommand(askCmd)
}
