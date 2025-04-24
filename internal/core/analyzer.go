package core

import (
	"strings"
)

func ExtractKeywords(text string) []string {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, "?", "")
	words := strings.Fields(text)
	return words
}
func AnalyzeQuestion(question string) string {
	knowledge, err := FindRelevantKnowledge(question)
	if err != nil || knowledge == "" {
		return "❌ Saya belum memiliki pengetahuan tentang itu."
	}

	// Kirim ke LLM
	answer, err := GenerateAnswerWithLLM(knowledge, question)
	if err != nil {
		return "⚠️ Gagal menghubungi LLM: " + err.Error()
	}

	return answer
}
// func AnalyzeQuestion(question string) string {
// 	// Cek interaksi mirip
// 	if prevQ, prevA, _ := FindSimilarInteraction(question); prevA != "" {
// 		return fmt.Sprintf("Pernah ditanyakan: \"%s\"\nJawaban: %s", prevQ, prevA)
// 	}
//
// 	// Kalau belum pernah, cari dari knowledge DB
// 	answer, err := FindRelevantKnowledge(question)
// 	if err != nil {
// 		return "Terjadi error saat mencari jawaban."
// 	}
// 	return answer
// }
