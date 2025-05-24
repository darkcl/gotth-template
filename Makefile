.PHONY: build

templ:
	templ generate --watch --proxy="http://localhost:8090" --open-browser=false -v

# Run air to detect any go file changes to re-build and re-run the server.
server:
	air \
	--build.cmd "go build -o tmp/bin/main ./cmd/main.go" \
	--build.bin "tmp/bin/main" \
	--build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

tailwind-clean:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --clean

# Run tailwindcss to generate the styles.css bundle in watch mode.
tailwind-watch:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

# Start development server
dev:
	make tailwind-clean
	make -j3 tailwind-watch templ server

# Production Build
build:
	rm -Rf build/
	mkdir -p build
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --minify
	go build -o build/server ./cmd/main.go
