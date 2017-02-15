VERSION=0.1
BRANCH=git rev-parse --abbrev-ref HEAD
COMMIT=git rev-parse HEAD
COMMIT_COUNT=git rev-list --count $$(${BRANCH})
LDFLAGS="-X main.Version=${VERSION} -X main.Commit=`${COMMIT}`"

mac:
	macpack build -version=${VERSION} -build-number=$$(${COMMIT_COUNT}) -ldflags=${LDFLAGS}
