URL_ICANN=https://data.iana.org/TLD/tlds-alpha-by-domain.txt

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
	rm parcels regexp.go go.sum

.PHONY: build
build: regexp.go
	go build

