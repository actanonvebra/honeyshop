// data stucture representing log records.
package models

type Log struct {
	ID        string `bson:"id"`
	Timestamp string `bson:"timestamp"`
	Level     string `bson:"level"`
	Message   string `bson:"message"`
}
