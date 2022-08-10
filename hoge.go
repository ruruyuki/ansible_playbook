package main

import "fmt"

func main() {

    idList := [...]map[string]string{
        {"ID": "1", "AID": "a1", "BID": "b1"},
        {"ID": "2", "AID": "a2", "BID": "b2"},
        {"ID": "3", "AID": "a3", "BID": "b3"},
    }

    mapList := []map[string]string{}
    for _, i := range idList {
        data := map[string]string{
            "id":     i["ID"],
            "aaa_id": i["AID"],
            "bbb_id": i["BID"],
            "memo":   "移動済み",
        }
        mapList = append(mapList, data)

    }
    fmt.Print(mapList)
}