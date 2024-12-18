package models

// data structure representing payment information.

type Checkout struct {
	ID        string  `bson:"id"`
	UserID    string  `bson:"userId"`
	Total     float64 `bson:"total"`
	CreatedAt string  `bson:"createdAt"`
}
