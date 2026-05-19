package auth

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileBackend_Flags(t *testing.T) {
	b := NewFileBackend()
	if b.IsSecure() {
		t.Error("FileBackend should not be secure")
	}
	if !b.IsAvailable() {
		t.Error("FileBackend should always be available")
	}
	if b.Name() == "" {
		t.Error("FileBackend.Name must not be empty")
	}
}

func TestFileBackend_GetMissing(t *testing.T) {
	setupTestDir(t)
	b := NewFileBackend()
	v, err := b.Get("default")
	if err != nil {
		t.Fatal(err)
	}
	if v != "" {
		t.Errorf("expected empty key when file missing, got %q", v)
	}
}

func TestFileBackend_SetGet(t *testing.T) {
	setupTestDir(t)
	b := NewFileBackend()
	if err := b.Set("default", "secret"); err != nil {
		t.Fatal(err)
	}
	v, err := b.Get("default")
	if err != nil {
		t.Fatal(err)
	}
	if v != "secret" {
		t.Errorf("got %q, want %q", v, "secret")
	}
}

func TestFileBackend_SetRejectsInvalidName(t *testing.T) {
	setupTestDir(t)
	b := NewFileBackend()
	if err := b.Set("", "secret"); err == nil {
		t.Error("expected error for empty profile name")
	}
	if err := b.Set("bad/name", "secret"); err == nil {
		t.Error("expected error for invalid profile name")
	}
}

func TestFileBackend_SetActivatesFirstProfile(t *testing.T) {
	setupTestDir(t)
	b := NewFileBackend()
	if err := b.Set("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	creds, _ := ReadCredentials()
	if creds == nil || creds.ActiveProfile != "alpha" {
		t.Fatalf("first Set should activate the profile, got %+v", creds)
	}
	if err := b.Set("beta", "k2"); err != nil {
		t.Fatal(err)
	}
	creds, _ = ReadCredentials()
	if creds.ActiveProfile != "alpha" {
		t.Fatalf("second Set must not change active profile, got %q", creds.ActiveProfile)
	}
}

func TestFileBackend_Delete(t *testing.T) {
	dir := setupTestDir(t)
	b := NewFileBackend()
	if err := b.Set("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	if err := b.Set("beta", "k2"); err != nil {
		t.Fatal(err)
	}
	ok, err := b.Delete("alpha")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Error("expected ok=true when deleting an existing profile")
	}
	creds, _ := ReadCredentials()
	if _, exists := creds.Profiles["alpha"]; exists {
		t.Error("alpha should be gone")
	}
	if creds.ActiveProfile != "beta" {
		t.Errorf("active profile should hop to remaining one, got %q", creds.ActiveProfile)
	}

	// Delete the last profile — file is removed.
	if _, err := b.Delete("beta"); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, "credentials.json")); !os.IsNotExist(err) {
		t.Errorf("credentials.json should be removed when no profiles remain, got %v", err)
	}
}

func TestFileBackend_DeleteUnknownProfile(t *testing.T) {
	setupTestDir(t)
	b := NewFileBackend()
	if err := b.Set("alpha", "k1"); err != nil {
		t.Fatal(err)
	}
	_, err := b.Delete("nonexistent")
	if err == nil {
		t.Fatal("expected error when deleting unknown profile")
	}
}
