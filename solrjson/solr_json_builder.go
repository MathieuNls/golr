package golr

import (
	"fmt"
	"strconv"
	"strings"
)

//SolrJSONBuilder represents a query that is being build for Solr
type SolrJSONBuilder struct {
	query   string
	filters []string
	offset  int
	limit   int
	fields  []string
	sort    string
	facets  []string
}

//NewSolrQuery creates a SolrQuery and initialize each field properly
func NewSolrJSONBuilder() *SolrJSONBuilder {

	query := new(SolrJSONBuilder)
	query.filters = []string{}
	query.fields = []string{}
	query.facets = []string{}

	return query
}

//Query sets the query field
func (solrQuery *SolrJSONBuilder) Query(query string) *SolrJSONBuilder {

	solrQuery.query = query
	return solrQuery
}

//Filter add a filter to the filter array
func (solrQuery *SolrJSONBuilder) Filter(field string, value string) *SolrJSONBuilder {
	solrQuery.filters = append(solrQuery.filters, "\""+field+":"+value+"\"")
	return solrQuery
}

//Offset sets the offset field
func (solrQuery *SolrJSONBuilder) Offset(offset int) *SolrJSONBuilder {
	solrQuery.offset = offset
	return solrQuery
}

//Limit sets the limit field
func (solrQuery *SolrJSONBuilder) Limit(limit int) *SolrJSONBuilder {
	solrQuery.limit = limit
	return solrQuery
}

//Field add a field to the array of selected fields
func (solrQuery *SolrJSONBuilder) Field(field string) *SolrJSONBuilder {
	solrQuery.fields = append(solrQuery.fields, "\""+field+"\"")
	return solrQuery
}

//Sort set the sort field
func (solrQuery *SolrJSONBuilder) Sort(field string, order string) *SolrJSONBuilder {
	solrQuery.sort = "\"" + field + " " + order + "\""
	return solrQuery
}

//Facet add a facet criterion to the facet array
func (solrQuery *SolrJSONBuilder) Facet(field string, value string) *SolrJSONBuilder {
	solrQuery.facets = append(solrQuery.facets, field+":\""+value+"\"")
	return solrQuery
}

//Prepare creates a []byte containing all the field and ready to be send via http
func (solrQuery *SolrJSONBuilder) Prepare() []byte {

	jsonValue := "{"

	if len(solrQuery.query) == 0 {
		solrQuery.query = "*:*"
	}

	jsonValue += "query : \"" + solrQuery.query + "\""

	if solrQuery.offset != 0 {
		jsonValue += ", offset : " + strconv.Itoa(solrQuery.offset)
	}

	if solrQuery.limit != 0 {
		jsonValue += ", limit : " + strconv.Itoa(solrQuery.limit)
	}

	if len(solrQuery.filters) > 0 {
		jsonValue += ", filter : [" + strings.Join(solrQuery.filters, ",") + "]"
	}

	if len(solrQuery.fields) > 0 {
		jsonValue += ", fields : [" + strings.Join(solrQuery.fields, ",") + "]"
	}

	if len(solrQuery.sort) > 0 {
		jsonValue += ", sort : " + solrQuery.sort
	}

	if len(solrQuery.facets) > 0 {
		jsonValue += ", facet : { " + strings.Join(solrQuery.facets, ",") + "}"
	}

	jsonValue += "}"

	fmt.Println(jsonValue)

	return []byte(jsonValue)
}
