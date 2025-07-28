package models

//ReviewRequest = JSON payload client sends to request code review
//AKA - defines what client sends backend during POST

type ReviewRequest struct {
	//json:"field" -> tells Go how to map JSON keys to go fields
	//binding: "required" -> Used by Gin to validate that field must be present
	Code     string `json:"code" binding:"required"`
	Language string `json:"language" binding:"required"`
}

// ReviewResponse = JSON response returned by API
// Defines what backend sends back to client for a GET request
type ReviewResponse struct {
	Summary     string   `json:"summary"`
	Issues      []string `json:"issues"`
	Suggestions []string `json:"suggestions"`
}
