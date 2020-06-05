
TAG := $(shell git describe --tags | grep "^[v0-9.]*" -o)

run:
	go run main.go

clean:
	rm -f redzilla

build:
	CGO_ENABLED=0 go build -a -ldflags '-s' -o redzilla

docker/build: build
	docker build . -t emuanalytics/redzilla:$(TAG)
	docker tag emuanalytics/redzilla:$(TAG) emuanalytics/redzilla:latest

docker/push:
	docker push emuanalytics/redzilla:$(TAG)
	docker push emuanalytics/redzilla:latest

docker/release: docker/build docker/push
