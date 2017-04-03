package golr

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//Solr represents a Solr cnx
type Solr struct {
	url     string
	timeout int
	commit  int
}

//NewSolr creates a Solr cnx
func NewSolr(url string, timeout int, commit int) *Solr {
	solr := new(Solr)
	solr.url = url
	solr.timeout = timeout
	solr.commit = commit
	return solr
}

//Update send a JSON update to the Solr API using marshmalling of data
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

//Delete sends a delete call to the Solr API to delete the ids
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

//Query sends the SolrQuery to the Solr API
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
