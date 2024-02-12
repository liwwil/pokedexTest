package database

import (
	"context"
	"fmt"
	"log"

	"github.com/TanyaEIEI/pokedex/graph/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const dbFile = "testdb1.db"

type Database struct {
	Db *gorm.DB
}

func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Pokemon{})
	return db, nil
}

func (pkb *Database) CreatePokemon(ctx context.Context, input Pokemon) (*Pokemon, error) {

	// Convert input from GraphQL model to database model
	newCharacter := Pokemon{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Abilities:   input.Abilities,
		Type:        input.Type,
	}
	if err := pkb.Db.Create(&newCharacter).Error; err != nil {
		fmt.Printf("Failed to create Pokemon: %v", err)
		return nil, fmt.Errorf("failed to create Pokemon")
	}

	return &newCharacter, nil
}

func (pkb *Database) ListPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	var pokemonList []*model.Pokemon
	err := pkb.Db.Find(&pokemonList).Error
	if err != nil {
		return nil, err
	}
	return pokemonList, nil
}

func (pkb *Database) UpdatePokemon(ctx context.Context, input UpdatePokemonInput) (*Pokemon, error) {
	var pokemonData *Pokemon
	if err := pkb.Db.First(&pokemonData, input.ID).Error; err != nil {
		fmt.Printf("Failed to find Pokemon for update: %v", err)
		return nil, fmt.Errorf("failed to find Pokemon for update")
	}

	pokemonData.Name = *input.Name
	pokemonData.Description = *input.Description
	pokemonData.Category = *input.Category
	pokemonData.Abilities = *input.Abilities
	pokemonData.Type = *input.Type

	// Update other fields as needed

	if err := pkb.Db.Save(&pokemonData).Error; err != nil {
		log.Printf("Failed to update Pokemon: %v", err)
		return nil, fmt.Errorf("failed to update Pokemon")
	}

	return pokemonData, nil
}

func (pkb *Database) DeletePokemon(ctx context.Context, id string) (bool, error) {
	// Delete the pokemon by ID
	if err := pkb.Db.Delete(&Pokemon{}, id).Error; err != nil {
		return false, err
	}

	// Return success
	success := true
	return success, nil
}

func (pkb *Database) SearchByID(ctx context.Context, id string) (*model.Pokemon, error) {
	var pokemonData *model.Pokemon

	if err := pkb.Db.First(&pokemonData, id).Error; err != nil {
		fmt.Printf("Failed to find Pokemon by using : %v", err)
		return nil, fmt.Errorf("failed to find Pokemon by using id")
	}

	return pokemonData, nil
}

func (pkb Database) SearchByName(ctx context.Context, name string) ([]*model.Pokemon, error) {
	var pokemonList []*model.Pokemon
	if err := pkb.Db.Where("name LIKE ?", "%"+name+"%").Find(&pokemonList).Error; err != nil {
		fmt.Printf("Failed to find Pokemon by using : %v", err)
		return nil, fmt.Errorf("failed to find Pokemon for update")
	}
	return pokemonList, nil
}
