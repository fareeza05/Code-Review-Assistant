package models

//ReviewRequest = JSON payload client sends to request code review
//AKA - defines what client sends backend during POST

type AnalyzeRequest struct {
	//json:"field" -> tells Go how to map JSON keys to go fields
	//binding: "required" -> Used by Gin to validate that field must be present
	Code     string `json:"code" binding:"required, min=10"`
	Language string `json:"language" binding:"required, oneof=go python javascript"`
}

// ReviewResponse = JSON response returned by API
// Defines what backend sends back to client for a GET request
// POST /analyze response (202 Accepted)
type AnalyzeResponse struct {
	ID      string `json:"id" binding:"required, uuid"` //UUID - required
	Message string `json:"message" binding:"required"`  //required
}

// GET /results/{id} response
type AnalysisResult struct {
	ID          string   `json:"id" binding:"required, uuid"`
	Summary     string   `json:"summary" binding:"required"`
	Issues      []string `json:"issues" binding:"required"`
	Suggestions []string `json:"suggestions" binding:"required"`
}

// GET /status/{id} response
type AnalysisStatus struct {
	ID     string `json:"id" binding:"required, uuid"`
	Status string `json:"status" binding:"required"`
}

// Generic Error response
type ErrorResponse struct {
	Error string `json:"error"`
}
