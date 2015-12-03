package main

import (
	"fmt"
	"runtime"
	"sample/controllers"
	"sample/core/authentication"
	"sample/settings"
	"sample/user_controllers"
	"time"

	"gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
)

func main() {
	settings.Init()
	ConfigRuntime()
	StartWorkers()
	StartGin()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
	fmt.Printf("Running with %d1 CPUs\n", nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartWorkers() {
	go statsWorker()
}

func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	gin.Logger()
	router.Use(rateLimit, gin.Recovery())
	router.Use(gin.Logger())
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
	router.GET("/", MyBenchLogger(), index)
	router.GET("/auth", authentication.RequireTokenAuthentication(), index)
	router.POST("/test", controllers.Login)
	router.GET("/room/:name", roomGET)
	router.POST("/room-post/:roomid", roomPOST)
	router.GET("/stream/:roomid", streamRoom)

	//mongodb user create
	uc := user_controllers.NewUserController(getSession())
	router.GET("/user", uc.GetUser)
	router.GET("/message", uc.GetMessage)
	router.POST("/message", uc.CreateMessage)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.RemoveUser)

	router.Run(":5001")

}

func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		fmt.Println(start, path)
		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)
		fmt.Println(latency)

	}
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
