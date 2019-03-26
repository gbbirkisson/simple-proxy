IMAGE:=gbbirkisson/simple-proxy:latest

build:
	docker build -t ${IMAGE} .

push:
	docker push ${IMAGE}