APP_NAME=go_blogs

build: 
	make swagger && go build -o bin/$(APP_NAME) cmd/main.go

swagger:
	swag init -g cmd/main.go

run:
	go run cmd/main.go