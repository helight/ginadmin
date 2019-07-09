TARGET=main/admin

all:
	go build -mod=vendor -o $(TARGET) main/main.go
#	go build -mod=vendor -ldflags -extldflags=-static -o $(TARGET) main/main.go
#	go build -mod=vendor -o demo main/demo.go

clean:
	go clean
	rm $(TARGET)