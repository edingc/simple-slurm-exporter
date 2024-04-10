.PHONY: all build

all: build

build:
  @echo "Building simple-slurm-json-exporter..."
  @go build -o simple-slurm-json-exporter main.go
