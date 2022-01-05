package decorator

import (
	"errors"
	"fmt"
)

type IngredientAdd interface {
	AddIngredient() (string, error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

type Meat struct {
	Ingredient IngredientAdd
}

type Onion struct {
	Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	// return "", errors.New("Not implemented yet")
	return "Pizza with the following ingredients:", nil
}

func (m *Meat) AddIngredient() (string, error) {
	// return "", errors.New("Not implemented yet")
	if m.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Meat")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "meat"), nil
}

func (o *Onion) AddIngredient() (string, error) {
	// return "", errors.New("Not implemented yet")
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "onion"), nil
}
