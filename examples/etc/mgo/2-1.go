package main

import (
	// native packages
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// 3rd packages
	"labix.org/v2/mgo/bson"
)


// Organizations数据模型结构
type Organizations struct {
	Name         string `json:"name"`
	Area         string `json:"area"`
	Type         string `json:"type"`
	CustomerType string `json:"customertype"`
	State        string `json:"state"`
	LastUpdated  string `json:"lastupdated"`
	Assigned     string `json:"assigned"`

	Address `json:"address"`
}

// for GET /Organizations
func (u *Organizations) Index(rw http.ResponseWriter, req *http.Request) {
	var org Organizations
	
	conditions := bson.M{"_id": bson.M{"$exist": 1}}
	
	result, _ := org.Retrieve(conditions)
	fmt.Fprint(rw, result)
}

// Organizations Find
func (o *Organizations) Retrieve(conditions map[string]interface{}) ([]Organizations, error) {
	session, err := getSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	collection := session.DB(DATABASE).C("organizations")

	result := []Organizations{}

	err = collection.Find(conditions).All(&result)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}


func (u Organizations) Index(rw http.ResponseWriter, req *http.Request) {
	notice := "Organizations#Index:" + req.Method + req.URL.String()
	log.Println(notice)
	values := req.URL.Query()

	var org model.Organizations
	conditions := bson.M{"_id": bson.M{"$exists": true}}

	if values.Get("name") != "" {
		conditions["name"] = values.Get("name")
	}
	if values.Get("area") != "" {
		conditions["area"] = values.Get("area")
	}
	if values.Get("type") != "" {
		conditions["type"] = values.Get("type")
	}

	result, _ := org.Retrieve(conditions)
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(result)
	if err != nil {
		log.Println(err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
	}
}