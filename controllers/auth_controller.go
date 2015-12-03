package controllers

import (
	"encoding/json"
	"net/http"
	"sample/services"
	"sample/services/models"
	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context) {
//func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := services.Login(requestUser)
	println(responseStatus)
//	w.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"token": token})
//	w.WriteHeader(responseStatus)
//	w.Write(token)
}

func RefreshToken(c *gin.Context) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&requestUser)
	c.JSON(http.StatusOK, gin.H{"token": services.RefreshToken(requestUser)})
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(services.RefreshToken(requestUser))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := services.Logout(r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
