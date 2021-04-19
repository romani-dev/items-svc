

build:
	CGO_ENABLED=0 go build -o bin/app main.go

docker: build
	docker build . -t items-svc:latest