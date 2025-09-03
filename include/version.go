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
		cwd, err := os.Getwd()

		if err != nil {
			return cwd, err
		} else {
			v, err := os.ReadFile(filepath.Join(cwd, "..", "VERSION"))

			if err != nil {
				return "", err
			} else {
				return strings.TrimSpace(string(v)), err
			}
		}
	}
}
