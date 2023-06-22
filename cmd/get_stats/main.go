package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/mihkulemin/go-elasticmon/pkg/metrics"
)

func main() {
	client, err := es.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Nodes.Stats(client.Nodes.Stats.WithLevel("shards"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %v\n", resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(data))

	var nodeStats metrics.NodeStatsResponse
	err = json.Unmarshal(data, &nodeStats)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nodeStats)
}
