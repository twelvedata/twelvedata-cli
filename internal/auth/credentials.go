package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// Storage tags the on-disk credentials file with where the actual key material
// lives. Profiles in a "secure_storage" file have empty APIKey fields; their
// keys live in the OS keyring under (ServiceName, profile-name).
type Storage string

const (
	StorageFile    Storage = "file"
	StorageSecure  Storage = "secure_storage"
	defaultProfile         = "default"
)

// Profile is one entry inside credentials.json. APIKey is only populated when
// the file backend is in use; for secure storage the field is omitted.
type Profile struct {
	APIKey string `json:"api_key,omitempty"`
}

// CredentialsFile is the persisted shape of credentials.json. Storage is
// optional; absence means "file".
type CredentialsFile struct {
	ActiveProfile string             `json:"active_profile"`
	Storage       Storage            `json:"storage,omitempty"`
	Profiles      map[string]Profile `json:"profiles"`
}

var profileNameRe = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)

// ValidateProfileName returns an error if the name is empty, too long, or
// contains characters outside [A-Za-z0-9._-].
func ValidateProfileName(name string) error {
	if name == "" {
		return fmt.Errorf("profile name must not be empty")
	}
	if len(name) > 64 {
		return fmt.Errorf("profile name must be 64 characters or fewer")
	}
	if !profileNameRe.MatchString(name) {
		return fmt.Errorf("profile name must contain only letters, numbers, dots, dashes, and underscores")
	}
	return nil
}

// ReadCredentials loads credentials.json. Returns (nil, nil) if the file is
// missing or unparseable — callers treat that as "no profiles configured".
func ReadCredentials() (*CredentialsFile, error) {
	data, err := os.ReadFile(CredentialsPath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var creds CredentialsFile
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, nil
	}
	if creds.Profiles == nil {
		return nil, nil
	}
	if creds.ActiveProfile == "" {
		creds.ActiveProfile = defaultProfile
	}
	return &creds, nil
}

// WriteCredentials persists creds to disk with file mode 0600 and parent dir
// 0700. Returns the path written.
func WriteCredentials(creds *CredentialsFile) (string, error) {
	dir := GetConfigDir()
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", err
	}
	path := filepath.Join(dir, credentialsFileName)
	buf, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return "", err
	}
	buf = append(buf, '\n')
	if err := os.WriteFile(path, buf, 0o600); err != nil {
		return "", err
	}
	// Best-effort chmod in case the file existed with looser permissions.
	_ = os.Chmod(path, 0o600)
	return path, nil
}

// DeleteCredentialsFile removes credentials.json. Missing file is not an error.
func DeleteCredentialsFile() error {
	err := os.Remove(CredentialsPath())
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// MaskKey returns "abc...wxyz" — first 3 and last 4 chars — for display.
func MaskKey(key string) string {
	if len(key) <= 7 {
		if len(key) <= 3 {
			return key + "..."
		}
		return key[:3] + "..."
	}
	return key[:3] + "..." + key[len(key)-4:]
}

// ListProfiles returns the configured profiles, in map iteration order, each
// annotated with whether it is the active one.
type ProfileInfo struct {
	Name   string
	Active bool
}

func ListProfiles() ([]ProfileInfo, error) {
	creds, err := ReadCredentials()
	if err != nil {
		return nil, err
	}
	if creds == nil {
		return nil, nil
	}
	out := make([]ProfileInfo, 0, len(creds.Profiles))
	for name := range creds.Profiles {
		out = append(out, ProfileInfo{Name: name, Active: name == creds.ActiveProfile})
	}
	return out, nil
}
