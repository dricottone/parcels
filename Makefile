INSTALL_DIR?=/usr/local/bin
URL_ICANN=https://data.iana.org/TLD/tlds-alpha-by-domain.txt
GO_FILES=$(shell find -name '*.go')

build/top-level-domains.txt:
	curl -o build/top-level-domains.txt $(URL_ICANN)

regexp.go: build/top-level-domains.txt
	tail +2 build/top-level-domains.txt \
		| perl -pe 'chomp if eof' \
		| tr '[:upper:]' '[:lower:]' \
		| tr '\n' '|' > build/regexp.go-part2
	cat <(perl -pe 'chomp if eof' build/regexp.go-part1) \
		<(perl -pe 'chomp if eof' build/regexp.go-part2) \
		build/regexp.go-part3 \
		> regexp.go

.PHONY: clean
clean:
	rm -f parcels regexp.go go.sum

parcels: $(GO_FILES) regexp.go
	go build

.PHONY: build
build: parcels

.PHONY: install
install: parcels
	install -m755 parcels $(INSTALL_DIR)/parcels

