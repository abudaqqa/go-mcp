to build the server for LibreChat docker container you must use those params
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o myserver
```
