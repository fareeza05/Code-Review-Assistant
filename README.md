# AI Code Reviewer 🧠💻
This is a simple Go-based web API for submitting code and receiving automated analysis results.
It is built using the [Gin](https://github.com/gin-gonic/gin) framework and structured for scalability and clean architecture

---
## Features
- `POST /analyze`: Submit code for review
- JSON request validation
- CORS support for frontend integration
- Modular structure with Go best practices

---

## Technical Stack
- Go
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- Standard Library
- Community Middleware

---

## Running the Server

1. Clone the repo:
```
git clone https://github.com/fareeza05/Code-Review-Assistant.git
cd Code-Review-Assistant
```

2. Install dependencies
`Go mod tidy`

3. Run the server
`go run ./cmd/server`

4. Test API with postman / curl
```
curl -X POST http://localhost:8080/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "language": "go",
    "filename": "main.go",
    "content": "package main\nfunc main() {}"
  }'

```

---

## Project Structure
```
ai-code-reviewer/
├── cmd/ # Application entry point
│ └── server/
│ └── main.go
├── internal/
│ ├── api/ # Request/response types and handlers
│ │ └── types.go
│ ├── models/ # Core domain models (if needed)
│ └── services/ # Business logic (future expansion)
├── go.mod # Go module definition
├── go.sum # Dependency checksums
├── .gitignore
└── README.md
```


---

## API Endpoints

### `POST /analyze`
Submit code for analysis

**Request Body (JSON)**
```
json 
{
    "language": "go",
  "filename": "main.go",
  "content": "package main\nfunc main() {}"
}
```

**Response**

```
{
  "message": "Code analysis complete",
  "filename": "main.go",
  "language": "go",
  "issues": []
}
```

---

## Future Roadmap / To-Do
-  Integrate AI code analysis
- Add authentication
- Add logging with request IDs
- Support file uploads
- Add frontend/CLI client
- Dockerize the app
