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
	insertPersonQuery = "INSERT INTO person (person_name) VALUES (?);"

	updateItemQuery = "UPDATE items SET purchased = ? where id = ?"

	selectPersonQuery = "SELECT id, person_name FROM person"

	selectItemQuery = "SELECT id, item_name, purchased FROM items WHERE person_id = ?"

	insertItemQuery = "INSERT INTO items (item_name, purchased, person_id) VALUES (?, ?, ?)"
)

func (a *WishListService) AddPerson(personName string) error {
	_, err := a.db.Exec(insertPersonQuery, personName)
	if err != nil {
		return err
	}

	return nil
}

func (a *WishListService) ListPeople() ([]*Person, error) {
	rows, err := a.db.Query(selectPersonQuery)
	if err != nil {
		return nil, err
	}

	var people []*Person
	for rows.Next() {
		var person Person

		err := rows.Scan(
			&person.ID,
			&person.Name,
		)
		if err != nil {
			return nil, err
		}

		people = append(people, &person)
	}

	return people, nil
}

func (a *WishListService) ShowItemList(PersonID int) ([]*Item, error) {
	rows, err := a.db.Query(selectItemQuery, PersonID)
	if err != nil {
		return nil, err
	}

	var items []*Item
	for rows.Next() {
		var item Item

		err := rows.Scan(
			&item.ID,
			&item.Item,
			&item.Purchased,
			//&item.PersonID,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	return items, nil
}

func (a *WishListService) AddItemsToPerson(itemName string, purchased bool, personID int) error {
	_, err := a.db.Exec(insertItemQuery, itemName, purchased, personID)
	if err != nil {
		return err
	}

	return nil
}

func (a *WishListService) UpdateItem(itemID int, purchased bool) error {
	_, err := a.db.Exec(updateItemQuery, purchased, itemID)
	if err != nil {
		return err
	}

	return nil
}
