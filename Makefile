.PHONY: run build mock test db proto

run: proto
	go run .

build: proto
	go build .

mock:
	mockgen -destination=./mocks/dependency_injector/mock_manager.go				-source=./dependency_injector/interface/interface.go  	-package=mock_dependency_injector	. DependencyInjector
	mockgen -destination=./mocks/logger/mock_logger.go								-source=./logger/interface.go  							-package=mock_logger				. Logger

test: mock
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out

db:
	docker-compose -f docker-compose.yml up --build --remove-orphans db

proto:
	sh ./compile_proto.sh