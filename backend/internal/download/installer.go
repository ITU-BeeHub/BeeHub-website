package download

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrUnsupportedPlatform = errors.New("unsupported platform")
	ErrInstallerNotFound   = errors.New("installer not found")
)

// Installer is an installer file on disk and the name it is served under.
type Installer struct {
	Path     string
	FileName string
}

// Catalog resolves the installer of the current release for each platform.
// Installers are expected at <dir>/installer_<version>.<ext>.
type Catalog struct {
	dir     string
	version string
}

func NewCatalog(dir, version string) Catalog {
	return Catalog{dir: dir, version: version}
}

var extensions = map[string]string{
	"windows": "exe",
	"mac":     "dmg",
}

// Lookup returns the installer for the given platform.
func (c Catalog) Lookup(platform string) (Installer, error) {
	ext, ok := extensions[platform]
	if !ok {
		return Installer{}, fmt.Errorf("%w: %q", ErrUnsupportedPlatform, platform)
	}

	path := filepath.Join(c.dir, fmt.Sprintf("installer_%s.%s", c.version, ext))
	if _, err := os.Stat(path); err != nil {
		return Installer{}, fmt.Errorf("%w: %s: %v", ErrInstallerNotFound, path, err)
	}

	return Installer{
		Path:     path,
		FileName: fmt.Sprintf("BeeHub-%s.%s", c.version, ext),
	}, nil
}
