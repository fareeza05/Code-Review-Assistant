package models

//ReviewRequest = JSON payload client sends to request code review
//AKA - defines what client sends backend during POST

type AnalyzeRequest struct {
	//json:"field" -> tells Go how to map JSON keys to go fields
	//binding: "required" -> Used by Gin to validate that field must be present
	Code     string `json:"code" binding:"required"`
	Language string `json:"language" binding:"required"`
}

// ReviewResponse = JSON response returned by API
// Defines what backend sends back to client for a GET request
// POST /analyze response (202 Accepted)
type AnalyzeResponse struct {
	ID      string `json:"id"` //UUID
	Message string `json:"message"`
}

// GET /results/{id} response
type AnalysisResult struct {
	ID          string   `json:"id"`
	Summary     string   `json:"summary"`
	Issues      []string `json:"issues"`
	Suggestions []string `json:"suggestions"`
}

// GET /status/{id} response
type AnalysisStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// Generic Error response
type ErrorResponse struct {
	Error string `json:"error"`
}
