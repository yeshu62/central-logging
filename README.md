# Distributed logging servers

A distributed logging system made using gin framework where multiple services can log messages to a central loggin server

### Requirements
- Go version 1.22.4

### Steps
- ```git clone https://github.com/yeshu62/central-logging.git```
- ```go mod init```
- ```go mod tidy```
- ```go run .```

## Viewing the Logs:
Navigate to http://localhost:8081/logs in your web browser to view the logs
or
```curl http://localhost:8081/logs```

  
