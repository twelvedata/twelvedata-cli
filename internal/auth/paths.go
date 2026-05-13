package auth

import (
	"os"
	"path/filepath"
	"runtime"
)

const credentialsFileName = "credentials.json"

// GetConfigDir returns the directory holding credentials.json. Resolution:
// TWELVEDATA_CONFIG_DIR escape hatch, then XDG_CONFIG_HOME, then %APPDATA% on
// Windows, then ~/.config.
func GetConfigDir() string {
	if v := os.Getenv("TWELVEDATA_CONFIG_DIR"); v != "" {
		return v
	}
	if v := os.Getenv("XDG_CONFIG_HOME"); v != "" {
		return filepath.Join(v, "twelvedata")
	}
	if runtime.GOOS == "windows" {
		if v := os.Getenv("APPDATA"); v != "" {
			return filepath.Join(v, "twelvedata")
		}
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(".", ".config", "twelvedata")
	}
	return filepath.Join(home, ".config", "twelvedata")
}

// CredentialsPath returns the full path to credentials.json.
func CredentialsPath() string {
	return filepath.Join(GetConfigDir(), credentialsFileName)
}
