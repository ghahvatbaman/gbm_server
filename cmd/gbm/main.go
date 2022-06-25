package main

import (
	"github.com/ghahvatbaman/gbm/config"
	"github.com/ghahvatbaman/gbm/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID    int64 `bun:",pk,autoincrement"`
	Name  string
	Email string
	Phome string
}

var DB *bun.DB

func main() {

	config.SetupConfig()
	db, err := config.SetupDb()
	DB = db

	if err != nil {
		panic(err)
	}

	app := fiber.New()
	router.SetupRoutes(app)

	app.Listen(":3000")
}
