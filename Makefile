default: build

build: html2words

clean:
	rm -f html2words

html2words: html2words.go
	go build ./...
