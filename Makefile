PROGRAM_NAME=gocr

init:
	dep ensure

build:
	go build -o $(PROGRAM_NAME)