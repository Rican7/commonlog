build:
	go build -v

install:
	go install

install-deps:
	go get -d -t ./...

install-deps-dev: install-deps
	go get github.com/golang/lint/golint

update-deps:
	go get -d -t -u ./...

update-deps-dev: update-deps
	go get -u github.com/golang/lint/golint

test:
	go test -v ./...

test-with-coverage:
	go test -cover ./...

test-with-coverage-formatted:
	go test -cover ./... | column -t | sort -r

lint: install-deps-dev
	errors=$$(gofmt -l .); if [ "$${errors}" != "" ]; then echo "$${errors}"; exit 1; fi
	errors=$$(golint -min_confidence=0.3 ./...); if [ "$${errors}" != "" ]; then echo "$${errors}"; exit 1; fi

vet:
	go vet ./...
