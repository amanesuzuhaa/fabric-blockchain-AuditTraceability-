package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)


//formatJSON format json data
func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, " ", ""); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return prettyJSON.String()
}

//constructDataByID structure data by ID
func constructDataByID(id string) Asset{
	asset := Asset{
		ID: 	id,
		USER:   randUser(),
		TIME: 	time.RFC3339,
		INFO: 	randInfo(),
	}
	return asset
}

// constructData construct data
func constructData(values ...string)  (Asset, error){

	var testAsset Asset
	getType := reflect.TypeOf(testAsset)
	//fmt.Println(len(values))
	//fmt.Println(getType.NumField())
	if len(values) != getType.NumField(){
		return testAsset, fmt.Errorf("error: paras failed")
	}
	asset := Asset{
		ID: 	values[0],
		TIME: 	values[1],
		USER:   values[2],
		INFO: 	values[3],
	}
	return asset, nil
}
