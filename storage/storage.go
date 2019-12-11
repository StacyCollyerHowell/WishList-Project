package storage

import (
	"encoding/json"
	"io/ioutil"

	"github.com/StacyCollyerHowell/WishList-Project/wishlistapi"
)

const filename = "wishlist.json"

func Load() error {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var savedWishList []*wishlistapi.Person
	err = json.Unmarshal(fileContents, &savedWishList)
	if err != nil {
		return err
	}

	wishlistapi.SetWishList(savedWishList)

	return nil
}

func Save() error {
	wishList := wishlistapi.ListWishlists()

	wishListBytes, err := json.MarshalIndent(wishList, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, wishListBytes, 0775)
	if err != nil {
		return err
	}

	return nil
}
