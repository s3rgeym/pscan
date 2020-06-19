PREFIX ?= /usr/local
# без этого флага не компилируется
export GO111MODULE = on

build:
	go build -v -o ./bin/pscan ./cmd/pscan/main.go

clean:
	rm -rf bin/*

install:
	install -m 755 ./bin/* $(PREFIX)/bin

test:
	go test -v ./...

uninstall:
	rm $(PREFIX)/bin/pscan

# если бы в каталоге был файл install, то `make install` запустило бы его без
# этой настройки
.PHONY: build clean install test uninstall
