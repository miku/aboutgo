data.db:
	sqlite3 $@ < create.sql

webcapture: main.go
	go build -o $@ $<

.PHONY: clean
clean:
	rm -f webcapture

