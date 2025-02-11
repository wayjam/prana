// Command Line Interface of GOM.
package main

import (
	"os"
	"sort"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/phogolabs/cli"
	"github.com/phogolabs/prana/cmd"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:   "log-level",
		Value:  "info",
		Usage:  "level of logging",
		EnvVar: "PRANA_LOG_LEVEL",
	},
	&cli.StringFlag{
		Name:   "log-format",
		Value:  "",
		Usage:  "format of the logs",
		EnvVar: "PRANA_LOG_FORMAT",
	},
	&cli.StringFlag{
		Name:   "database-url",
		Value:  "sqlite3://prana.db",
		Usage:  "Database URL",
		EnvVar: "PRANA_DB_URL",
	},
}

func main() {
	var (
		migration  = &cmd.SQLMigration{}
		routine    = &cmd.SQLRoutine{}
		model      = &cmd.SQLModel{}
		repository = &cmd.SQLRepository{}
	)

	commands := []*cli.Command{
		migration.CreateCommand(),
		routine.CreateCommand(),
		model.CreateCommand(),
		repository.CreateCommand(),
	}

	app := &cli.App{
		Name:      "prana",
		HelpName:  "prana",
		Usage:     "Golang Database Manager",
		UsageText: "prana [global options]",
		Version:   "1.0-beta-05",
		Writer:    os.Stdout,
		ErrWriter: os.Stderr,
		Flags:     flags,
		Before:    cmd.BeforeEach,
		Commands:  commands,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	for _, command := range commands {
		sort.Sort(cli.FlagsByName(command.Flags))
		sort.Sort(cli.CommandsByName(command.Commands))
	}

	app.Run(os.Args)
}
