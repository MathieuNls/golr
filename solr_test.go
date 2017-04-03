package golr

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestNewSolr(t *testing.T) {

	solr := NewSolr("http://127.0.0.1:8983/solr/aa", 10, 10)

	assert := assert.New(t)
	assert.NotNil(solr)
}

func TestQuerySolr(t *testing.T) {

	query := NewSolrQuery()
	solr := NewSolr("http://127.0.0.1:8983/solr/aa", 10, 10)
	query.Query("*:*")

	mock := "{ \"responseHeader\":{ \"status\":0, \"QTime\":0, \"params\":{ \"q\":\"*:*\", \"indent\":\"on\", \"wt\":\"json\", \"_\":\"1491244092319\"}}, \"response\":{\"numFound\":30,\"start\":0,\"docs\":[ { \"id\":\"GB18030TEST\", \"name\":\"Test with some GB18030 encoded characters\", \"features\":[\"No accents here\", \"\u8FD9\u662F\u4E00\u4E2A\u529F\u80FD\", \"This is a feature (translated)\", \"\u8FD9\u4EFD\u6587\u4EF6\u662F\u5F88\u6709\u5149\u6CFD\", \"This document is very shiny (translated)\"], \"price\":0.0, \"price_c\":\"0.0,USD\", \"inStock\":true, \"_version_\":1563670104514756608}, { \"id\":\"SP2514N\", \"name\":\"Samsung SpinPoint P120 SP2514N - hard drive - 250 GB - ATA-133\", \"manu\":\"Samsung Electronics Co. Ltd.\", \"manu_id_s\":\"samsung\", \"cat\":[\"electronics\", \"hard drive\"], \"features\":[\"7200RPM, 8MB cache, IDE Ultra ATA-133\", \"NoiseGuard, SilentSeek technology, Fluid Dynamic Bearing (FDB) motor\"], \"price\":92.0, \"price_c\":\"92.0,USD\", \"popularity\":6, \"inStock\":true, \"manufacturedate_dt\":\"2006-02-13T15:26:37Z\", \"manufacturedate_pdt\":\"2006-02-13T15:26:37Z\", \"store\":\"35.0752,-97.032\", \"_version_\":1563670104566136832}, { \"id\":\"6H500F0\", \"name\":\"Maxtor DiamondMax 11 - hard drive - 500 GB - SATA-300\", \"manu\":\"Maxtor Corp.\", \"manu_id_s\":\"maxtor\", \"cat\":[\"electronics\", \"hard drive\"], \"features\":[\"SATA 3.0Gb/s, NCQ\", \"8.5ms seek\", \"16MB cache\"], \"price\":350.0, \"price_c\":\"350.0,USD\", \"popularity\":6, \"inStock\":true, \"store\":\"45.17614,-93.87341\", \"manufacturedate_dt\":\"2006-02-13T15:26:37Z\", \"manufacturedate_pdt\":\"2006-02-13T15:26:37Z\", \"_version_\":1563670104585011200}, { \"id\":\"F8V7067-APL-KIT\", \"name\":\"Belkin Mobile Power Cord for iPod w/ Dock\", \"manu\":\"Belkin\", \"manu_id_s\":\"belkin\", \"cat\":[\"electronics\", \"connector\"], \"features\":[\"car power adapter, white\"], \"weight\":4.0, \"price\":19.95, \"price_c\":\"19.95,USD\", \"popularity\":1, \"inStock\":false, \"store\":\"45.18014,-93.87741\", \"manufacturedate_dt\":\"2005-08-01T16:30:25Z\", \"manufacturedate_pdt\":\"2005-08-01T16:30:25Z\", \"_version_\":1563670104591302656}, { \"id\":\"IW-02\", \"name\":\"iPod & iPod Mini USB 2.0 Cable\", \"manu\":\"Belkin\", \"manu_id_s\":\"belkin\", \"cat\":[\"electronics\", \"connector\"], \"features\":[\"car power adapter for iPod, white\"], \"weight\":2.0, \"price\":11.5, \"price_c\":\"11.50,USD\", \"popularity\":1, \"inStock\":false, \"store\":\"37.7752,-122.4232\", \"manufacturedate_dt\":\"2006-02-14T23:55:59Z\", \"manufacturedate_pdt\":\"2006-02-14T23:55:59Z\", \"_version_\":1563670104592351232}, { \"id\":\"MA147LL/A\", \"name\":\"Apple 60 GB iPod with Video Playback Black\", \"manu\":\"Apple Computer Inc.\", \"manu_id_s\":\"apple\", \"cat\":[\"electronics\", \"music\"], \"features\":[\"iTunes, Podcasts, Audiobooks\", \"Stores up to 15,000 songs, 25,000 photos, or 150 hours of video\", \"2.5-inch, 320x240 color TFT LCD display with LED backlight\", \"Up to 20 hours of battery life\", \"Plays AAC, MP3, WAV, AIFF, Audible, Apple Lossless, H.264 video\", \"Notes, Calendar, Phone book, Hold button, Date display, Photo wallet, Built-in games, JPEG photo playback, Upgradeable firmware, USB 2.0 compatibility, Playback speed control, Rechargeable capability, Battery level indication\"], \"includes\":\"earbud headphones, USB cable\", \"weight\":5.5, \"price\":399.0, \"price_c\":\"399.00,USD\", \"popularity\":10, \"inStock\":true, \"store\":\"37.7752,-100.0232\", \"manufacturedate_dt\":\"2005-10-12T08:00:00Z\", \"manufacturedate_pdt\":\"2005-10-12T08:00:00Z\", \"_version_\":1563670104597594112}, { \"id\":\"adata\", \"compName_s\":\"A-Data Technology\", \"address_s\":\"46221 Landing Parkway Fremont, CA 94538\", \"_version_\":1563670104604934144}, { \"id\":\"apple\", \"compName_s\":\"Apple\", \"address_s\":\"1 Infinite Way, Cupertino CA\", \"_version_\":1563670104605982720}, { \"id\":\"asus\", \"compName_s\":\"ASUS Computer\", \"address_s\":\"800 Corporate Way Fremont, CA 94539\", \"_version_\":1563670104605982721}, { \"id\":\"ati\", \"compName_s\":\"ATI Technologies\", \"address_s\":\"33 Commerce Valley Drive East Thornhill, ON L3T 7N6 Canada\", \"_version_\":1563670104607031296}] }}"
	var f interface{}
	err := json.Unmarshal([]byte(mock), &f)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// mock to list out the articles
	httpmock.RegisterResponder("GET", "http://127.0.0.1:8983/solr/aa/query",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, f)

			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	val, err := solr.Query(query)

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(f, val)
}

