// the data structure representnig the user table.
package models

type User struct {
	ID        int    `bson:"id"`
	Username  string `bson:"username"`
	Password  string `bson:"password"`
	Email     string `bson:"email"`
	CreatedAt string `bson:"createdAt`
	UpdatedAt string `bson:"updatedAt"`
}
