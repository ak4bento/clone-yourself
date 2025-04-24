package config

import (
  "fmt"
  "os"
  "path/filepath"
  "github.com/spf13/viper"
)

// Config struct to hold configuration values
type Config struct {
  Port int
  DBPath string
  KnowledgeBasePath string
  LogFilePath string
  ModelPath string
  TempDir string
  TempFilePath string
  TempFileName string
  TempFileExt string
}
