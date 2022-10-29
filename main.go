package main

import (
	"fmt"

	"github.com/fuadop/altschool-go-exam/shop"
)

func main() {
	item1 := shop.NewCar("Ford", "Mustang", "red")
	item2 := shop.NewCar("Ford", "Mustang", "blue")
	item3 := shop.NewCar("Dodge", "Challenger", "orange")

	store := shop.NewStore()
	store.AddItem(item1, 20000)
	store.AddItem(item1, 20000)
	store.AddItem(item1, 20000)

	store.AddItem(item2, 30000)
	store.AddItem(item2, 30000)

	store.AddItem(item3, 40000) // rare collection haha

	fmt.Println("Initial products")
	fmt.Println()
	store.DisplayItems()

	fmt.Println("Selling some items")
	fmt.Println()
	ok, err := store.SellItem(item1)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
	}

	if ok {
		fmt.Println("Sold =>", item1.GetName())
		fmt.Println()
	}

	store.DisplayOutflow()

	fmt.Println("Sell another item")
	fmt.Println()
	ok, err = store.SellItem(item2)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
	}

	if ok {
		fmt.Println("Sold =>", item1.GetName())
		fmt.Println()
	}

	store.DisplayOutflow()

	fmt.Println("Done with preview. Thanks for previewing")
}
