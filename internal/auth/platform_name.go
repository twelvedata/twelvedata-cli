package auth

import "runtime"

func platformKeyringName() string {
	switch runtime.GOOS {
	case "darwin":
		return "macOS Keychain"
	case "windows":
		return "Windows Credential Manager"
	case "linux":
		return "Secret Service (libsecret)"
	default:
		return "OS keyring"
	}
}
