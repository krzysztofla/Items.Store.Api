run:
	go run main.go
test:
	go test
swagger: 
	swag init
docker:
	docker build -t getting-started .