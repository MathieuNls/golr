package golr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	query := NewSolrJSONBuilder()

	assert := assert.New(t)

	assert.NotNil(query.filters)
	assert.NotNil(query.fields)
	assert.NotNil(query.facets)
}

func TestQuery(t *testing.T) {

	query := NewSolrJSONBuilder()
	query.Query("*:*")

	assert := assert.New(t)
	assert.Equal("*:*", query.query)

}

func TestFilter(t *testing.T) {

	query := NewSolrJSONBuilder()
	query.
		Filter("a", "b").
		Filter("c", "d")

	assert := assert.New(t)
	assert.Equal("\"a:b\"", query.filters[0])
	assert.Equal("\"c:d\"", query.filters[1])

}

func TestLimit(t *testing.T) {

	query := NewSolrJSONBuilder()
	query.Limit(1)

	assert := assert.New(t)
	assert.Equal(1, query.limit)

}

func TestOffset(t *testing.T) {

	query := NewSolrJSONBuilder()
	query.Offset(1)

	assert := assert.New(t)
	assert.Equal(1, query.offset)

}

func TestField(t *testing.T) {

	query := NewSolrJSONBuilder()
	query.Field("a").Field("b")

	assert := assert.New(t)
	assert.Equal("\"a\"", query.fields[0])
	assert.Equal("\"b\"", query.fields[1])
}

func TestSort(t *testing.T) {

	query := NewSolrJSONBuilder()
	query.Sort("a", "desc")

	assert := assert.New(t)
	assert.Equal("\"a desc\"", query.sort)

}

func TestFacet(t *testing.T) {

	query := NewSolrJSONBuilder()
	query.Facet("avg_price", "avg(price)")

	assert := assert.New(t)
	assert.Equal("avg_price:\"avg(price)\"", query.facets[0])

}

func TestPrepare(t *testing.T) {

	query := NewSolrJSONBuilder()

	query.Facet("avg_price", "avg(price)").
		Sort("a", "desc").
		Field("a").
		Field("b").
		Offset(1).Limit(1).Filter("a", "b").
		Filter("c", "d").Query("*:*")

	assert := assert.New(t)
	assert.Equal("{query : \"*:*\", offset : 1, limit : 1, filter : [\"a:b\",\"c:d\"], filter : [\"a\",\"b\"], sort : \"a desc\", facet : { avg_price:\"avg(price)\"}}", string(query.Prepare()))

}
