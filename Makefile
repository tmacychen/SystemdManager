all: *.go
	go build .

clean: SystemdManager
	rm SystemdManager
	rm log
