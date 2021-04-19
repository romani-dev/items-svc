

build:
	CGO_ENABLED=0 go build main.go

docker: build
	docker build . -t items-svc:latest