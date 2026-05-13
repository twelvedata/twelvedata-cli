package auth

import (
	"os"
	"sync"
)

// ServiceName is the identifier under which keys are stored in the OS keyring.
const ServiceName = "twelvedata-cli"

// CredentialBackend abstracts where the API key material lives. There are two
// implementations: FileBackend (plaintext credentials.json) and KeyringBackend
// (OS-native secure storage).
type CredentialBackend interface {
	Get(account string) (string, error)
	Set(account, secret string) error
	Delete(account string) (bool, error)
	IsAvailable() bool
	Name() string
	IsSecure() bool
}

var (
	cachedBackend CredentialBackend
	backendOnce   sync.Once
)

// GetBackend returns the active backend, picking secure storage when available
// and falling back to file storage. TWELVEDATA_CREDENTIAL_STORE overrides:
// "file" forces file storage; "secure_storage" forces keyring (and falls back
// to file when the keyring backend reports itself unavailable).
func GetBackend() CredentialBackend {
	backendOnce.Do(func() {
		cachedBackend = pickBackend()
	})
	return cachedBackend
}

// ResetBackend clears the cached backend. Useful only in tests.
func ResetBackend() {
	backendOnce = sync.Once{}
	cachedBackend = nil
}

func pickBackend() CredentialBackend {
	override := os.Getenv("TWELVEDATA_CREDENTIAL_STORE")
	if override == "file" {
		return NewFileBackend()
	}
	kr := NewKeyringBackend()
	if override == "secure_storage" {
		if kr.IsAvailable() {
			return kr
		}
		return NewFileBackend()
	}
	if kr.IsAvailable() {
		return kr
	}
	return NewFileBackend()
}
