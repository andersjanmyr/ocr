sources = $(shell find . -path ./vendor -prune -o -name '*.go' -print)

dist:
	mkdir dist

dist/ocr.exe: $(sources) | dist
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o dist/ocr.exe

dist/ocr-osx: $(sources) | dist
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -installsuffix cgo -ldflags '-s' -o dist/ocr-osx

dist/ocr-linux: $(sources) | dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags '-s' -o dist/ocr-linux

.PHONY: build run tag release clean
build: dist/ocr.exe dist/ocr-osx dist/ocr-linux

run:
	go run main.go

tag:
	./scripts/tag.sh $(VERSION)

release: tag build
	./scripts/release.sh ocr $(VERSION) dist/*

clean :
	-rm -r dist
