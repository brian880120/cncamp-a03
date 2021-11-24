dep:
	go mod vendor

run: dep
	go run main.go
