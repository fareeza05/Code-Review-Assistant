package main //Marks file as standalone executable program (not reusable). All programs must start with a main package to run

import (
	"log"

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

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	//Define string with port value
	port := "8080"
	//Logs message so server is running
	log.Printf("Listening on http://localhost:%s", port)

	/*
		Declares new var - err
		Runs route
		The value returned by running that is then assigned to err
		If err is nil error hasnt occured, server is properly listening
	*/
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
