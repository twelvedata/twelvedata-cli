package auth

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestStoreAPIKey_FileBackend(t *testing.T) {
	dir := setupTestDir(t)
	path, backend, err := StoreAPIKey("default", "secret")
	if err != nil {
		t.Fatal(err)
	}
	if backend.IsSecure() {
		t.Error("test setup forces file backend; got secure")
	}
	if path != filepath.Join(dir, "credentials.json") {
		t.Errorf("path = %q, want under %q", path, dir)
	}
	creds, err := ReadCredentials()
	if err != nil {
		t.Fatal(err)
	}
	if creds.Profiles["default"].APIKey != "secret" {
		t.Errorf("expected stored APIKey, got %+v", creds.Profiles["default"])
	}
	if creds.ActiveProfile != "default" {
		t.Errorf("active profile should be set on first store, got %q", creds.ActiveProfile)
	}
}

func TestStoreAPIKey_InvalidName(t *testing.T) {
	setupTestDir(t)
	if _, _, err := StoreAPIKey("", "secret"); err == nil {
		t.Error("expected error for empty profile name")
	}
}

func TestStoreAPIKey_DoesNotChangeActiveOnSecondAdd(t *testing.T) {
	setupTestDir(t)
	if _, _, err := StoreAPIKey("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	if _, _, err := StoreAPIKey("beta", "k2"); err != nil {
		t.Fatal(err)
	}
	creds, _ := ReadCredentials()
	if creds.ActiveProfile != "alpha" {
		t.Errorf("active profile should remain alpha, got %q", creds.ActiveProfile)
	}
}

func TestRemoveProfile_File(t *testing.T) {
	setupTestDir(t)
	if _, _, err := StoreAPIKey("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	if _, _, err := StoreAPIKey("beta", "k2"); err != nil {
		t.Fatal(err)
	}
	if err := RemoveProfile("alpha"); err != nil {
		t.Fatal(err)
	}
	creds, _ := ReadCredentials()
	if _, ok := creds.Profiles["alpha"]; ok {
		t.Error("alpha should be removed")
	}
	if creds.ActiveProfile != "beta" {
		t.Errorf("active profile should hop to beta, got %q", creds.ActiveProfile)
	}
}

func TestRemoveProfile_LastProfileRemovesFile(t *testing.T) {
	dir := setupTestDir(t)
	if _, _, err := StoreAPIKey("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	if err := RemoveProfile("alpha"); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, "credentials.json")); !os.IsNotExist(err) {
		t.Errorf("expected file removed, got err = %v", err)
	}
}

func TestRemoveProfile_NoCredentialsFile(t *testing.T) {
	setupTestDir(t)
	err := RemoveProfile("default")
	if err == nil {
		t.Fatal("expected error when no credentials file exists")
	}
}

func TestRemoveProfile_UnknownProfile(t *testing.T) {
	setupTestDir(t)
	if _, _, err := StoreAPIKey("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	err := RemoveProfile("ghost")
	var pne *ProfileNotFoundError
	if !errors.As(err, &pne) {
		t.Fatalf("expected ProfileNotFoundError, got %v", err)
	}
}

func TestRemoveAll_NoFile(t *testing.T) {
	setupTestDir(t)
	if err := RemoveAll(); err != nil {
		t.Errorf("RemoveAll on empty config should be a no-op, got %v", err)
	}
}

func TestRemoveAll_DeletesFile(t *testing.T) {
	dir := setupTestDir(t)
	if _, _, err := StoreAPIKey("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	if _, _, err := StoreAPIKey("beta", "k2"); err != nil {
		t.Fatal(err)
	}
	if err := RemoveAll(); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, "credentials.json")); !os.IsNotExist(err) {
		t.Errorf("expected file removed, got %v", err)
	}
}

func TestRenameProfile_File(t *testing.T) {
	setupTestDir(t)
	if _, _, err := StoreAPIKey("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	if err := RenameProfile("alpha", "gamma"); err != nil {
		t.Fatal(err)
	}
	creds, _ := ReadCredentials()
	if _, ok := creds.Profiles["alpha"]; ok {
		t.Error("old name should be gone")
	}
	if creds.Profiles["gamma"].APIKey != "k1" {
		t.Errorf("expected k1 under new name, got %+v", creds.Profiles["gamma"])
	}
	if creds.ActiveProfile != "gamma" {
		t.Errorf("active profile should be renamed, got %q", creds.ActiveProfile)
	}
}

func TestRenameProfile_SameNameIsNoOp(t *testing.T) {
	setupTestDir(t)
	if _, _, err := StoreAPIKey("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	if err := RenameProfile("alpha", "alpha"); err != nil {
		t.Errorf("renaming to same name should be a no-op, got %v", err)
	}
}

func TestRenameProfile_NewNameInvalid(t *testing.T) {
	setupTestDir(t)
	if _, _, err := StoreAPIKey("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	if err := RenameProfile("alpha", "bad name"); err == nil {
		t.Error("expected error for invalid new name")
	}
}

func TestRenameProfile_TargetExists(t *testing.T) {
	setupTestDir(t)
	if _, _, err := StoreAPIKey("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	if _, _, err := StoreAPIKey("beta", "k2"); err != nil {
		t.Fatal(err)
	}
	if err := RenameProfile("alpha", "beta"); err == nil {
		t.Error("expected error renaming to an existing name")
	}
}

func TestRenameProfile_Unknown(t *testing.T) {
	setupTestDir(t)
	if _, _, err := StoreAPIKey("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	if err := RenameProfile("ghost", "phantom"); err == nil {
		t.Error("expected error renaming an unknown profile")
	}
}

func TestRenameProfile_NoCredentialsFile(t *testing.T) {
	setupTestDir(t)
	if err := RenameProfile("a", "b"); err == nil {
		t.Error("expected error when no credentials file exists")
	}
}
