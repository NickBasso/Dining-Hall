package food

import (
	"dining-hall/src/cookingApparatus"
)

type Food struct {
	Id               int16
	Name             string
	PreparationTime  int
	Complexity       byte
	CookingApparatus cookingApparatus.CookingApparatus
}

func GetFoodList () []Food {
	foodList := append(make([]Food, 0), Food{
		Id: 1,
		Name: "pizza",
		PreparationTime: 20,
		Complexity: 2,
		CookingApparatus: cookingApparatus.Oven},
	Food{
		Id: 2,
		Name: "salad",
		PreparationTime: 10,
		Complexity: 1,
		CookingApparatus: cookingApparatus.None},
	Food{
		Id: 3,
		Name: "zeama",
		PreparationTime: 7,
		Complexity: 1,
		CookingApparatus: cookingApparatus.Stove},	
	Food{
		Id: 4,
		Name: "Scallop Sashimi with Meyer Lemon Confit",
		PreparationTime: 32,
		Complexity: 3,
		CookingApparatus: cookingApparatus.None},
	Food{
		Id: 5,
		Name: "Island Duck with Mulberry Mustard",
		PreparationTime: 35,
		Complexity: 3,
		CookingApparatus: cookingApparatus.Oven},
	Food{
		Id: 6,
		Name: "Waffles",
		PreparationTime: 10,
		Complexity: 1,
		CookingApparatus: cookingApparatus.Stove},
	Food{
		Id: 7,
		Name: "Aubergine",
		PreparationTime: 20,
		Complexity: 2,
		CookingApparatus: cookingApparatus.None},
	Food{
		Id: 8,
		Name: "Lasagna",
		PreparationTime: 30,
		Complexity: 2,
		CookingApparatus: cookingApparatus.Oven},
	Food{
		Id: 9,
		Name: "Burger",
		PreparationTime: 15,
		Complexity: 1,
		CookingApparatus: cookingApparatus.Oven},
	Food{
		Id: 10,
		Name: "Gyros",
		PreparationTime: 15,
		Complexity: 1,
		CookingApparatus: cookingApparatus.None},)

	return foodList
}

func GetFoodMap() map[int16]Food {
	foodList := GetFoodList();
	foodMap := make(map[int16]Food)

	for i := 0; i < len(foodList); i++ {
		foodMap[foodList[i].Id] = foodList[i]
	}
	
	return foodMap
}