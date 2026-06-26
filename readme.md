# Monach C2 - A simple C2 Client/Server in Python & Go

## Educational Use Only

This project is intended solely for educational, research, and authorized security testing purposes. It must only be used in environments where you have explicit permission to do so. Any unauthorized or malicious use is strictly prohibited. The author assumes no responsibility for misuse or damages resulting from the use of this software.

## About 
This project demonstrates the fundamentals of a Command and Control (C2) architecture. The server coordinates communications and task management, while the client establishes a connection to receive instructions and return results.

## Getting started 

You will need :
- Python3
- Go 1.21+

### Server
First you need to deploy the server.
After creating your own venv environment

```bash
    pip install -r requirements.txt

    # Then 

    fastapi dev
```

### Client 
For now, the client will be compiled with thoose. 
We'll add an Make in the next versions 

```bash
go mod init

# For Windows
GOOS=windows GOARCH=amd64 go build -o client.exe .

# Linux 
GOOS=linux GOARCH=amd64 go build -o client .

# MacOS
GOOS=darwin GOARCH=amd64 go build -o client_mac .

# If you are using Powershell, you'll need to do it this way 
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o client .
```

### Next steps 

- Soon
