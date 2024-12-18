package models

// data structure representing the cart table.

type Cart struct {
	ID       string   `bson:"id"`
	UserID   string   `bson:"UserId"`
	Products []string `bson:"products"`
}
