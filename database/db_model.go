package database

type Pokemon struct {
	ID          int `gorm:"AUTO_INCREMENT:1;PRIMARY_KEY;not null"`
	Name        string
	Description string
	Category    string
	Abilities   string
	Type        string
}

type CreatePokemonInput struct {
	Name        *string
	Description *string
	Category    *string
	Abilities   *string
	Type        *string
}

type UpdatePokemonInput struct {
	ID          *string
	Name        *string
	Description *string
	Category    *string
	Abilities   *string
	Type        *string
}

