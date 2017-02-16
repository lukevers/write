VERSION=0.1.0
BRANCH=git rev-parse --abbrev-ref HEAD
COMMIT=git rev-parse HEAD
COMMIT_COUNT=git rev-list --count $$(${BRANCH})
LDFLAGS="-X main.Version=${VERSION} -X main.Commit=`${COMMIT}`"

default: sass macos

build: go sass macos

go:
	go get -u -v

sass:
	sass --style compressed resources/scss/write.scss:resources/css/write.css

macos:
	macpack build -version=${VERSION} -build-number=$$(${COMMIT_COUNT}) -ldflags=${LDFLAGS}
