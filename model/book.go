package model

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

var StoreBooks = []Book{
	{
		ID:    "1",
		Title: "The Hobbit",
		Author: Author{
			ID:       "2",
			FistName: "John",
			LastName: "Doe",
		},
	},
	{
		ID:    "2",
		Title: "Moon Landing",
		Author: Author{
			ID:       "1",
			FistName: "Kim",
			LastName: "Beckley",
		},
	},
}
