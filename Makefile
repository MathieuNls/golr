test:
	go test -v --coverprofile coverage ./engine
	go tool cover -func=coverage
	go test -v --coverprofile coverage ./solrjson
	go tool cover -func=coverage