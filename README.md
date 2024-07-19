
# WebSocket-Based Real-Time Messaging Application 

This project implements a real-time messaging application using WebSocket technology. The backend is developed in Go, leveraging the Gin framework to handle WebSocket connections and serve static files. The frontend, composed of HTML, CSS, and JavaScript, interacts with the WebSocket server to enable real-time communication between clients.








## Table of Content
- Prerequisites
- Folder Structure
- Build and Run the Application
- Accessing the Application
- Configuration Details
- Screeshots


## Prerequisites

Before running the application , ensure you have the following installed:

- Go(1.22 or higher)
- Docker

Though Docker can be exempted , You may clone the application and run locally with go compiler



## Folder Structure
```bash
.
├── Client
│   ├── index.html
│   ├── script.js
│   └── style.css
├── Dockerfile
└── Server
    ├── go.mod
    ├── go.sum
    └── main.go
```
## Build and Run the Application

1. Build the Docker Image:

From the root directory of the project run :

```bash
docker build -t websocket-chat-application .
```

2. Run the Docker Container

```bash
docker run -d -p 8080:8080 websocket-chat-application
```

3. Verify the Container is Running

```bash
docker ps
```

4. Check the logs of the Application

```bash
docker logs <container-id>
```

Without Docker (Locally)

1. Clone the repository

```bash 
git clone https://github.com/Nikhils-179/go-chat-application.git
```

2. Navigate to the Server directory

```bash
cd server
```

3. Run the Go Application

```bash
go build -o main . 
```

4. Run the Go Application
```bash
./main
```

5. Access the Application 

Open your web browser and navigate to http://localhost to access the WebSocket chat application.



## Configuration Details

1. Backend Configuration: The Go server implements the WebSocket protocol using Gorilla WebSocket. It handles incoming messages, reverses them, maintains a history of the last 5 messages, and serves static files from the /static directory.

2. Frontend Configuration: The static files (index.html, style.css, script.js) are served by the Go server and provide the user interface and WebSocket interactions.

3. Security Configuration : Added security to restrict WebSocket connections to only allow requests from `localhost:8080`. This is enforced by checking the `Origin` header in the WebSocket upgrade request. Connections from other origins will be rejected to prevent unauthorized access.
