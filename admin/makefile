TARGET=admin

all:
	go build -mod=vendor -o $(TARGET) main.go
#	go build -mod=vendor -ldflags -extldflags=-static -o $(TARGET) main.go
#	go build -mod=vendor -o demo main/demo.go

clean:
	go clean
	rm $(TARGET)
