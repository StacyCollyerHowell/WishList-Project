package wishlistapi

var (
	wishLists []*Person
)

var (
	Items []*Item
)

type WishList struct {
	People []*Person
}

type Person struct {
	Name  string
	Items []*Item
}

type Item struct {
	Item      string
	Purchased bool
}

func SetWishList(w []*Person) {
	wishLists = w
}

func AddWishlist(wishList *Person) {
	wishLists = append(wishLists, wishList)
}

func ListWishlists() []*Person {
	return wishLists
}

func (plist *Person) AddItemToPerson(it *Item) {
	plist.Items = append(plist.Items, it)
}

func (y *Person) ShowItemList() []*Item {

	return y.Items
}

func (list *WishList) AddPersonToList(p *Person) {
	list.People = append(list.People, p)
}

// func (wlist *wishLists) AddPersonToWishList(p *Person) {
// 	wlist.Person = append(wlist.Person, p)
// }
