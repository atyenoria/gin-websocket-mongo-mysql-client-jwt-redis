package main

import (
	"fmt"
	"runtime"
	"sample/controllers"
	"sample/core/authentication"
	"sample/settings"
	"time"

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
	router.GET("/room/:roomid", roomGET)
	router.POST("/room-post/:roomid", roomPOST)
	router.GET("/stream/:roomid", streamRoom)
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
