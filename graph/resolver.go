package graph

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/TanyaEIEI/pokedex/graph/model"
	"gorm.io/gorm"
)

type Resolver struct {
}

type CustomContext struct {
	Database *gorm.DB
}

var customContextKey string = "CUSTOM_CONTEXT"

func CreateContext(args *CustomContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customContext := &CustomContext{
			Database: args.Database,
		}
		requestWithCtx := r.WithContext(context.WithValue(r.Context(), customContextKey, customContext))
		next.ServeHTTP(w, requestWithCtx)
	})
}

func GetContext(ctx context.Context) *CustomContext {
	customContext, ok := ctx.Value(customContextKey).(*CustomContext)
	if !ok {
		return nil
	}
	return customContext
}

func getCurrentId(list interface{}) (int, error) {
	switch list.(type) {
	case []*model.PokemonAbility:
		newItem := list.([]*model.PokemonAbility)
		return strconv.Atoi(newItem[len(newItem)-1].ID)
	case []*model.PokemonType:
		newItem := list.([]*model.PokemonType)
		return strconv.Atoi(newItem[len(newItem)-1].ID)
	case []*model.PokemonCategory:
		newItem := list.([]*model.PokemonCategory)
		return strconv.Atoi(newItem[len(newItem)-1].ID)
	case []*model.Pokemon:
		newItem := list.([]*model.Pokemon)
		return strconv.Atoi(newItem[len(newItem)-1].ID)
	default:
		return 0, errors.New("unsupported type")
	}
}
