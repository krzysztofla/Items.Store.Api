run:
	go run main.go
test:
	go test
swagger: 
	swagger generate spec -o ./swagger.yaml --scan-models