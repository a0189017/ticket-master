
docker-build:
	docker build --build-arg SOURCE_FILES=main -t hongyu/ticket-master:latest .

mockery:
	mockery --all --dir=pkg/repository --output=./mocks --outpkg=mocks --with-expecter

test:
	go test -v ./pkg/usecase/...