package storage

import (
	"fmt"
	"path/filepath"
	"time"
)

func GenerateFilename(original string) string {
	ext := filepath.Ext(original)
	name := filepath.Base(original[:len(original)-len(ext)])
	timestamp := time.Now().UnixNano()

	return fmt.Sprintf("%s_%d%s", name, timestamp, ext)
}
