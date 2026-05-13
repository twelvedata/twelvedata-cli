package auth

import (
	"errors"
	"os"
)

// Source labels where an API key came from. It powers `td whoami` and lets
// callers report a precise origin in error messages.
type Source string

const (
	SourceFlag   Source = "flag"
	SourceEnv    Source = "env"
	SourceConfig Source = "config"
	SourceSecure Source = "secure_storage"
	envAPIKey           = "TWELVEDATA_API_KEY"
	envProfile          = "TWELVEDATA_PROFILE"
)

// ResolvedKey is the result of the resolution chain. Profile is empty when
// Source is Flag or Env.
type ResolvedKey struct {
	Key     string
	Source  Source
	Profile string
}

// ErrNoAPIKey is returned when nothing in the resolution chain yielded a key.
// The internal/output classifier maps it to exit code 3 (auth failure).
var ErrNoAPIKey = errors.New("no Twelve Data API key configured. Run `td login` or set TWELVEDATA_API_KEY")

// ResolveProfileName picks the profile name to use, in order: explicit flag,
// TWELVEDATA_PROFILE env var, credentials.json active_profile, "default".
func ResolveProfileName(flagValue string) string {
	if flagValue != "" {
		return flagValue
	}
	if v := os.Getenv(envProfile); v != "" {
		return v
	}
	creds, _ := ReadCredentials()
	if creds != nil && creds.ActiveProfile != "" {
		return creds.ActiveProfile
	}
	return defaultProfile
}

// ResolveAPIKey applies the full resolution chain: flag → env → active profile
// (secure storage or file). The returned ResolvedKey records which step won.
func ResolveAPIKey(flagValue, profileFlag string) (*ResolvedKey, error) {
	if flagValue != "" {
		return &ResolvedKey{Key: flagValue, Source: SourceFlag}, nil
	}
	if v := os.Getenv(envAPIKey); v != "" {
		return &ResolvedKey{Key: v, Source: SourceEnv}, nil
	}
	creds, err := ReadCredentials()
	if err != nil {
		return nil, err
	}
	profile := ResolveProfileName(profileFlag)
	if creds == nil {
		return nil, ErrNoAPIKey
	}
	entry, ok := creds.Profiles[profile]
	if !ok {
		return nil, ErrNoAPIKey
	}
	if creds.Storage == StorageSecure {
		backend := GetBackend()
		key, err := backend.Get(profile)
		if err != nil {
			return nil, err
		}
		if key == "" {
			return nil, ErrNoAPIKey
		}
		return &ResolvedKey{Key: key, Source: SourceSecure, Profile: profile}, nil
	}
	if entry.APIKey == "" {
		return nil, ErrNoAPIKey
	}
	return &ResolvedKey{Key: entry.APIKey, Source: SourceConfig, Profile: profile}, nil
}

// SetActiveProfile updates the active_profile field in credentials.json.
func SetActiveProfile(name string) error {
	if err := ValidateProfileName(name); err != nil {
		return err
	}
	creds, err := ReadCredentials()
	if err != nil {
		return err
	}
	if creds == nil {
		return errors.New("no credentials file found. Run `td login` first")
	}
	if _, ok := creds.Profiles[name]; !ok {
		return errProfileNotFound(name, creds)
	}
	creds.ActiveProfile = name
	if _, err := WriteCredentials(creds); err != nil {
		return err
	}
	return nil
}

func errProfileNotFound(name string, creds *CredentialsFile) error {
	names := make([]string, 0, len(creds.Profiles))
	for n := range creds.Profiles {
		names = append(names, n)
	}
	return &ProfileNotFoundError{Name: name, Available: names}
}

// ProfileNotFoundError carries the requested name and the available list so
// callers can format a useful message.
type ProfileNotFoundError struct {
	Name      string
	Available []string
}

func (e *ProfileNotFoundError) Error() string {
	if len(e.Available) == 0 {
		return "profile \"" + e.Name + "\" not found. No profiles configured"
	}
	out := "profile \"" + e.Name + "\" not found. Available: "
	for i, n := range e.Available {
		if i > 0 {
			out += ", "
		}
		out += n
	}
	return out
}
