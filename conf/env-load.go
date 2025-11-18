package conf

import (
	"os"
	"path/filepath"
	"strings"
)

func LoadEnv() {
	data, err := os.ReadFile(findEnvPath())
	if err != nil {
		return
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		if os.Getenv(parts[0]) == "" {
			_ = os.Setenv(parts[0], parts[1])
		}
	}
}

func findEnvPath() string {
	dir, _ := os.Getwd()
	for dir != "" && dir != string(filepath.Separator) {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}
		dir = filepath.Dir(dir)
	}
	return ".env"
}
