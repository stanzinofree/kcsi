package version

import (
	_ "embed" // Embed directive is used to embed version.yaml into the binary
	"fmt"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed version.yaml
var manifestData []byte

// Build information injected via ldflags during build
var (
	version   string // Injected from git tag (e.g., "v0.6.3" -> "0.6.3")
	buildDate string
	gitCommit string
)

// Manifest holds the version and metadata information
type Manifest struct {
	Version      string   `yaml:"version"`
	Name         string   `yaml:"name"`
	FullName     string   `yaml:"fullName"`
	Description  string   `yaml:"description"`
	Author       string   `yaml:"author"`
	License      string   `yaml:"license"`
	Repository   string   `yaml:"repository"`
	BuildDate    string   `yaml:"buildDate"`
	Spirit       string   `yaml:"spirit"`
	Contributors []string `yaml:"contributors"`
	Releases     struct {
		Latest string `yaml:"latest"`
		Stable string `yaml:"stable"`
	} `yaml:"releases"`
}

var manifest Manifest

func init() {
	if yaml.Unmarshal(manifestData, &manifest) != nil {
		// Fallback to defaults if manifest cannot be read
		manifest.Version = "dev"
		manifest.Name = "kcsi"
		manifest.FullName = "Kubectl Smart Interactive"
		manifest.Description = "A kubectl wrapper with intelligent autocompletion"
		manifest.Author = "Alessandro"
	}
}

// GetVersion returns the current version
// Priority: injected version (from git tag) > manifest version > "dev"
func GetVersion() string {
	if version != "" {
		return version
	}
	if manifest.Version != "" {
		return manifest.Version
	}
	return "dev"
}

// GetVersionInfo returns formatted version information
func GetVersionInfo() string {
	return fmt.Sprintf("%s version %s", manifest.Name, manifest.Version)
}

// GetDetailedVersion returns detailed version information
func GetDetailedVersion() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s (%s)\n", manifest.FullName, manifest.Name))
	sb.WriteString(fmt.Sprintf("Version: %s\n", manifest.Version))
	sb.WriteString(fmt.Sprintf("Author: %s\n", manifest.Author))

	// Use injected buildDate if available, otherwise fall back to manifest
	if buildDate != "" {
		sb.WriteString(fmt.Sprintf("Build Date: %s\n", buildDate))
	} else if manifest.BuildDate != "" {
		sb.WriteString(fmt.Sprintf("Build Date: %s\n", manifest.BuildDate))
	}

	// Show git commit if available
	if gitCommit != "" {
		sb.WriteString(fmt.Sprintf("Git Commit: %s\n", gitCommit))
	}

	sb.WriteString(fmt.Sprintf("Go Version: %s\n", runtime.Version()))
	sb.WriteString(fmt.Sprintf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH))

	return sb.String()
}

// GetAbout returns comprehensive information about the project
func GetAbout() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("╔══════════════════════════════════════════════════════════════╗\n"))
	sb.WriteString(fmt.Sprintf("║  %s v%-50s║\n", manifest.FullName, manifest.Version))
	sb.WriteString(fmt.Sprintf("╚══════════════════════════════════════════════════════════════╝\n\n"))

	sb.WriteString(fmt.Sprintf("%s\n\n", manifest.Description))

	sb.WriteString("Spirit & Philosophy:\n")
	sb.WriteString(strings.TrimSpace(manifest.Spirit))
	sb.WriteString("\n\n")

	sb.WriteString(fmt.Sprintf("Author: %s\n", manifest.Author))

	if len(manifest.Contributors) > 1 {
		sb.WriteString("Contributors:\n")
		for _, contributor := range manifest.Contributors {
			sb.WriteString(fmt.Sprintf("  - %s\n", contributor))
		}
		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf("License: %s\n", manifest.License))
	sb.WriteString(fmt.Sprintf("Repository: %s\n", manifest.Repository))
	sb.WriteString(fmt.Sprintf("Version: %s\n", manifest.Version))

	if manifest.BuildDate != "" {
		sb.WriteString(fmt.Sprintf("Build Date: %s\n", manifest.BuildDate))
	}

	sb.WriteString(fmt.Sprintf("\nBuilt with Go %s for %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH))

	return sb.String()
}

// GetManifest returns the full manifest
func GetManifest() Manifest {
	return manifest
}

// GetAuthor returns the author name
func GetAuthor() string {
	return manifest.Author
}

// GetBuildDate returns the build date (injected via ldflags or from manifest)
func GetBuildDate() string {
	if buildDate != "" {
		return buildDate
	}
	if manifest.BuildDate != "" {
		return manifest.BuildDate
	}
	return "unknown"
}

// GetGitCommit returns the git commit hash (injected via ldflags)
func GetGitCommit() string {
	if gitCommit != "" {
		return gitCommit
	}
	return "unknown"
}
