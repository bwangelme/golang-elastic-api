package Routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bwangelme/golang-elastic-api/app/Entity"
	"github.com/bwangelme/golang-elastic-api/app/Models"
	"github.com/bwangelme/golang-elastic-api/app/Utils"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
)

var Index = "search"

func GetHandler(w http.ResponseWriter, r *http.Request) {
	client := Models.GetElasticCon(Utils.ElasticUrl())
	dat := Utils.BodyToJson(r)
	query_type := dat["query_type"].(string)
	child_type := dat["child_type"].(string)
	start := int(dat["start"].(float64))
	array_of_json := dat["query_json"].([]interface{})
	size := int(dat["size"].(float64))
	sorting := dat["sort"].(map[string]interface{})

	var fieldName string
	var sortType bool
	for i := range sorting {
		if i == "field" {
			fieldName = sorting[i].(string)
		} else if i == "asc" {
			sortType = true
		}
	}
	bq := elastic.NewBoolQuery()
	if query_type == "parent" {
		datRecord := array_of_json[0]
		res := datRecord.(map[string]interface{})
		key := res["key"].(string)
		value := res["value"].(string)

		matchChildQuery := elastic.NewHasChildQuery(child_type, elastic.NewMatchQuery(key, value)).
			InnerHit(elastic.NewInnerHit().Name("messages"))
		bq = bq.Must(elastic.NewMatchAllQuery())
		bq = bq.Filter(matchChildQuery)
	} else {
		for i := 0; i < len(array_of_json); i++ {
			datRecord := array_of_json[i]
			res := datRecord.(map[string]interface{})
			qType := res["query_type"].(string)
			matchQueryType := res["match"].(string)
			key := res["key"].(string)
			value := res["value"].(interface{})
			var matchType *elastic.MatchQuery
			var termQuery *elastic.TermQuery
			var rangeQuery *elastic.RangeQuery
			match := 0
			switch matchQueryType {
			case "text":
				value := res["value"].(string)
				matchType = elastic.NewMatchQuery(key, value)
				break
			case "keyword":
				match = 1
				termQuery = elastic.NewTermQuery(key, value)
				break
			case "range":
				match = 2
				rangeQuery = elastic.NewRangeQuery(key)
				valueRange := value.(map[string]interface{})

				for i := range valueRange {
					switch i {
					case "gte":
						rangeQuery = rangeQuery.Gte(valueRange[i])
						break
					case "gt":
						rangeQuery = rangeQuery.Gt(valueRange[i])
						break
					case "lte":
						rangeQuery = rangeQuery.Lte(valueRange[i])
						break
					case "lt":
						rangeQuery = rangeQuery.Lt(valueRange[i])
						break
					}

				}
				break
			}
			switch qType {
			case "must":
				if match == 0 {
					bq = bq.Must(matchType)
				} else {
					bq = bq.Must(termQuery)
				}
				break
			case "filter":
				if match == 0 {
					bq = bq.Filter(matchType)
				} else {
					bq = bq.Filter(termQuery)
				}
				break
			case "must_not":
				if match == 0 {
					bq = bq.MustNot(matchType)
				} else {
					bq = bq.MustNot(termQuery)
				}
				break
			case "should":
				if match == 0 {

					bq = bq.Should(matchType)
				} else {
					bq = bq.Should(termQuery)

				}
				break
			}
		}
	}

	var searchResult *elastic.SearchResult

	// 输出查询的结构
	//ss := elastic.NewSearchSource().Query(bq).From(start).Size(size).Sort(fieldName, sortType)
	//d, _ := ss.Source()
	//data, _ := json.Marshal(d)
	//log.Println(string(data))

	eQuery := client.Search(Index).
		Query(bq).
		From(start).Size(size)

	if fieldName != "" {
		eQuery = eQuery.Sort(fieldName, sortType)
	}

	searchResult, err := eQuery.Pretty(true).Do(context.Background())
	if err != nil {
		panic(err)
	}
	hits := searchResult.Hits.Hits

	datArray := make([]map[string]interface{}, len(hits))
	var dat1 map[string]interface{}

	for i := 0; i < len(hits); i++ {
		hit := searchResult.Hits.Hits[i]
		if err := json.Unmarshal(hit.Source, &dat1); err != nil {
			panic(err)
		}
		datArray[i] = dat1
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	response := Entity.JsonResponse{"data_source": datArray, "status": true, "length": len(hits)}
	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(b))

}
