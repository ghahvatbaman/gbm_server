package migrations

import (
	"context"
	"fmt"
	"log"

	"github.com/ghahvatbaman/gbm/internal/models"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] create users table")
		res, err := db.NewCreateTable().Model((*models.User)(nil)).Exec(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res)
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] drop users table")
		res, err := db.NewDropTable().Model((*models.User)(nil)).Exec(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res)
		return nil
	})
}
