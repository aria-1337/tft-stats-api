package main

import (
    "net/http"
    "encoding/json"
    "time"
    "fmt"
)

var client = &http.Client{Timeout: 10 * time.Second}

func request(url string, target interface{}, key string) error {
    r, err := client.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

// /riot/account/v1/accounts/by-riot-id/{gameName}/{tagLine}
// {gameName} = going 8th
// {tagLin} = 7330 (no #)
// Returns { puuid (string), gameName (string), tagLine (string) }
type RAccount struct {
    Puuid string `json:"puuid"`
    GameName string `json:"gameName"`
    TagLine string `json:"tagLine"`
}

func r_account_riot_id(key string, region string, gameName string, tagLine string) *RAccount {
    response := &RAccount{}
    err := request("https://" + region + ".api.riotgames.com/riot/account/v1/accounts/by-riot-id/" + gameName + "/" + tagLine + "?api_key=" + key, response, key)
    if err != nil {
        // TODO: This should eventaully return an error
        fmt.Println("Error when calling: r_account_riot_id: ", err)
    }
    return response
}
