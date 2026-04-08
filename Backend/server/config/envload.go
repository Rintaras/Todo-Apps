package Config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadDotEnv はカレントディレクトリから親へ辿って最初に見つかった .env を読み込みます。
// `go run ./server` を Backend から、バイナリをルートから実行するどちらでも拾えるようにするためのものです。
func LoadDotEnv() (loadedPath string) {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for {
		path := filepath.Join(dir, ".env")
		if st, err := os.Stat(path); err == nil && !st.IsDir() {
			_ = godotenv.Load(path)
			return path
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}
