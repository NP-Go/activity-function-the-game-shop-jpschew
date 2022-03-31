package main

func main() {
	//Create objects
	game1 := game{"Minecraft", 5}
	game2 := game{"World of Warcraft", 19}
	game3 := game{"Elite Dangerours", 54}

	var gameList list
	gameList = append(gameList, game1, game2, game3)
	// fmt.Println(list)
	gameList.printContent()

	for _, item := range gameList {
		item.printPrice()
	}
}
