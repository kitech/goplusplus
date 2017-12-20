VERSION         :=      $(shell cat ./VERSION)
# VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`


# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD} -X main.Entry=main"
GOVVV=`govvv -flags -version ${VERSION}|sed 's/=/=GOVVV-/g'`
# 	go build -v -ldflags "-w -s ${GOVVV}" .

all:

ins:
	go install -v

bud:
	go build -v ${LDFLAGS}

test:
	cd ./lib && go test -v

fmt:
	go fmt

# See more at `gorelease` github repo.
release:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)
	goreleaser --rm-dist

dist:
	goxc -bc="linux,!arm windows,386 darwin"

.PHONY: install test fmt release
