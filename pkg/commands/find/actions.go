package find

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func actionFindDotGitFolders(cliContext *cli.Context) error {
	folders, _ := dotGitWalkDir(cliContext.String("directory"))
	for _, folder := range folders {
		fmt.Print(folder)
	}
	return nil
}

func dotGitWalkDir(startingDirectory string) (folders []string, err error) {
	startingDirectory = filepath.Clean(startingDirectory)

	err = filepath.WalkDir(startingDirectory, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && d.Name() == ".git" {
			folders = append(folders, path)
		}
		return nil
	})
	return folders, err
}
