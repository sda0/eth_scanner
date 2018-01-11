RELEASE?=1.0.0
GOOS?=linux

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

.PHONY: check
check:
	go vet $$(go list ./... | grep -v /vendor/)

.PHONY: test
test:
	go test -v $$(go list ./... | grep -v /vendor/)

.PHONY: build
build: clean
	CGO_ENABLED=0 GOOS=${GOOS} go build -o bin/${GOOS}/ethScanner \
		-ldflags "-X main.version=${RELEASE} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_TIME}" \
		main.go

.PHONY: clean
clean:
	@rm -rf bin/${GOOS}/*


HAS_GLIDE := $(shell command -v glide;)

.PHONY: vendor
vendor: prepare_glide
	glide install

.PHONY: prepare_glide
prepare_glide:
ifndef HAS_GLIDE
	curl https://glide.sh/get | sh
endif

.PHONY: initdb
initdb:
	docker-compose up -d db
	sleep 3
	docker exec -t `docker ps -f name=_db_1 -q | head -n1` bash -c "psql -U postgres < /data/postgres/backup/dump.sql"

