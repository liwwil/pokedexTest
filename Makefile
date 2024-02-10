run:
	go run main.go

gqlgen:
	go get github.com/99designs/gqlgen@v0.17.43

	go run github.com/99designs/gqlgen generate