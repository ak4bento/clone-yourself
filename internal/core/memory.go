package core

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
	"strings"
)

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
			topic TEXT,
			content TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
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

func SaveKnowledge(topic, content string) error {
	_, err := db.Exec(`INSERT INTO knowledge (topic, content) VALUES (?, ?)`, topic, content)
	return err
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

	var best entry

	for rows.Next() {
		var topic, content string
		if err := rows.Scan(&topic, &content); err != nil {
			continue
		}

		score := 0
		for _, word := range keywords {
			if strings.Contains(strings.ToLower(topic), word) || strings.Contains(strings.ToLower(content), word) {
				score++
			}
		}

		if score > best.score {
			best = entry{topic, content, score}
		}
	}

	if best.score == 0 {
		return "Maaf, aku belum punya pengetahuan soal itu.", nil
	}

	return best.content, nil
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