func TestDelete(t *testing.T) {
	solr := NewSolr("http://127.0.0.1:8983/solr/aa", 10, 10)

	mock := "{\"responseHeader\":{\"status\":0,\"QTime\":1}}"
	var f interface{}
	err := json.Unmarshal([]byte(mock), &f)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// mock to list out the articles
	httpmock.RegisterResponder("POST", "http://127.0.0.1:8983/solr/aa/update",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, f)

			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	val, err := solr.Delete([]string{"SP2514N"})

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(f, val)
}

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

func TestUpdate(t *testing.T) {

	solr := NewSolr("http://127.0.0.1:8983/solr/aa", 10, 10)

	mock := "{\"responseHeader\":{\"status\":0,\"QTime\":1}}"
	var f interface{}
	err := json.Unmarshal([]byte(mock), &f)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// mock to list out the articles
	httpmock.RegisterResponder("POST", "http://127.0.0.1:8983/solr/aa/update?commitWithin"+strconv.Itoa(solr.commit),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, f)

			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	products := make([]*techProduct, 3, 3)

	products[0] = newTechProduct("GB18030TEST", "Test")
	products[1] = newTechProduct("SP2514N", "Test2")
	products[2] = newTechProduct("NewOne", "NewOne")

	val, err := solr.Update(products)

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(f, val)

}
