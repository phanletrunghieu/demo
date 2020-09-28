package sqlite3

import (
	"context"

	"service-1/model"

	otgorm "github.com/smacker/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var gDB *gorm.DB

func init() {
	var err error

	gDB, err = gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	// register callbacks must be called for a root instance of your gorm.DB
	otgorm.AddGormCallbacks(gDB)

	gDB.AutoMigrate(model.Bank{})
}

func GetDB(ctx context.Context) *gorm.DB {
	return otgorm.SetSpanToGorm(ctx, gDB)
}
