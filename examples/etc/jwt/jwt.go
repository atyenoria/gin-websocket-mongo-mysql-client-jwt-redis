package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
)

const myToken = "sample"

type token struct {
	Header map[string]string
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}

var pool = newPool()

func redistest() {
	c := pool.Get()
	defer c.Close()

	test2, _ := c.Do("HGETALL", "test:2")
	fmt.Println(test2)
}

func main() {

	//redis
	c := pool.Get()
	defer c.Close()
//	redistest()
	redistest, _ := c.Do("SET","aasdfa","aaasdfa")
	fmt.Println(redistest)

	// New web token.
	token := jwt.New(jwt.SigningMethodHS256)

	// Set a header and a claim
	token.Header["kid"] = "JWT"
	token.Header["test"] = "samle"
	token.Header["te"] = "samleasdf"
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString([]byte(myToken))
	fmt.Println(t)

	dtoken, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		//		return myLookupKey(token.Header["kid"])
		return []byte(myToken), nil
	})

	if err == nil && dtoken.Valid {
		fmt.Println("ok")
		fmt.Println(dtoken)
	} else {
		fmt.Println("miss")
	}

}

func myLookupKey(a string) int {
	fmt.Println(a)
	return 1
}
