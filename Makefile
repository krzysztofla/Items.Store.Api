run:
	go run main.go
test:
	go test
swagger: 
	swag init
continer:
	docker build -t items-store-api .
	docker run --name items-store-go-api -dp 8080:8080 items-store-api