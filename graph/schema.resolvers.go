package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"strconv"

	"github.com/TanyaEIEI/pokedex/database"
	"github.com/TanyaEIEI/pokedex/graph/model"
)

// CreatePokemon is the resolver for the createPokemon field.
func (r *mutationResolver) CreatePokemon(ctx context.Context, input model.PokemonInput) (*model.Pokemon, error) {
	newpokemon, err := r.Pokedex.CreatePokemon(ctx, database.Pokemon{
		Name:        *input.Name,
		Description: *input.Description,
		Category:    *input.Category,
		Abilities:   *input.Abilities,
		Type:        *input.Type,
	})
	if err != nil {
		return nil, err
	}

	var responsePokemon model.Pokemon
	responsePokemon.ID = strconv.Itoa(newpokemon.ID)
	responsePokemon.Name = newpokemon.Name
	responsePokemon.Category = newpokemon.Category
	responsePokemon.Description = newpokemon.Description
	responsePokemon.Abilities = &newpokemon.Abilities
	responsePokemon.Type = &newpokemon.Type

	return &responsePokemon, nil
}

// UpdatePokemon is the resolver for the updatePokemon field.
func (r *mutationResolver) UpdatePokemon(ctx context.Context, input model.PokemonInput) (*model.Pokemon, error) {
	updatePokemon, err := r.Pokedex.UpdatePokemon(ctx, database.UpdatePokemonInput{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Abilities:   input.Abilities,
		Type:        input.Type,
	})
	if err != nil {
		return nil, err
	}
	var responsePokemon model.Pokemon
	responsePokemon.ID = strconv.Itoa(updatePokemon.ID)
	responsePokemon.Name = updatePokemon.Name
	responsePokemon.Category = updatePokemon.Category
	responsePokemon.Description = updatePokemon.Description
	responsePokemon.Abilities = &updatePokemon.Abilities
	responsePokemon.Type = &updatePokemon.Type

	return &responsePokemon, nil
}

// DeletePokemon is the resolver for the deletePokemon field.
func (r *mutationResolver) DeletePokemon(ctx context.Context, id string) (bool, error) {
	return r.Pokedex.DeletePokemon(ctx, id)
}

// SearchPokemonByID is the resolver for the searchPokemonById field.
func (r *queryResolver) SearchPokemonByID(ctx context.Context, id string) (*model.Pokemon, error) {
	return r.Pokedex.SearchByID(ctx, id)
}

// SearchPokemonByName is the resolver for the searchPokemonByName field.
func (r *queryResolver) SearchPokemonByName(ctx context.Context, name string) ([]*model.Pokemon, error) {
	return r.Pokedex.SearchByName(ctx, name)
}

// Pokemons is the resolver for the pokemons field.
func (r *queryResolver) Pokemons(ctx context.Context) ([]*model.Pokemon, error) {
	return r.Pokedex.ListPokemon(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
