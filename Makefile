
go-build:
	go build -o app-core

docker-build:
	docker build -t app-core .


docker-run:
	docker run app-core