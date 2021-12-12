build:
	go build -o ytmapi main.go

run: build
	go run ./main.go
