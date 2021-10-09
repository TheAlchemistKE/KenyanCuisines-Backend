package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/TheAlchemistKE/KenyanCuisines-Backend/graph/generated"
	"github.com/TheAlchemistKE/KenyanCuisines-Backend/graph/model"
)

func (r *mutationResolver) CreateRecipe(ctx context.Context, input model.NewRecipe) (*model.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddRecipeInstruction(ctx context.Context, input *model.RecipeInstruction) (*model.Instruction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddRecipeIngredient(ctx context.Context, input *model.RecipeIngredient) (*model.Ingredient, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Recipes(ctx context.Context) ([]*model.Recipe, error) {
	var recipes []*model.Recipe
	var ingredients []*model.Ingredient
	var instructions []*model.Instruction

	var ingr = model.Ingredient{
		Name:     "Parsley",
		Quantity: "One bunch",
		Recipe:   nil,
	}
	ingredients = append(ingredients, &ingr)

	var ingr1 = model.Ingredient{
		Name:     "Tomato",
		Quantity: "Two",
		Recipe:   nil,
	}
	ingredients = append(ingredients, &ingr1)

	var instruction1 = model.Instruction{
		Info:   "Fry tomatoes then add parsley.",
		Recipe: nil,
	}
	instructions = append(instructions, &instruction1)

	var dummyRecipe = model.Recipe{
		Name: "Recipe1",
		Ingredients: ingredients,
		Instructions: instructions,
	}
	recipes = append(recipes, &dummyRecipe)
	return recipes, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
