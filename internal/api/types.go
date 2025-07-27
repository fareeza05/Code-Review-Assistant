package api

/*
What client's JSON request structure will look like
*/
type CodeSubmissionRequest struct {
	Language string `json:"language" binding:"required"`
	Filename string `json:"filename" binding:"required"`
	Content  string `json:"content" binding:"required"`
}
