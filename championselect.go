package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiKey             = "RGAPI-faecaef7-51c3-401f-902a-1072e311d77d"
	championV3Url      = "https://euw1.api.riotgames.com/lol/platform/v3/champions"
	lolStaticDataV3Url = "https://euw1.api.riotgames.com/lol/static-data/v3/champions"
)

type lolStaticJSON struct {
	Title string `json:"title"`
	Name  string `json:"name"`
	Key   string `json:"key"`
	ID    int    `json:"id"`
}

type championJSON struct {
	Champions []championJSONInner `json:"champions"`
}

type championJSONInner struct {
	RankedPlayEnabled bool `json:"rankedPlayEnabled"`
	BotEnabled        bool `json:"botEnabled"`
	BotMmEnabled      bool `json:"botMmEnabled"`
	Active            bool `json:"active"`
	FreeToPlay        bool `json:"freeToPlay"`
	ID                int  `json:"id"`
}

type championInfo struct {
	ID   int
	Name string
}

func main() {
	ids := getChampionIDs()
	champions := getChampionInfo(ids...)
	fmt.Println(champions)
}

func getChampionIDs() []int {
	c := championJSON{}
	getJSON(championV3Url, &c)
	var ids []int
	for _, v := range c.Champions {
		ids = append(ids, v.ID)
	}
	return ids
}

func getChampionInfo(ids ...int) []championInfo {
	var c []championInfo
	for _, id := range ids {
		url := fmt.Sprintf("%s%s%d", lolStaticDataV3Url, "/", id)
		b := lolStaticJSON{}
		getJSON(url, &b)
		c = append(c, championInfo{ID: id, Name: b.Name})
		fmt.Println(id, b.Name)
	}
	return c
}

func getJSON(baseURL string, target interface{}) error {
	url := baseURL + "?api_key=" + apiKey
	res, err := http.Get(url)
	check(err)
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
