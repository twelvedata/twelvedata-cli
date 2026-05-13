package auth

import (
	"errors"

	"github.com/zalando/go-keyring"
)

// KeyringBackend stores API keys in the OS-native secure store (Keychain on
// macOS, Credential Manager on Windows, Secret Service on Linux) via
// github.com/zalando/go-keyring.
type KeyringBackend struct{}

func NewKeyringBackend() *KeyringBackend { return &KeyringBackend{} }

func (KeyringBackend) Name() string   { return platformKeyringName() }
func (KeyringBackend) IsSecure() bool { return true }

// IsAvailable probes the OS keyring with a harmless Get for a non-existent
// account. ErrNotFound means "backend works"; any other error means the daemon
// or service is missing (typical on headless Linux without libsecret).
func (KeyringBackend) IsAvailable() bool {
	_, err := keyring.Get(ServiceName, "__twelvedata_cli_probe__")
	if err == nil {
		return true
	}
	return errors.Is(err, keyring.ErrNotFound)
}

func (KeyringBackend) Get(account string) (string, error) {
	v, err := keyring.Get(ServiceName, account)
	if err != nil {
		if errors.Is(err, keyring.ErrNotFound) {
			return "", nil
		}
		return "", err
	}
	return v, nil
}

func (KeyringBackend) Set(account, secret string) error {
	return keyring.Set(ServiceName, account, secret)
}

func (KeyringBackend) Delete(account string) (bool, error) {
	if err := keyring.Delete(ServiceName, account); err != nil {
		if errors.Is(err, keyring.ErrNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
