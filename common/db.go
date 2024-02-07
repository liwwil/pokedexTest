package common

import (
	"github.com/TanyaEIEI/pokedex/graph/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.Pokemon{}, &model.PokemonType{}, &model.PokemonCategory{}, &model.PokemonAbility{})
	return db, nil
}
