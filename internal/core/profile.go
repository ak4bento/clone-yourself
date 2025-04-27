package core

import (
	"os"
	"gopkg.in/yaml.v3"
)

type AIProfile struct {
  Name       string `yaml:"name"`
  Tone       string `yaml:"tone"` // e.g. "friendly", "formal", "casual", "to-the-point"
	Style      string `yaml:"style"`// e.g. "simple explanation", "uses analogies", "technical"
	Language   string `yaml:"language"`// e.g. "id", "en"
	Signature  string `yaml:"signature"`// optional: penutup atau cara menyapa khas
}

func LoadAIProfile(path string) (*AIProfile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var profile AIProfile
	err = yaml.Unmarshal(data, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
