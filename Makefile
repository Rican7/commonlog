build:
	go build -v

install:
	go install

install-deps:
	glide install

install-deps-dev: install-deps
	go get github.com/golang/lint/golint

update-deps:
	glide update

update-deps-dev: update-deps
	go get -u github.com/golang/lint/golint

test:
	go test -v $$(glide novendor)

test-with-coverage:
	go test -cover $$(glide novendor)

test-with-coverage-formatted:
	go test -cover $$(glide novendor) | column -t | sort -r

lint: install-deps-dev
	errors=$$(gofmt -l .); if [ "$${errors}" != "" ]; then echo "$${errors}"; exit 1; fi
	errors=$$(glide novendor | xargs -n 1 golint -min_confidence=0.3); if [ "$${errors}" != "" ]; then echo "$${errors}"; exit 1; fi

vet:
	go vet $$(glide novendor)
