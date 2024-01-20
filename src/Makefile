run:
	go run .\src\cmd\rest\main.go

test:
	go test -v -cover ./...

build:
	docker build --tag keyword-generator:1.0 .

docker-run:
	docker run -d -p 8010:8010 --env-file .env keyword-generator:1.0