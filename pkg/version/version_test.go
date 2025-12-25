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
	if version != "0.5.3" {
		t.Errorf("Expected version 0.5.3, got %s", version)
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
	if manifest.Version != "0.5.3" {
		t.Errorf("Expected manifest version '0.5.3', got %s", manifest.Version)
	}
	if manifest.License != "MIT" {
		t.Errorf("Expected manifest license 'MIT', got %s", manifest.License)
	}
}
