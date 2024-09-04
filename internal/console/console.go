package console

import (
	"log"
	"os"

	"github.com/rkashapov2015/webproject/internal/commands"
	"github.com/rkashapov2015/webproject/internal/database/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
)

type App struct {
	db  *bun.DB
	cli *cli.App
}

func New(db *bun.DB) *App {
	cli := &cli.App{
		Name: "Console",
		Commands: []*cli.Command{
			commands.NewDBCommand(migrate.NewMigrator(db, migrations.Migrations)),
			commands.NewUserCommand(db),
		},
	}

	app := &App{
		db:  db,
		cli: cli,
	}

	return app
}

func (app *App) Run() error {
	if err := app.cli.Run(os.Args); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
