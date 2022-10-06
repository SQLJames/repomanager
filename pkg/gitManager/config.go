package gitmanager

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/sqljames/repomanager/pkg/util/jlogr"
)

func GetOrigin(dotGitFolder string) (urls []string) {
	repository, err := git.PlainOpen(filepath.Clean(dotGitFolder))
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "this sucks", "path", filepath.Clean(dotGitFolder))

		return nil
	}

	remote, err := repository.Remote("origin")
	if err != nil {
		jlogr.Logger.ILog.Error(err, "Seems like Repo doesn't have an Origin.", "path", filepath.Clean(dotGitFolder))

		return nil
	}

	return remote.Config().URLs
}

// https://github.com/SQLJames/gSupportCollector.
// git@github.com:SQLJames/factorio-docker.git.
func ParseOriginToDirectory(originURL string) (directory string) {
	if strings.Contains(originURL, "@") {
		return parseSSHOrigin(originURL)
	}

	return parseHTTPOrigin(originURL)
}

// https://github.com/SQLJames/gSupportCollector.
func parseHTTPOrigin(originURL string) (directory string) {
	return strings.Split(originURL, "//")[1]
}

// git@github.com:SQLJames/factorio-docker.git.
func parseSSHOrigin(originURL string) (directory string) {
	originURL = strings.Split(originURL, "@")[1]
	originURL = strings.ReplaceAll(originURL, ".git", "")

	return path.Join(strings.Split(originURL, ":")...)
}
