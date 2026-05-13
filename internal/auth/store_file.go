package auth

import "fmt"

// FileBackend stores API keys directly inside credentials.json. It is always
// available and never considered secure.
type FileBackend struct{}

func NewFileBackend() *FileBackend { return &FileBackend{} }

func (FileBackend) Name() string   { return "plaintext file" }
func (FileBackend) IsSecure() bool { return false }
func (FileBackend) IsAvailable() bool { return true }

func (FileBackend) Get(account string) (string, error) {
	creds, err := ReadCredentials()
	if err != nil {
		return "", err
	}
	if creds == nil {
		return "", nil
	}
	return creds.Profiles[account].APIKey, nil
}

func (FileBackend) Set(account, secret string) error {
	if err := ValidateProfileName(account); err != nil {
		return err
	}
	creds, err := ReadCredentials()
	if err != nil {
		return err
	}
	if creds == nil {
		creds = &CredentialsFile{ActiveProfile: defaultProfile, Profiles: map[string]Profile{}}
	}
	creds.Profiles[account] = Profile{APIKey: secret}
	if len(creds.Profiles) == 1 {
		creds.ActiveProfile = account
	}
	// Don't clear an existing "secure_storage" marker — other profiles may still
	// have keys in the keyring. The resolver falls through correctly in either case.
	if _, err := WriteCredentials(creds); err != nil {
		return err
	}
	return nil
}

func (FileBackend) Delete(account string) (bool, error) {
	creds, err := ReadCredentials()
	if err != nil {
		return false, err
	}
	if creds == nil {
		return false, nil
	}
	if _, ok := creds.Profiles[account]; !ok {
		return false, fmt.Errorf("profile %q not found", account)
	}
	delete(creds.Profiles, account)
	if creds.ActiveProfile == account {
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
		if err := DeleteCredentialsFile(); err != nil {
			return false, err
		}
		return true, nil
	}
	if _, err := WriteCredentials(creds); err != nil {
		return false, err
	}
	return true, nil
}
