package wishlistapi

import "database/sql"

type WishListService struct {
	db *sql.DB
}

func NewService(db *sql.DB) *WishListService {
	return &WishListService{
		db: db,
	}
}

const (
	insertPersonQuery = "INSERT INTO person (id, person_name) VALUES (?, ?);"

	selectPersonQuery = "SELECT id, person_name, FROM person"

	selectItemQuery = "SELECT id, item_name, purchased FROM items WHERE person_id = ?"

	insertItemQuery = "INSERT INTO items (id, item_name, purchased, person_id) VALUES (?, ?, ?,?)"
)

func (a *WishListService) AddWishlist(personName string) error {
	_, err := a.db.Exec(insertPersonQuery, Person.Name)
	if err != nil {
		return err
	}

	return nil
}

func (a *WishListService) ListWishlists() ([]Person, error) {
	rows, err := a.db.Query(selectPersonQuery)
	if err != nil {
		return nil, err
	}

	var people []Person
	for rows.Next() {
		var person Person

		err := rows.Scan(
			&person.ID,
			&person.Name,
			&person.ItemID,
		)
		if err != nil {
			return nil, err
		}

		people = append(people, person)
	}

	return people, nil
}

func (a *WishListService) ShowItemList(PersonID int) ([]Items, error) {
	rows, err := a.db.Query(selectItemQuery, ItemID)
	if err != nil {
		return nil, err
	}

	var items []Item
	for rows.Next() {
		var item Item

		err := rows.Scan(
			&item.ID,
			&item.Item,
			&item.Purchased,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (a *WishListService) AddItemsToPerson(itemID int, personID int) error {
	_, err := a.db.Exec(insertItemQuery, itemID, personID)
	if err != nil {
		return err
	}

	return nil
}
