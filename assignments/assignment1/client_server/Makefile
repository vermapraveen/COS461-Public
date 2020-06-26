default: go

all: go

go: client-go server-go

client-go: client-go.go
	go build client-go.go

server-go: server-go.go
	go build server-go.go

clean:
	rm -f server-go client-go
