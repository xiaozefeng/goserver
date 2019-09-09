all: gotool
	@go build -v .
clean:
	rm -f goserver
gotool:
	gofmt -w .
	go tool vet . |& grep -v vendor;true

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"

.PHONY: clean gotool help