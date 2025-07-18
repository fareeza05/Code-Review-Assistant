package main //Marks file as standalone executable program (not reusable). All programs must start with a main package to run

/*Pulls in Go's standard libraries
"fmt" //For formatted output like print
"log" //For logging messages to terminal
"net/http" //To build HTTP servers
*/
import (
	"fmt"
	"log"
	"net/http"
)

// Every Go program starts execution at main() function
func main() {
	/*
		Sets up a route. When someone visits /(root of site) this will run
		HandleFunc -> takes in path, and handler function which is func
			Log -> Logs incoming method request
			fmt -> Writes response back to user, w is ResponseWriter(how you send data back to client)
				Fprintln -> prints text followed by new line
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)
		fmt.Fprintln(w, "Welcome to the AI Code Reviewer!")
	})

	//Define string with port value
	port := "8080"
	//Logs message so server is running
	log.Printf("Listening on http://localhost:%s", port)

	/*
		Declares new var - err
		Runs ListenAndServe -> tries to connect to the port specified
		The value returned by running that is then assigned to err
		If err is nil error hasnt occured, server is properly listening
		ListenandServe is both the entry point to the server and also what continues to listen and handle any changes
	*/
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
