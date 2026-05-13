package auth

import (
	"errors"
	"fmt"
	"os"
)

// StoreAPIKey writes the key for `profile` to the active backend and updates
// credentials.json accordingly. Returns the config path and the backend used.
func StoreAPIKey(profile, apiKey string) (configPath string, backend CredentialBackend, err error) {
	if err := ValidateProfileName(profile); err != nil {
		return "", nil, err
	}
	backend = GetBackend()
	creds, err := ReadCredentials()
	if err != nil {
		return "", nil, err
	}
	if creds == nil {
		creds = &CredentialsFile{ActiveProfile: profile, Profiles: map[string]Profile{}}
	}
	if backend.IsSecure() {
		if err := backend.Set(profile, apiKey); err != nil {
			return "", backend, err
		}
		creds.Storage = StorageSecure
		creds.Profiles[profile] = Profile{}
	} else {
		creds.Profiles[profile] = Profile{APIKey: apiKey}
		// Leave any existing Storage marker alone — other profiles may still live
		// in secure storage.
	}
	if len(creds.Profiles) == 1 {
		creds.ActiveProfile = profile
	}
	path, err := WriteCredentials(creds)
	if err != nil {
		return "", backend, err
	}
	return path, backend, nil
}

// RemoveProfile deletes a single profile from both the backend and the
// credentials file. If the profile was active, the next available profile (or
// "default") becomes active. If no profiles remain, the file is deleted.
func RemoveProfile(profile string) error {
	creds, err := ReadCredentials()
	if err != nil {
		return err
	}
	if creds == nil {
		return errors.New("no credentials file found")
	}
	if _, ok := creds.Profiles[profile]; !ok {
		return errProfileNotFound(profile, creds)
	}
	if creds.Storage == StorageSecure {
		backend := GetBackend()
		if backend.IsSecure() {
			if _, err := backend.Delete(profile); err != nil {
				return fmt.Errorf("failed to remove credential from %s: %w", backend.Name(), err)
			}
		}
	}
	delete(creds.Profiles, profile)
	if creds.ActiveProfile == profile {
		creds.ActiveProfile = ""
		for name := range creds.Profiles {
			creds.ActiveProfile = name
			break
		}
		if creds.ActiveProfile == "" {
			creds.ActiveProfile = defaultProfile
		}
	}
	if len(creds.Profiles) == 0 {
		return DeleteCredentialsFile()
	}
	_, err = WriteCredentials(creds)
	return err
}

// RemoveAll wipes every profile, removing each from secure storage when in use,
// and then deletes credentials.json. Partial failures rewrite the file with
// successfully-removed profiles dropped, so retries only re-attempt failures.
func RemoveAll() error {
	creds, err := ReadCredentials()
	if err != nil {
		return err
	}
	if creds == nil {
		path := CredentialsPath()
		if _, statErr := os.Stat(path); statErr == nil {
			return os.Remove(path)
		}
		return nil
	}
	if creds.Storage == StorageSecure {
		backend := GetBackend()
		if backend.IsSecure() {
			var failed []string
			for name := range creds.Profiles {
				if _, err := backend.Delete(name); err != nil {
					failed = append(failed, name)
				}
			}
			if len(failed) > 0 {
				for name := range creds.Profiles {
					if !contains(failed, name) {
						delete(creds.Profiles, name)
					}
				}
				if creds.ActiveProfile != "" {
					if _, ok := creds.Profiles[creds.ActiveProfile]; !ok {
						creds.ActiveProfile = ""
						for n := range creds.Profiles {
							creds.ActiveProfile = n
							break
						}
						if creds.ActiveProfile == "" {
							creds.ActiveProfile = defaultProfile
						}
					}
				}
				_, _ = WriteCredentials(creds)
				return fmt.Errorf("failed to remove credentials for: %v. Retry to clean them up", failed)
			}
		}
	}
	return DeleteCredentialsFile()
}

// RenameProfile swaps a profile's name in-place. When the storage backend is
// secure, the secret is migrated under (ServiceName, newName) first and the
// old entry is removed; if removal fails the migration is rolled back.
func RenameProfile(oldName, newName string) error {
	if oldName == newName {
		return nil
	}
	if err := ValidateProfileName(newName); err != nil {
		return err
	}
	creds, err := ReadCredentials()
	if err != nil {
		return err
	}
	if creds == nil {
		return errors.New("no credentials file found")
	}
	entry, ok := creds.Profiles[oldName]
	if !ok {
		return errProfileNotFound(oldName, creds)
	}
	if _, exists := creds.Profiles[newName]; exists {
		return fmt.Errorf("profile %q already exists", newName)
	}
	if creds.Storage == StorageSecure {
		backend := GetBackend()
		if backend.IsSecure() {
			key, err := backend.Get(oldName)
			if err != nil {
				return fmt.Errorf("failed to read credential from %s: %w", backend.Name(), err)
			}
			if key != "" {
				if err := backend.Set(newName, key); err != nil {
					return fmt.Errorf("failed to write credential to %s: %w", backend.Name(), err)
				}
				if _, err := backend.Delete(oldName); err != nil {
					if _, rbErr := backend.Delete(newName); rbErr != nil {
						return fmt.Errorf("failed to remove old credential %q from %s; rollback of new credential also failed", oldName, backend.Name())
					}
					return fmt.Errorf("failed to remove old credential %q from %s; rename rolled back", oldName, backend.Name())
				}
			}
		}
	}
	creds.Profiles[newName] = entry
	delete(creds.Profiles, oldName)
	if creds.ActiveProfile == oldName {
		creds.ActiveProfile = newName
	}
	_, err = WriteCredentials(creds)
	return err
}

func contains(s []string, v string) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}
