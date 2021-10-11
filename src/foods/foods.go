package foods

import "dining-hall/src/cookingApparatus"

type Foods struct {
	id               string
	name             string
	preparationTime  int
	complexity       byte
	cookingApparatus cookingApparatus.CookingApparatus
}