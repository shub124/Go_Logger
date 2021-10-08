package solretl

import (
	meta "consumeservice/Meta"
	"time"

	"log"

	solr "github.com/sendgrid/go-solr"
)

func Updatetosolr(cric *meta.Cricketmeta) {
	solrzk := solr.NewSolrZK("localhost:2182", "solr", "Cricket_Info")
	err := solrzk.Listen()
	if err != nil {
		panic(err.Error())
	}
	https, err := solrzk.UseHTTPS()
	if err != nil {
		panic(err.Error())
	}
	solrhttp, err := solr.NewSolrHTTP(https, "Cricket_Info", solr.MinRF(2), solr.QueryRouter(solr.NewRoundRobinRouter()))
	if err != nil {
		panic(err.Error())
	}
	solrClient := solr.NewSolrHttpRetrier(solrhttp, 5, 100*time.Millisecond)

	locator := solrzk.GetSolrLocator()

	data := map[string]interface{}{
		"name": cric.Cricketer_name, "runs": cric.Runs, "average": cric.Average, "matches": cric.Matches,
	}
	all, err := locator.GetLeadersAndReplicas(data["name"].(string))
	if err != nil {
		panic(err.Error())

	}
	err = solrClient.Update(all, true, data, solr.Commit(false))
	if err != nil {
		panic(err.Error())
	}

}

func UpdateCricketKPI(cric *meta.Cricketmeta) {

	log.Println("Reached here")
	go Updatetosolr(cric)
}
