package migrations

import (
	"context"
	"fmt"

	"github.com/rkashapov2015/webproject/internal/database/models"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		_, err := db.NewCreateTable().
			Model((*models.Note)(nil)).
			ForeignKey(`("author_id") REFERENCES "users" ("id")`).
			Exec(ctx)
		if err != nil {
			panic(err)
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		_, err := db.NewDropTable().Model((*models.Note)(nil)).IfExists().Exec(ctx)
		if err != nil {
			panic(err)
		}
		return nil
	})
}
