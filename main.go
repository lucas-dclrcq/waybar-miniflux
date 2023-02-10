package main

import (
	"encoding/json"
	"flag"
	"fmt"
	miniflux "miniflux.app/client"
	"strconv"
)

func main() {
	url := flag.String("url", "", "Miniflux url")
	token := flag.String("token", "", "Miniflux API token")
	flag.Parse()

	if *url == "" {
		fmt.Println("url should be specified")
		return
	}

	if *token == "" {
		fmt.Println("token should be specified")
		return
	}

	client := miniflux.New(*url, *token)

	counters, err := client.FetchCounters()

	if err != nil {
		fmt.Println(err)
		return
	}

	totalUnreads := getTotalUnreads(counters.UnreadCounters)

	fmt.Println(formatWaybarInput(totalUnreads))
}

func formatWaybarInput(totalUnreads int) string {
	waybarInput := newWaybarInput(strconv.Itoa(totalUnreads))
	jsonOutput, _ := json.Marshal(waybarInput)
	return string(jsonOutput)
}

func getTotalUnreads(unreads map[int64]int) int {
	var totalUnread = 0
	for _, val := range unreads {
		totalUnread += val
	}
	return totalUnread
}

type WaybarInput struct {
	Text       string `json:"text"`
	Alt        string `json:"alt"`
	Tooltip    string `json:"tooltip"`
	Class      string `json:"class"`
	Percentage string `json:"percentage"`
}

func newWaybarInput(text string) *WaybarInput {
	w := WaybarInput{Text: text}
	return &w
}
