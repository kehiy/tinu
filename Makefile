check:
	golangci-lint run \
		--build-tags "${BUILD_TAG}" \
		--timeout=20m0s \
		--enable=gofmt \
		--enable=unconvert \
		--enable=unparam \
		--enable=asciicheck \
		--enable=misspell \
		--enable=revive \
		--enable=decorder \
		--enable=reassign \
		--enable=usestdlibvars \
		--enable=nilerr \
		--enable=gosec \
		--enable=exportloopref \
		--enable=whitespace \
		--enable=goimports \
		--enable=gocyclo \
		