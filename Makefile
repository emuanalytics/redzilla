
run:
	go run main.go

clean:
	rm -f redzilla

build:
	CGO_ENABLED=0 go build -a -ldflags '-s' -o redzilla

docker/build:
	docker build . -t opny/redzilla:latest

docker/push: docker/build
	docker push opny/redzilla:latest
