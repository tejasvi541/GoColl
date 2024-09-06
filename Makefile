#DEV

build-dev:
	docker build -t gocoll -f containers/images/Dockerfile . && docker build -t turn -f containers/turn/Dockerfile.turn .
clean-dev:
	docker-compose -f containers/composes/dc.dev.yml down
run-dev:
	docker-compose -f containers/composes/dc.dev.yml up -d
