package auth

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestGetConfigDir_Precedence(t *testing.T) {
	t.Run("TWELVEDATA_CONFIG_DIR wins", func(t *testing.T) {
		t.Setenv("TWELVEDATA_CONFIG_DIR", "/tmp/explicit")
		t.Setenv("XDG_CONFIG_HOME", "/tmp/xdg")
		if got := GetConfigDir(); got != "/tmp/explicit" {
			t.Errorf("GetConfigDir() = %q, want /tmp/explicit", got)
		}
	})

	t.Run("XDG_CONFIG_HOME beats default", func(t *testing.T) {
		t.Setenv("TWELVEDATA_CONFIG_DIR", "")
		t.Setenv("XDG_CONFIG_HOME", "/tmp/xdg")
		want := filepath.Join("/tmp/xdg", "twelvedata")
		if got := GetConfigDir(); got != want {
			t.Errorf("GetConfigDir() = %q, want %q", got, want)
		}
	})

	t.Run("Default under home", func(t *testing.T) {
		t.Setenv("TWELVEDATA_CONFIG_DIR", "")
		t.Setenv("XDG_CONFIG_HOME", "")
		got := GetConfigDir()
		if runtime.GOOS != "windows" {
			if !filepath.IsAbs(got) {
				t.Errorf("expected absolute path, got %q", got)
			}
			if filepath.Base(got) != "twelvedata" {
				t.Errorf("default dir should end with twelvedata, got %q", got)
			}
		}
	})
}

func TestCredentialsPath(t *testing.T) {
	t.Setenv("TWELVEDATA_CONFIG_DIR", "/tmp/x")
	want := filepath.Join("/tmp/x", "credentials.json")
	if got := CredentialsPath(); got != want {
		t.Errorf("CredentialsPath() = %q, want %q", got, want)
	}
}
