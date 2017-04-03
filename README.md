[![Build Status](https://travis-ci.org/MathieuNls/golr.png)](https://travis-ci.org/MathieuNls/golr)
[![GoDoc](https://godoc.org/github.com/MathieuNls/golr?status.png)](https://godoc.org/github.com/MathieuNls/golr)
[![codecov](https://codecov.io/gh/MathieuNls/golr/branch/master/graph/badge.svg)](https://codecov.io/gh/MathieuNls/golr)

# golr

A simple lib for Solr written in Go.

# Install

```bash
go get github.com/mathieunls/golr
```

# Examples

### Update

```go

type techProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func newTechProduct(id string, name string) *techProduct {
	product := new(techProduct)
	product.ID = id
	product.Name = name
	return product
}

func main(){

    //URL, timeout, commit delay
    solr := NewSolr("http://127.0.0.1:8983/solr/aa", 10, 10)
    products := make([]*techProduct, 3, 3)

    //Existing products get updated
	products[0] = newTechProduct("GB18030TEST", "Test")
	products[1] = newTechProduct("SP2514N", "Test2")
    //New product gets created
	products[2] = newTechProduct("NewOne", "NewOne")

	val, err := solr.Update(products)
}
```

### Select

```go
func main(){
    query := NewSolrQuery()
 
	solr := NewSolr("http://127.0.0.1:8983/solr/aa", 10, 10)
	
    /** Produces 
    {  
        query:"*:*",
        offset:1,
        limit:1,
        filter:[  
            "a",
            "b"
        ],
        sort:"a desc",
        facet:{  
            avg_price:"avg(price)"
        }
    }*/
    query.Facet("avg_price", "avg(price)").
		Sort("a", "desc").
		Field("a").
		Field("b").
		Offset(1).Limit(1).Filter("a", "b").Query("*:*")

    val, err := solr.Query(query)

    fmt.Println(val) // Maps that contain the query result
}
```

### Delete

```go
func main(){
    solr := NewSolr("http://127.0.0.1:8983/solr/aa", 10, 10)
    solr.Delete([]string{"SP2514N"})
}
```