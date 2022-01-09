package prototype

import (
	"errors"
	"fmt"
)

// ShirtCloner ...
type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	// White ...
	White = 1
	// Black ...
	Black = 2
	// Blue ...
	Blue = 3
)

// GetShirtsCloner ...
func GetShirtsCloner() ShirtCloner {
	return nil
}

// ShirtsCache ...
type ShirtsCache struct{}

// GetClone ...
func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

// ItemInfoGetter ...
type ItemInfoGetter interface {
	GetInfo() string
}

// ShirtColor ...
type ShirtColor byte

// Shirt ...
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

// GetInfo ...
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}
var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}
var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

// GetPrice ...
func (s *Shirt) GetPrice() float32 {
	return s.Price
}
