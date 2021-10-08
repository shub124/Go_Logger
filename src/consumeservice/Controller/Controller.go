package consumercontroller

import (
	meta "consumeservice/Meta"
	"encoding/json"
	"log"

	elastic "gopkg.in/olivere/elastic.v7"

	//solr "solretl"
	"context"
	"fmt"
)

func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

var Controllermap = map[string]func(msg []byte){
	"Cricket_Log": cricketlogcontroller,
	"User_Log":    userlogcontroller,
}

func cricketlogcontroller(msg []byte) {

	log.Println("Hit")

	cricmeta := new(meta.Cricketmeta)
	err := json.Unmarshal(msg, &cricmeta)
	if err != nil {
		log.Println("Cricket_KPI_ERROR", err.Error())
	}
	UpdatetElasticsearch(cricmeta)
	//solr.UpdateCricketKPI(cricmeta)

}

func userlogcontroller(msg []byte) {

}

func UpdatetElasticsearch(cric *meta.Cricketmeta) {

	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	data, err := json.Marshal(cric)
	js := string(data)
	ind, err := esclient.Index().
		Index("cricket").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")
	fmt.Println(ind)

}
