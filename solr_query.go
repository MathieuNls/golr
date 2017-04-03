package golr

import (
	"strconv"
	"strings"
)

//SolrQuery represents a query that is being build for Solr
type SolrQuery struct {
	query   string
	filters []string
	offset  int
	limit   int
	fields  []string
	sort    string
	facets  []string
}

//NewSolrQuery creates a SolrQuery and initialize each field properly
func NewSolrQuery() *SolrQuery {

	query := new(SolrQuery)
	query.filters = []string{}
	query.fields = []string{}
	query.facets = []string{}

	return query
}

//Query sets the query field
func (solrQuery *SolrQuery) Query(query string) *SolrQuery {

	solrQuery.query = query
	return solrQuery
}

//Filter add a filter to the filter array
func (solrQuery *SolrQuery) Filter(field string, value string) *SolrQuery {
	solrQuery.filters = append(solrQuery.filters, "\""+field+":"+value+"\"")
	return solrQuery
}

//Offset sets the offset field
func (solrQuery *SolrQuery) Offset(offset int) *SolrQuery {
	solrQuery.offset = offset
	return solrQuery
}

//Limit sets the limit field
func (solrQuery *SolrQuery) Limit(limit int) *SolrQuery {
	solrQuery.limit = limit
	return solrQuery
}

//Field add a field to the array of selected fields
func (solrQuery *SolrQuery) Field(field string) *SolrQuery {
	solrQuery.fields = append(solrQuery.fields, "\""+field+"\"")
	return solrQuery
}

//Sort set the sort field
func (solrQuery *SolrQuery) Sort(field string, order string) *SolrQuery {
	solrQuery.sort = "\"" + field + " " + order + "\""
	return solrQuery
}

//Facet add a facet criterion to the facet array
func (solrQuery *SolrQuery) Facet(field string, value string) *SolrQuery {
	solrQuery.facets = append(solrQuery.facets, field+":\""+value+"\"")
	return solrQuery
}

//Prepare creates a []byte containing all the field and ready to be send via http
func (solrQuery *SolrQuery) Prepare() []byte {

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
		jsonValue += ", filter : [" + strings.Join(solrQuery.fields, ",") + "]"
	}

	if len(solrQuery.sort) > 0 {
		jsonValue += ", sort : " + solrQuery.sort
	}

	if len(solrQuery.facets) > 0 {
		jsonValue += ", facet : { " + strings.Join(solrQuery.facets, ",") + "}"
	}

	jsonValue += "}"

	return []byte(jsonValue)
}
