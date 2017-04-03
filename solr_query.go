package golr

import (
	"strconv"
	"strings"
)

type SolrQuery struct {
	query   string
	filters []string
	offset  int
	limit   int
	fields  []string
	sort    string
	facets  []string
}

func NewSolrQuery() *SolrQuery {

	query := new(SolrQuery)
	query.filters = []string{}
	query.fields = []string{}
	query.facets = []string{}

	return query
}

func (solrQuery *SolrQuery) Query(query string) *SolrQuery {

	solrQuery.query = query
	return solrQuery
}

func (solrQuery *SolrQuery) Filter(field string, value string) *SolrQuery {
	solrQuery.filters = append(solrQuery.filters, "\""+field+":"+value+"\"")
	return solrQuery
}

func (solrQuery *SolrQuery) Offset(offset int) *SolrQuery {
	solrQuery.offset = offset
	return solrQuery
}

func (solrQuery *SolrQuery) Limit(limit int) *SolrQuery {
	solrQuery.limit = limit
	return solrQuery
}

func (solrQuery *SolrQuery) Field(field string) *SolrQuery {
	solrQuery.fields = append(solrQuery.fields, "\""+field+"\"")
	return solrQuery
}

func (solrQuery *SolrQuery) Sort(field string, order string) *SolrQuery {
	solrQuery.sort = "\"" + field + " " + order + "\""
	return solrQuery
}

func (solrQuery *SolrQuery) Facet(field string, value string) *SolrQuery {
	solrQuery.facets = append(solrQuery.facets, field+":\""+value+"\"")
	return solrQuery
}

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
