package food

import (
	"dininghall/src/components/constants"
	"dininghall/src/components/types/apparatus"
)

type Food struct {
	ID              int
	Name            string
	PreparationTime int
	Complexity      int
	Apparatus       apparatus.Apparatus
}

func GetMenuArray() []Food {
	foodList := append(make([]Food, 0), Food{
		ID:              1,
		Name:            "pizza",
		PreparationTime: 20,
		Complexity:      2,
		Apparatus:       apparatus.Oven},
		Food{
			ID:              2,
			Name:            "salad",
			PreparationTime: 10,
			Complexity:      1,
			Apparatus:       apparatus.None},
		Food{
			ID:              3,
			Name:            "zeama",
			PreparationTime: 7,
			Complexity:      1,
			Apparatus:       apparatus.Stove},
		Food{
			ID:              4,
			Name:            "Scallop Sashimi with Meyer Lemon Confit",
			PreparationTime: 32,
			Complexity:      3,
			Apparatus:       apparatus.None},
		Food{
			ID:              5,
			Name:            "Island Duck with Mulberry Mustard",
			PreparationTime: 35,
			Complexity:      3,
			Apparatus:       apparatus.Oven},
		Food{
			ID:              6,
			Name:            "Waffles",
			PreparationTime: 10,
			Complexity:      1,
			Apparatus:       apparatus.Stove},
		Food{
			ID:              7,
			Name:            "Aubergine",
			PreparationTime: 20,
			Complexity:      2,
			Apparatus:       apparatus.None},
		Food{
			ID:              8,
			Name:            "Lasagna",
			PreparationTime: 30,
			Complexity:      2,
			Apparatus:       apparatus.Oven},
		Food{
			ID:              9,
			Name:            "Burger",
			PreparationTime: 15,
			Complexity:      1,
			Apparatus:       apparatus.Oven},
		Food{
			ID:              10,
			Name:            "Gyros",
			PreparationTime: 15,
			Complexity:      1,
			Apparatus:       apparatus.None})

	return foodList
}

func GetMenuMap () map[int]Food {
	menuArray := GetMenuArray()
	menuMap := make(map[int]Food, constants.MenuCount)

	for i := 0; i < len(menuArray); i++ {
		menuMap[menuArray[i].ID] = menuArray[i] 
	}

	return menuMap
}