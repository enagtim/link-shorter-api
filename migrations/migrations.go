package migrations

import (
	"go/adv-demo/internal/link"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/db"
)

func Migrate(db *db.Db) {
	err := db.AutoMigrate(&link.Link{}, &user.User{})
	if err != nil {
		panic(err)
	}
}
