package main

import (
	"fmt"
	"strconv"

	"github.com/StacyCollyerHowell/WishList-Project/storage"

	"github.com/StacyCollyerHowell/WishList-Project/wishlistapi"
	"github.com/manifoldco/promptui"
)

const (
	addPerson  = "Add Name"
	addItem    = "Add Item"
	showPeople = "Show List of People"
	// viewItems  = "Show Item List"
)

func main() {

	err := storage.Load()
	if err != nil {
		fmt.Println("Error Loading WishLists from file", err)
	}

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

		wishlistapi.AddWishlist(newPerson)

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

	availablePeople := wishlistapi.ListWishlists()

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

	//chosenPerson.viewItems(chosenPerson)

	//ViewItems(chosenPerson)

	viewItems(chosenPerson)

	// if err != nil {
	// 	return err
	// }

	if err != nil {
		return err
	}

	return nil
}

func viewItems(chosenPerson *wishlistapi.Person) error {

	availableItems := chosenPerson.ShowItemList()

	if len(availableItems) == 0 {
		fmt.Println("No items to select!")
		return nil
	}

	fmt.Println(availableItems)

	var options []string
	for _, personList := range availableItems {
		options = append(options, personList.Item)
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

	//implemnt viewItems function similar to view person

	promptForBool(chosenItem.Item)

	return nil
}

//show list of items when select person
//let them select if they have bought the item or not

//do the same thing as show people to show list of items
