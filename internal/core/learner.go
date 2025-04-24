package core

import "fmt"

func LearnFromInteraction(input, response string) {
	fmt.Printf("[LEARN] Mempelajari input: '%s' dengan respon: '%s'\n", input, response)
	err := LogInteraction(input, response)
	if err != nil {
		fmt.Println("Gagal menyimpan log interaksi:", err)
	}
}
