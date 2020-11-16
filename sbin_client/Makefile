build: folder
	go build -o bin/client
folder:
	mkdir -p bin

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o src/client


	