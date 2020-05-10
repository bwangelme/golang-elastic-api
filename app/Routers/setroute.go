package Routers

import (
	//"fmt"
	"net/http"

	"github.com/bwangelme/golang-elastic-api/app/Models"
	"github.com/bwangelme/golang-elastic-api/app/Utils"
	"golang.org/x/net/context"
	//"github.com/olivere/elastic/v7"
)

func SetHandler(w http.ResponseWriter, r *http.Request) {
	dat := Utils.BodyToJson(r)
	bodyData := dat["source"]
	id := dat["id"].(string)
	parent_id := dat["parent_id"].(string)
	operation := dat["operation"].(string)

	client := Models.GetElasticCon(Utils.ElasticUrl())
	indexService := client.Index().Index(Index)
	updateSevice := client.Update().Index(Index)
	deleteService := client.Delete().Index(Index)
	if operation == "add" {
		if parent_id != "" {
			indexService = indexService.Parent(parent_id)
		}
		_, _ = indexService.Id(id).BodyJson(bodyData).Do(context.Background())
	} else if operation == "update" {
		if parent_id != "" {
			updateSevice = updateSevice.Parent(parent_id)
		}
		_, _ = updateSevice.Id(id).Doc(bodyData).DetectNoop(true).Do(context.TODO())
	} else if operation == "delete" {
		if parent_id != "" {
			deleteService = deleteService.Id(id)
		}
		_, _ = deleteService.Do(context.TODO())
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok"}`))
}
