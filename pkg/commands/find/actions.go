package find

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"

	cp "github.com/otiai10/copy"
	"github.com/sqljames/repomanager/pkg/gitManager"
	"github.com/sqljames/repomanager/pkg/util"
	"github.com/sqljames/repomanager/pkg/util/jlogr"
	"github.com/urfave/cli/v2"
)

type GitRepository struct {
	Path string
}

func actionFindDotGitFolders(cliContext *cli.Context) error {
	programPath := util.MakeStorageLocation()
	folders, _ := dotGitWalkDir(cliContext.String("directory"))
	for _, folder := range folders {
		urls := gitManager.GetOrigin(folder)
		for _, url := range urls {
			originalDestination := folder[0 : len(folder)-4]
			finalDestination := path.Join(programPath, gitManager.ParseOriginToDirectory(url))
			jlogr.Logger.ILog.Debug("Moving files", "From", originalDestination, "To", finalDestination, "origin", url)
			err := os.MkdirAll(finalDestination, util.FolderPermissons)
			if err != nil {
				return err
			}
			cp.Copy(originalDestination, finalDestination)
		}

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
