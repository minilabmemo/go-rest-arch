
BIN_FILE_NAME=app-core

go-build:
	@ printf "Building aplication... "
	@ go build -o $(BIN_FILE_NAME) .
	@ echo "done"

docker-build:
	docker build -t app-core .


docker-run:
	docker run app-core