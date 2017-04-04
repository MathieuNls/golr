test:
	go test -v --coverprofile coverage ./engine
	go tool cover -func=coverage
	go test -v --coverprofile coverage ./solrjson
	go tool cover -func=coverage

build:
	go get github.com/mathieunls/golr/engine
	go get github.com/mathieunls/golr/solrjson
	go get github.com/jarcoal/httpmock
	go get github.com/stretchr/testify/assert