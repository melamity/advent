all: build/2024-01 build/2024-02

build/2024-01:
	go build -C build advent/2024-01

build/2024-02:
	go build -C build advent/2024-02