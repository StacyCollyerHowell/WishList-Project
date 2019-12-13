package main

//Change all the functions that reference the methods and functions in wishList.go file to reference the new functions in Model

//copy whole DB and config file to get data base connected use slice flie to get one part connected, one part at a time

import (
	"fmt"
	"os"
	"strconv"

	"github.com/StacyCollyerHowell/WishList-Project/db"
	"github.com/StacyCollyerHowell/WishList-Project/storage"
	_ "github.com/go-sql-driver/mysql"

	"github.com/StacyCollyerHowell/WishList-Project/wishlistapi"
	"github.com/manifoldco/promptui"
)

const (
	addPerson  = "Add Name"
	addItem    = "Add Item"
	showPeople = "Show List of People"
	// viewItems  = "Show Item List"
)

var wishListService *wishlistapi.WishListService

func main() {

	db, err := db.ConnectDatabase("wishlist_db.config")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	wishListService = wishlistapi.NewService(db)

	for {
		fmt.Println()

		prompt := promptui.Select{
			Label: "Choose one",
			Items: []string{
				addPerson,
				showPeople,
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case addPerson:
			err := addPersonPrompt()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

		case showPeople:
			err := ViewPerson()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			// case viewItems:
			// 	err := viewPersonList()
			// 	if err != nil {
			// 		fmt.Printf("Prompt failed %v\n", err)
			// 		return
			// 	}
		}
		fmt.Println(result)
	}
}

func addPersonPrompt() error {
	for {

		namePrompt := promptui.Prompt{
			Label: "Name",
		}
		name, err := namePrompt.Run()
		if err != nil {
			return err
		}

		newPerson := &wishlistapi.Person{
			Name: name,
		}

		wishListService.AddPerson(newPerson.Name)

		fmt.Println("Added Name", newPerson)

		item, err := promptForItem()
		if err != nil {
			return err
		}

		newPerson.AddItemToPerson(item)

		fmt.Println("Added Item", item)

		err = storage.Save()
		if err != nil {
			return err
		}

		//add struct to a person
		//add a for loop to get multiple items in

		return nil
	}
}

func promptForItem() (*wishlistapi.Item, error) {

	for {
		itemPrompt := promptui.Prompt{
			Label: "Item",
		}

		item, err := itemPrompt.Run()
		if err != nil {
			return nil, err
		}

		purchased, err := promptForBool("Purchased?")
		if err != nil {
			return nil, err
		}

		newItem := &wishlistapi.Item{
			Item:      item,
			Purchased: purchased,
		}

		done, err := promptForBool("Done Adding Items?")
		if err != nil {
			return nil, err
		}

		if done {
			return newItem, nil
		}
	}

	return nil, nil
}

func promptForBool(label string) (bool, error) {

	purchasedPrompt := promptui.Prompt{
		Label: label,
	}

	purchasedStr, err := purchasedPrompt.Run()
	if err != nil {
		return false, err
	}

	purchased, err := strconv.ParseBool(purchasedStr)
	if err != nil {
		return false, err
	}

	return purchased, nil
}

func ViewPerson() error {

	availablePeople, err := wishListService.ListPeople()
	if err != nil {
		return err
	}

	if len(availablePeople) == 0 {
		fmt.Println("No names to select!")
		return nil
	}

	var options []string
	for _, wishList := range availablePeople {
		options = append(options, wishList.Name)
	}

	selectNamePrompt := promptui.Select{
		Label: "Select Name",
		Items: options,
	}

	chosenIndex, _, err := selectNamePrompt.Run()
	if err != nil {
		return err
	}

	chosenPerson := availablePeople[chosenIndex]

	viewItems(chosenPerson)

	if err != nil {
		return err
	}

	return nil
}

func viewItems(chosenPerson *wishlistapi.Person) error {

	availableItems, err := wishListService.ShowItemList(chosenPerson.ID)

	if err != nil {
		return err
	}
	if len(availableItems) == 0 {
		fmt.Println("No items to select!")
		return nil
	}

	fmt.Println(availableItems)

	var options []string
	for _, personList := range availableItems {
		if personList.Purchased == true {
			options = append(options, personList.Item+" ✓")
		} else {
			options = append(options, personList.Item+" ☐")

		}
	}
	selectItemPrompt := promptui.Select{
		Label: "Select item",
		Items: options,
	}

	chosenIndex, _, err := selectItemPrompt.Run()
	if err != nil {
		return err
	}

	chosenItem := availableItems[chosenIndex]

	if err != nil {
		return err
	}

	promptForBool(chosenItem.Item)

	chosenItem.Purchased = !chosenItem.Purchased

	wishListService.UpdateItem(chosenItem.ID, chosenItem.Purchased)

	return nil
}

//show list of items when select person
//let them select if they have bought the item or not

//do the same thing as show people to show list of items
