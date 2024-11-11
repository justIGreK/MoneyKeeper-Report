PWD=$(shell pwd)

.PHONY: generate
generate: # generates *.go files from proto.
	@docker run -it --rm \
		--volume $(PWD):/workspace \
		--workdir /workspace \
		bufbuild/buf:1.36.0 generate --include-imports
