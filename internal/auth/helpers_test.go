package auth

import (
	"testing"
)

// setupTestDir isolates a single test from any real on-disk credentials. It
// points GetConfigDir at a fresh temp dir, forces the file backend so the OS
// keyring isn't touched, and clears the API-key / profile env vars so the
// resolver behaves deterministically.
func setupTestDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("TWELVEDATA_CONFIG_DIR", dir)
	t.Setenv("TWELVEDATA_CREDENTIAL_STORE", "file")
	t.Setenv("TWELVEDATA_API_KEY", "")
	t.Setenv("TWELVEDATA_PROFILE", "")
	t.Setenv("XDG_CONFIG_HOME", "")
	ResetBackend()
	t.Cleanup(ResetBackend)
	return dir
}
