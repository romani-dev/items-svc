

build:
	CGO_ENABLED=0 go build -o bin/app main.go

docker: build
ifndef TAG
	$(error TAG is not set)
endif
	docker build . -t clazz/items-svc:$(TAG)

push: docker
ifndef TAG
	$(error TAG is not set)
endif
	docker push clazz/items-svc:$(TAG)
