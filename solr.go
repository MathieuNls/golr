package golr

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Solr struct {
	url     string
	timeout int
	commit  int
}

func NewSolr(url string, timeout int, commit int) *Solr {
	solr := new(Solr)
	solr.url = url
	solr.timeout = timeout
	solr.commit = commit
	return solr
}

func (solr *Solr) Update(data interface{}) (interface{}, error) {

	value, err := json.Marshal(data)

	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", solr.url+"/update?commitWithin"+strconv.Itoa(solr.commit), bytes.NewBuffer(value))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var f interface{}
	err = json.Unmarshal(body, &f)

	if err != nil {
		return false, err
	}

	return f, nil
}

func (solr *Solr) Delete(ids []string) (interface{}, error) {

	//append " for json marshalling
	for index := 0; index < len(ids); index++ {
		ids[index] = "\"" + ids[index] + "\""
	}

	jsonIds := "{ \"delete\":[" + strings.Join(ids, ", ") + "] }"

	req, err := http.NewRequest("POST", solr.url+"/update", bytes.NewBuffer([]byte(jsonIds)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var f interface{}
	err = json.Unmarshal(body, &f)

	if err != nil {
		return false, err
	}

	return f, nil
}

func (solr *Solr) Query(query *SolrQuery) (interface{}, error) {

	req, err := http.NewRequest("GET", solr.url+"/query", bytes.NewBuffer(query.Prepare()))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var f interface{}
	err = json.Unmarshal(body, &f)

	if err != nil {
		return nil, err
	}

	return f, nil
}
