package models

import "gopkg.in/mgo.v2/bson"

type Contest struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	Link        string        `json:"link" bson:"link"`
}
