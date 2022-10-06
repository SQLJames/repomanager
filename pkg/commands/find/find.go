package find

import "github.com/urfave/cli/v2"

func New() *cli.Command {
	return &cli.Command{
		Name:    "find",
		Usage:   "Allows users to find objects",
		Aliases: []string{"f"},
		Subcommands: []*cli.Command{
			repos(),
		},
	}
}

func repos() *cli.Command {
	return &cli.Command{
		Name:   "repos",
		Usage:  "Prints all the notebooks in your journal",
		Action: actionFindDotGitFolders,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "directory",
				Usage:    "used to define folder to search in",
				Required: true,
			},
		},
	}
}
