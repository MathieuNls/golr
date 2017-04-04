test:
	go test -coverprofile=engine.cover.out -covermode=atomic ./engine
	go test -coverprofile=solr.cover.out -covermode=atomic  ./solrjson
	cat *.cover.out >> coverage.txt && rm *cover.out
build:
	go get -t -v ./...