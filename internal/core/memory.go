package core

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
	"strings"
  "sort"
  "fmt"
)

type KnowledgeEntry struct {
	Category       string
	Topic          string
	Content        string
	Tags           string
	RelevanceScore float64
}

var allowedCategories = map[string]bool{
	"fact": true, "story": true, "opinion": true, "experience": true,
}

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite", "./knowledge/ai4ben.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS knowledge (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      category TEXT DEFAULT 'fact',
      topic TEXT NOT NULL,
      content TEXT NOT NULL,
      tags TEXT DEFAULT '',
      relevance_score REAL DEFAULT 0.7,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  )`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS context_history (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      user_input TEXT NOT NULL,
      ai_response TEXT,
      referenced_knowledge_ids TEXT, -- comma-separated IDs dari table knowledge
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP
  )`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS knowledge_awareness (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      knowledge_id INTEGER NOT NULL,
      understanding_level TEXT CHECK(understanding_level IN ('unread', 'skimmed', 'read', 'deep', 'uncertain')) DEFAULT 'unread',
      notes TEXT,
      last_read DATETIME,
      FOREIGN KEY (knowledge_id) REFERENCES knowledge(id)
  )`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS interactions (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      question TEXT,
      answer TEXT,
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`)

	if err != nil {
		log.Fatal(err)
	}
}

func SaveKnowledge(entry KnowledgeEntry) error {
	// 1. Validasi dan Set default
	if entry.Category == "" || !allowedCategories[entry.Category] {
		entry.Category = "fact"
	}

	if entry.Tags == "" {
		entry.Tags = strings.ToLower(strings.ReplaceAll(entry.Topic, " ", ","))
	}

	if entry.RelevanceScore == 0 {
		entry.RelevanceScore = 0.7
	}

	// 2. Insert ke database
	query := `INSERT INTO knowledge (category, topic, content, tags, relevance_score) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, entry.Category, entry.Topic, entry.Content, entry.Tags, entry.RelevanceScore)
	if err != nil {
		return fmt.Errorf("gagal insert knowledge: %w", err)
	}
	return nil
}


func FindRelevantKnowledge(question string) (string, error) {
	keywords := ExtractKeywords(question)

	rows, err := db.Query(`SELECT topic, content FROM knowledge`)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	type entry struct {
		topic   string
		content string
		score   int
	}

	var matched []entry

	for rows.Next() {
		var topic, content string
		if err := rows.Scan(&topic, &content); err != nil {
			continue
		}

		score := 0
		for _, word := range keywords {
			word = strings.ToLower(word)
			if strings.Contains(strings.ToLower(topic), word) || strings.Contains(strings.ToLower(content), word) {
				score++
			}
		}

		if score > 0 {
			matched = append(matched, entry{topic, content, score})
		}
	}

	if len(matched) == 0 {
		return "Maaf, aku belum punya pengetahuan soal itu.", nil
	}

	// Urutkan dari skor tertinggi ke terendah
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].score > matched[j].score
	})

	var answer strings.Builder
	answer.WriteString("Berikut beberapa informasi yang relevan:\n\n")

	for _, e := range matched {
		answer.WriteString(fmt.Sprintf("🔹 *%s*:\n%s\n\n", e.topic, e.content))
	}

	return answer.String(), nil
}


func LogInteraction(question, answer string) error {
	_, err := db.Exec(`INSERT INTO interactions (question, answer) VALUES (?, ?)`, question, answer)
	return err
}

func FindSimilarInteraction(question string) (string, string, error) {
	keywords := ExtractKeywords(question)

	rows, err := db.Query(`SELECT question, answer FROM interactions`)
	if err != nil {
		return "", "", err
	}
	defer rows.Close()

	type pair struct {
		q     string
		a     string
		score int
	}

	var best pair

	for rows.Next() {
		var q, a string
		if err := rows.Scan(&q, &a); err != nil {
			continue
		}

		score := 0
		for _, word := range keywords {
			if strings.Contains(strings.ToLower(q), word) {
				score++
			}
		}

		if score > best.score {
			best = pair{q, a, score}
		}
	}

	if best.score == 0 {
		return "", "", nil
	}

	return best.q, best.a, nil
}
