package main //Marks file as standalone executable program (not reusable). All programs must start with a main package to run

import (
	"log"
	"net/http"
	"time"

	api "github.com/fareeza05/Code-Review-Assistant/internal/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
) /*Pulls in Go's standard libraries
"fmt" //For formatted output like print
"log" //For logging messages to terminal
"net/http" //To build HTTP servers
*/

// Every Go program starts execution at main() function

func main() {
	/*
		Setting up a router with Gin
	*/

	//Set up router with default middleware (logger + recovery)
	router := gin.Default()

	/*
		Setting up CORS Middleware
		router.Use -> how you register middlware in Gin -> Gets passed a function that runs before request handler
		cors.New -> what runs before request handler -> creates middleware instance with custom config
		cors.Config -> holds all CORS settings, is a struct

	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},                            // dictates which origins to allow requests from, * indicates all
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Lists HTTP methods we are okay with clients from other origins using
		AllowHeaders: []string{"Origin", "Content-Type"},       // Requests we're allowing from cross-origin requests,
		//'Origin' needed for CORS, Content-Type needed if frontend uses JSON
		ExposeHeaders:    []string{"Content-Length"}, //Response headers browser is allowed to see
		AllowCredentials: true,                       //Allows cookies/auth tokens to be sent w requests
		MaxAge:           12 * time.Hour,             //How long browser should cache CORS preflight response (avoids sending preflight every time)

		// Preflight -> when browser asks server if a req is okay
	}))

	router.Use(api.MaxBodySize(5 << 20))

	router.Use(api.ErrorHandler())

	/*Define first route for root
	GET - registers a get route
		first arg is path
		second arg is anonymous func, which will be called whenever someone access route

	func(c *gin.Context) - handler function
		c- context provided by Gin (gives access to request, response and helpers)

	c.JSON - sends json message back to client
		200 - HTTPS status code
		gin.H - shortcut for creating JSON object
		"Send this JSON along with a code 200OK"
	*/

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the AI Code Reviewer!",
		})
	})

	//API versioning
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})

		/*
			Post route -> /analyze
			func(gin context) -> handler function
			c-> context
			empty variable req is created with type CodeSubmissionRequest

		*/
		v1.POST("/analyze", func(c *gin.Context) {
			var req api.CodeSubmissionRequest

			// err Tries to parse incoming json body into req struct
			// if err is not nil (aka -> it receives error message)
			// If there's error we send a 400 Bad response
			// gin.H is helper for building JSON map
			//return -> used to exit early
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			//If there's no error, we return a 200 OK code and the following response
			c.JSON(http.StatusOK, gin.H{
				"message":  "Code analysis complete",
				"filename": req.Filename,
				"language": req.Language,
				"issues":   []string{}, //Placeholder issues
			})
		})

	}

	//Define string with port value
	port := "8080"
	//Logs message so server is running
	log.Printf("Listening on http://localhost:%s", port)

	//Create custom HTTP server with timeouts
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 5 << 20,
	}

	//Start
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}

}
