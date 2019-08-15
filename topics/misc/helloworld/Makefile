CGO_ENABLED = 1
TARGETS = main maindl

.PHONY: all
all: $(TARGETS)

main: main.go
	go build main.go

maindl: maindl.go
	CGO_ENABLED=$(CGO_ENABLED) go build maindl.go

.PHONY: clean
clean:
	rm -f main maindl
