package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}

type Message struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Theme  string        `json:"theme" bson:"theme"`
	Detail string        `json:"detail" bson:"detail"`
}
