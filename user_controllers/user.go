package user_controllers

import (
	"fmt"
	"sample/models"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	UserController struct {
		session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}




func (uc UserController) GetUser(c *gin.Context) {
	name := c.Request.Header.Get("name")
	id := c.Request.Header.Get("id")
	gender := c.Request.Header.Get("gender")

	if !bson.IsObjectIdHex(id) {
		c.Writer.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)
	fmt.Println(oid)
	u := models.User{}
	fmt.Println(id)
	fmt.Println(name)
	//	if err := uc.session.DB("go_rest_tutorial").C("users").FindId(oid).One(&u); err != nil {
	if err := uc.session.DB("go_rest_tutorial").C("users").Find(bson.M{"gender": gender}).One(&u); err != nil {
		c.Writer.WriteHeader(404)
		return
	}
	c.JSON(200, u)
}


func (uc UserController) CreateUser(c *gin.Context) {
	u := models.User{}
	c.BindJSON(&u)
	u.Id = bson.NewObjectId()
	uc.session.DB("go_rest_tutorial").C("users").Insert(u)
	c.JSON(201, u)
}

func (uc UserController) RemoveUser(c *gin.Context) {
	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		c.Writer.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("go_rest_tutorial").C("users").RemoveId(oid); err != nil {
		c.Writer.WriteHeader(404)
		return
	}
	c.Writer.WriteHeader(200)
}












func (uc UserController) GetMessage(c *gin.Context) {
	theme := c.Request.Header.Get("theme")
	u := models.Message{}
	if err := uc.session.DB("go_rest_tutorial").C("messages").Find(bson.M{"theme": theme}).One(&u); err != nil {
		c.Writer.WriteHeader(404)
		return
	}
	c.JSON(200, u)
}

func (uc UserController) CreateMessage(c *gin.Context) {
	u := models.Message{}
	c.BindJSON(&u)
	u.Id = bson.NewObjectId()
	uc.session.DB("go_rest_tutorial").C("messages").Insert(u)
	c.JSON(201, u)
}