package database

import (
	"context"
	"fmt"

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
	if input.Name != nil {
		pokemonData.Name = *input.Name
	}
	if input.Description != nil {
		pokemonData.Description = *input.Description
	}
	if input.Category != nil {
		pokemonData.Category = *input.Category
	}
	if input.Abilities != nil {
		pokemonData.Abilities = *input.Abilities
	}
	if input.Type != nil {
		pokemonData.Type = *input.Type
	}

	// Update other fields as needed

	if err := pkb.Db.Save(&pokemonData).Error; err != nil {
		fmt.Printf("Failed to update Pokemon: %v", err)
		return nil, fmt.Errorf("failed to update Pokemon")
	}

	return pokemonData, nil
}

func (pkb *Database) DeletePokemon(ctx context.Context, id string) (bool, error) {
	// Check if the Pokemon exists
	var existingPokemon Pokemon
	if err := pkb.Db.First(&existingPokemon, "id = ?", id).Error; err != nil {
		return false, err
	}

	if err := pkb.Db.Delete(&Pokemon{}, id).Error; err != nil {
		return false, err
	}

	return true, nil
}


func (pkb *Database) SearchByID(ctx context.Context, id string) (*model.Pokemon, error) {
	var pokemonData *model.Pokemon

	if err := pkb.Db.First(&pokemonData, id).Error; err != nil {
		fmt.Printf("Failed to find this pokemon's id : %v", err)
		return nil, fmt.Errorf("Failed to find this pokemon's id ")
	}

	return pokemonData, nil
}

func (pkb Database) SearchByName(ctx context.Context, name string) ([]*model.Pokemon, error) {
	var pokemonList []*model.Pokemon
	if err := pkb.Db.Where("name LIKE ?", "%"+name+"%").Find(&pokemonList).Error; err != nil {
		fmt.Printf("failed to find this pokemon's name: %v", err)
		return nil, fmt.Errorf("failed to find this pokemon's name")
	}
	return pokemonList, nil
}

// Additional Func
func CheckInput(input model.PokemonInput) error {
	if input.Name == nil {
		return fmt.Errorf("name must not be null")
	}
	if input.Description == nil {
		return fmt.Errorf("description must not be null")
	}
	if input.Category == nil {
		return fmt.Errorf("category must not be null")
	}
	if input.Abilities == nil {
		return fmt.Errorf("abilities must not be null")
	}
	if input.Type == nil {
		return fmt.Errorf("type must not be null")
	}

	return nil
}
