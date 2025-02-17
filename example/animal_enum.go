// Code generated by go-enum DO NOT EDIT.
// Version: example
// Revision: example
// Build Date: example
// Built By: example

package example

import (
	"fmt"
)

const (
	// AnimalCat is a Animal of type Cat.
	AnimalCat Animal = iota
	// AnimalDog is a Animal of type Dog.
	AnimalDog
	// AnimalFish is a Animal of type Fish.
	AnimalFish
	// AnimalFishPlusPlus is a Animal of type Fish++.
	AnimalFishPlusPlus
	// AnimalFishSharp is a Animal of type Fish#.
	AnimalFishSharp
)

const _AnimalName = "CatDogFishFish++Fish#"

var _AnimalMap = map[Animal]string{
	AnimalCat:          _AnimalName[0:3],
	AnimalDog:          _AnimalName[3:6],
	AnimalFish:         _AnimalName[6:10],
	AnimalFishPlusPlus: _AnimalName[10:16],
	AnimalFishSharp:    _AnimalName[16:21],
}

// String implements the Stringer interface.
func (x Animal) String() string {
	if str, ok := _AnimalMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Animal(%d)", x)
}

var _AnimalValue = map[string]Animal{
	_AnimalName[0:3]:   AnimalCat,
	_AnimalName[3:6]:   AnimalDog,
	_AnimalName[6:10]:  AnimalFish,
	_AnimalName[10:16]: AnimalFishPlusPlus,
	_AnimalName[16:21]: AnimalFishSharp,
}

// ParseAnimal attempts to convert a string to a Animal.
func ParseAnimal(name string) (Animal, error) {
	if x, ok := _AnimalValue[name]; ok {
		return x, nil
	}
	return Animal(0), fmt.Errorf("%s is not a valid Animal", name)
}
