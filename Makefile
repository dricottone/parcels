INSTALL_DIR?=/usr/local/bin
URL_ICANN=https://data.iana.org/TLD/tlds-alpha-by-domain.txt
GO_FILES=$(shell find -name '*.go')

scripts/top-level-domains.txt:
	curl -o scripts/top-level-domains.txt $(URL_ICANN)

common/regexp.go: scripts/top-level-domains.txt
	tail +2 scripts/top-level-domains.txt \
		| perl -pe 'chomp if eof' \
		| tr '[:upper:]' '[:lower:]' \
		| tr '\n' '|' > scripts/regexp.go-part2
	cat <(perl -pe 'chomp if eof' scripts/regexp.go-part1) \
		<(perl -pe 'chomp if eof' scripts/regexp.go-part2) \
		scripts/regexp.go-part3 \
		> common/regexp.go

.PHONY: clean
clean:
	rm -f parcels common/regexp.go go.sum

parcels: common/regexp.go $(GO_FILES)
	go build

.PHONY: build
build: parcels

.PHONY: install
install: parcels
	install -m755 parcels $(INSTALL_DIR)/parcels

