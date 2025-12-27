package version

import (
	"strings"
	"testing"
)

func TestGetVersion(t *testing.T) {
	version := GetVersion()
	if version == "" {
		t.Error("GetVersion() returned empty string")
	}
	// Version should follow semantic versioning pattern (e.g., 0.6.0, 1.0.0)
	if !strings.Contains(version, ".") {
		t.Errorf("Version should contain dots (semantic versioning), got %s", version)
	}
}

func TestGetAuthor(t *testing.T) {
	author := GetAuthor()
	if author == "" {
		t.Error("GetAuthor() returned empty string")
	}
	if !strings.Contains(author, "Alessandro") {
		t.Errorf("Expected author to contain 'Alessandro', got %s", author)
	}
}

func TestGetDetailedVersion(t *testing.T) {
	detailed := GetDetailedVersion()
	if detailed == "" {
		t.Error("GetDetailedVersion() returned empty string")
	}
	if !strings.Contains(detailed, "Version:") {
		t.Error("GetDetailedVersion() should contain 'Version:'")
	}
	if !strings.Contains(detailed, "Author:") {
		t.Error("GetDetailedVersion() should contain 'Author:'")
	}
}

func TestGetAbout(t *testing.T) {
	about := GetAbout()
	if about == "" {
		t.Error("GetAbout() returned empty string")
	}
	if !strings.Contains(about, "kcsi") {
		t.Error("GetAbout() should contain 'kcsi'")
	}
}

func TestGetManifest(t *testing.T) {
	manifest := GetManifest()
	if manifest.Name != "kcsi" {
		t.Errorf("Expected manifest name 'kcsi', got %s", manifest.Name)
	}
	// Check version is not empty and follows semantic versioning
	if manifest.Version == "" {
		t.Error("Manifest version should not be empty")
	}
	if !strings.Contains(manifest.Version, ".") {
		t.Errorf("Manifest version should follow semantic versioning (contain dots), got %s", manifest.Version)
	}
	if manifest.License != "MIT" {
		t.Errorf("Expected manifest license 'MIT', got %s", manifest.License)
	}
}
