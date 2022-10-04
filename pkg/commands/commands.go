package commands

import (
	"github.com/sqljames/repomanager/pkg/flags"
	"github.com/sqljames/repomanager/pkg/info"
	"github.com/sqljames/repomanager/pkg/commands/find"
	"github.com/sqljames/repomanager/pkg/util/jlogr"
	"github.com/sqljames/repomanager/pkg/version"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func NewApp() *cli.App {
	modifyCLIDefaultVersion()

	app := &cli.App{
		Name:      info.GetApplicationName(),
		Usage:     info.Description,
		Flags:     flags.GenerateGlobalFlags(),
		Version:   version.Version.String(),
		Before:    beforeTasks,
		Authors:   info.Authors,
		Copyright: info.Copyright,
		Suggest:   true,
		Commands:  []*cli.Command{
			find.New(),
		},
	}

	return app
}

func beforeTasks(cliContext *cli.Context) error {
	taskList := []func(*cli.Context) error{
		instrumentLoggingFlags,
	}

	for _, t := range taskList {
		if err := t(cliContext); err != nil {
			return err
		}
	}

	return nil
}

// Care of: https://github.com/physcat/klog-cli/blob/main/main.go
func instrumentLoggingFlags(cliContext *cli.Context) error {
	// Command line flags always overwrite configuration files
	first := altsrc.InitInputSourceWithContext(cliContext.App.Flags, altsrc.NewYamlSourceFromFlagFunc("config"))

	err := first(cliContext)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
	}

	// The second config map will not overwrite the first
	second := altsrc.InitInputSourceWithContext(cliContext.App.Flags, altsrc.NewYamlSourceFromFlagFunc("global-config"))

	err = second(cliContext)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
	}

	jlogr.InitializeLogger(cliContext.Int("verbosity"))

	return err
}

func modifyCLIDefaultVersion() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "Prints the version",
		Value: false,
	}
	cli.VersionPrinter = func(cCtx *cli.Context) {
		version.Version.CliPrinter()
	}
}
