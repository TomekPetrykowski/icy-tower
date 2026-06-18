alias r := run

build:
    @echo "Building the project..."
    go build -o bin/icy-tower cmd/icy-tower/main.go

run: build
    @echo "Running the project..."
    ./bin/icy-tower
