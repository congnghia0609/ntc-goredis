package main

import (
	"encoding/json"
	"fmt"
)

func sjson2map(s string) map[string]interface{} {
	// string json to map
	if len(s) > 0 {
		var mapData map[string]interface{}
		in := []byte(s)
		json.Unmarshal(in, &mapData)
		return mapData
	}
	return nil
}

func map2sjson(data map[string]interface{}) string {
	if data != nil {
		out, _ := json.Marshal(data)
		return string(out)
	}
	return ""
}

func exampleUJson() {
	blockId := uint64(10)
	blockHeader := "0x04a4fcf765d61e99fc2a9c785f4505f32de74c38ec2d0d120b5c278d5659e087"
	seedHash := "0x0000000000000000000000000000000000000000000000000000000000000000"
	difficulty := "0x00007fe007fe007fe007fe007fe007fe007fe007fe007fe007fe007fe007fe00"
	mapData := make(map[string]interface{})
	mapData["type_xmpp"] = "XMPP_Topic_Disc_TaskMn"
	mapData["type_msg"] = "DTM_Task_PoW"
	mapData["blockId"] = blockId
	mapData["blockHeader"] = blockHeader
	mapData["seedHash"] = seedHash
	mapData["difficulty"] = difficulty
	sjson := map2sjson(mapData)
	fmt.Println("sjson=", sjson)
	data := sjson2map(sjson)
	fmt.Println("type_xmpp:", data["type_xmpp"], ", type_msg:", data["type_msg"], ", blockId:", data["blockId"], ", blockHeader:", data["blockHeader"],
		", difficulty:", data["difficulty"])
}
