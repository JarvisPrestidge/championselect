package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	apiKey             = "RGAPI-faecaef7-51c3-401f-902a-1072e311d77d"
	championV3Url      = "https://euw1.api.riotgames.com/lol/platform/v3/champions?freeToPlay=false"
	lolStaticDataV3Url = "https://euw1.api.riotgames.com/lol/static-data/v3/champions/1?locale=en_GB"
)

type lolStaticV3Response struct {
	Title string `json:"title"`
	Name  string `json:"name"`
	Key   string `json:"key"`
	ID    int    `json:"id"`
}

type championV3Response struct {
	Champion []struct {
		RankedPlayEnabled bool `json:"rankedPlayEnabled"`
		BotEnabled        bool `json:"botEnabled"`
		BotMmEnabled      bool `json:"botMmEnabled"`
		Active            bool `json:"active"`
		FreeToPlay        bool `json:"freeToPlay"`
		ID                int  `json:"id"`
	} `json:"champions"`
}

func main() {

	body := genericAPICall(lolStaticDataV3Url, &lolStaticV3Response{})
	a := lolStaticV3Response{}
	err = json.Unmarshal(body, &a)
	check(err)

}

func genericAPICall(baseURL string) *[]byte {
	url := baseURL + "&api_key=" + apiKey
	res, err := http.Get(url)
	check(err)

	body, err := ioutil.ReadAll(res.Body)
	check(err)
	defer res.Body.Close()

	return &body
}

// check is a helper error function
func check(e error) {
	if e != nil {
		panic(e)
	}
}
