package cmd

import (
	"fmt"
	"github.com/ak4bento/clone-yourself/internal/core"
	"github.com/spf13/cobra"
  "strconv"
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Input pengetahuan baru",
	Run: func(cmd *cobra.Command, args []string) {
		topic, err := cmd.Flags().GetString("topic")
		if err != nil {
			fmt.Println("Gagal membaca flag topic:", err)
			return
		}

		content, err := cmd.Flags().GetString("content")
		if err != nil {
			fmt.Println("Gagal membaca flag content:", err)
			return
		}

		category, _ := cmd.Flags().GetString("category") // optional
		tags, _ := cmd.Flags().GetString("tags")         // optional
		relevanceStr, _ := cmd.Flags().GetString("relevance")

		relevance := 0.7
		if relevanceStr != "" {
			r, err := strconv.ParseFloat(relevanceStr, 64)
			if err == nil {
				relevance = r
			}
		}

		entry := core.KnowledgeEntry{
			Category:       category,
			Topic:          topic,
			Content:        content,
			Tags:           tags,
			RelevanceScore: relevance,
		}

		err = core.SaveKnowledge(entry)
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
	inputCmd.Flags().String("category", "", "Kategori (fact/story/opinion/experience)")
	inputCmd.Flags().String("tags", "", "Tags (contoh: golang,concurrency)")
	inputCmd.Flags().String("relevance", "", "Skor relevansi (0.0 - 1.0)")
	inputCmd.MarkFlagRequired("topic")
	inputCmd.MarkFlagRequired("content")
	rootCmd.AddCommand(inputCmd)
}
