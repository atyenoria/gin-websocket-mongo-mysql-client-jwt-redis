package main

import (
    "fmt"
    mgo "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

// Personに共通なものはここにまとめたいですよね
type Person struct {
    ID    bson.ObjectId `bson:"_id"`
    Name  string        `bson:"name"`
    Level int           `bson:"level"`
}

func (p *Person) Greet() string {
    return fmt.Sprintf("My name is %s.", p.Name)
}

// 派生したstructもつくりたいですよね
type Admin struct {
//    Person
	Person      `bson:",inline"`
    Password    string   `bson:"password"`
    Permissions []string `bson:"permissions"`
}

func main() {
    session, _ := mgo.Dial("mongodb://localhost/test1")
    defer session.Close()
    db := session.DB("test")

    admin := &Admin{
        Person{
            bson.NewObjectId(),
            "otiai10",
            10,
        },
        "hoge",
        []string{"write", "read"},
    }

    db.C("people").Insert(admin)
}