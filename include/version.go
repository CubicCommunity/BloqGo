package include

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/semver"
)

var VERSION string = "0.0.1"

// Gets the string of the current version of BloqGo
func Version() (string, error) {
	if semver.IsValid(fmt.Sprintf("v%s", VERSION)) {
		return VERSION, nil
	} else {
		exePath, err := os.Executable()

		if err != nil {
			return "", err
		} else {
			projectRoot := filepath.Dir(filepath.Dir(exePath))

			v, e := os.ReadFile(filepath.Join(projectRoot, "VERSION"))

			if e != nil {
				return "", e
			} else {
				return strings.TrimSpace(string(v)), e
			}
		}
	}
}
