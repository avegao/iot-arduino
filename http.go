package main

import (
    "github.com/hooklift/httpclient"
    "encoding/json"
)

func GetJson(url string, target interface{}) error {
    client := httpclient.Default()
    response, err := client.Get(url)

    if err != nil {
        return err
    }

    defer response.Body.Close()

    return json.NewDecoder(response.Body).Decode(target)
}
