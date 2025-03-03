package main

import "fmt"

func main() {
	fmt.Println("Welcome!")

	// hobbies := [4]string{"cars", "programming", "hiking", "walking"}
	// hobbies_copy := hobbies
	// fmt.Println("Hobbies: ", hobbies)
	// fmt.Println("Copy of Hobbies: ", hobbies_copy)
	// start_index := 2
	// end_index := 4

	// hobbies_subset := hobbies[start_index:end_index]
	// fmt.Println("Hobbies subset: ", hobbies_subset)

	// fmt.Println("First element of Hobbies: ", hobbies[0])
	// fmt.Println("Second and third element of Hobbies: ", hobbies[1:3])

	// favorite_hobbies := hobbies[0:2]
	// fmt.Println("Favorite hobbies: ", favorite_hobbies)

	// more_favorite_hobbies []string = favorite_hobbies[1:3]
	// fmt.Println("More Favorite hobbies: ", more_favorite_hobbies)

	// goals := []string{"job", "debt"}
	// fmt.Println("Goals: ", goals)
	// // change goal
	// goals[1] = "rx7"
	// fmt.Println("Goals: ", goals)

	// var updated_goals []string = goals
	// fmt.Println("Updated Goals: ", updated_goals)

	// updated_goals = append(updated_goals, "535xi", "528i")

	// fmt.Println("Updated Goals: ", updated_goals)

	// fmt.Println("Length updated goals: ", len(updated_goals))
	// fmt.Println("capacity updated goals: ", cap(updated_goals))

	// type product struct {
	// 	title string
	// 	id    int
	// 	price float64
	// }

	// productlist := []product{
	// 	{
	// 		"Book1",
	// 		1,
	// 		10.99,
	// 	},
	// 	{
	// 		"Book2",
	// 		2,
	// 		12.99,
	// 	},
	// }

	// fmt.Println("Product list: ", productlist)

	// new_product := product{"book3", 3, 20.99}

	// productlist = append(productlist, new_product)

	// fmt.Println("Updated Product list: ", productlist)

	people := make([]string, 1, 10)

	people[0] = "neil"
	people = append(people, "Shawn")
	people = append(people, "Frank")
	people = append(people, "Julie")
	people = append(people, "Dave")
	people = append(people, "Jack")
	people = append(people, "Pheebe")
	people = append(people, "Han")

	// for i, v := range people {
	// 	fmt.Println(i, v)
	// }

	var lenght int = len(people)

	for i := 1; i < lenght; i = i + 2 {
		fmt.Println(i, people[i])
	}

}
