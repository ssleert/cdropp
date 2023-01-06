build:
	go build -ldflags "-s -w" cmd/cdropp/cdropp.go

install:
	install ./cdropp /usr/local/bin

uninstall:
	rm -rf /usr/local/bin/cdropp