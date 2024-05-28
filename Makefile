.PHONY: build run clean

build:
    go build -o todo-app

run: build
    ./todo-app

clean:
    rm -f todo-app
