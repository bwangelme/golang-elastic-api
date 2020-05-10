package Routers

import (
	"fmt"
	"net/http"

	"github.com/bwangelme/golang-elastic-api/app/Models"
	"github.com/bwangelme/golang-elastic-api/app/Utils"

	"context"
	"encoding/json"
)

func MappingHandler(w http.ResponseWriter, r *http.Request) {
	client := Models.GetElasticCon(Utils.ElasticUrl())
	data := Utils.BodyToJson(r)
	mappingData := data["mapping_json"].(map[string]interface{})
	//mappingIndex := elastic.NewAliasAddAction()
	putMappingResponse, err := client.PutMapping().Index(Utils.DefaultIndex()).BodyJson(mappingData).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(putMappingResponse)
	newMapping, err := client.GetMapping().Index(Utils.DefaultIndex()).Do(context.Background())
	//var b map[string]interface{}
	b, marshalErr := json.Marshal(newMapping)
	if marshalErr != nil {
		panic(marshalErr)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte("Gorilla Map!\n"))
	//fmt.Println(client.ClusterState())
	fmt.Fprint(w, string(b))

}
