package data

import (
	"encoding/json"
	"groupie/structdata"
	"net/http"
)

func FetchData(apiURL string) ([]structdata.GroupData, error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data []structdata.GroupData
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func FetchDataL(apiURL string) (structdata.LocationData, error) {
	var data structdata.LocationData

	response, err := http.Get(apiURL)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func FetchDataD(apiURL string) (structdata.DatesData, error) {
	var data structdata.DatesData

	response, err := http.Get(apiURL)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func FetchDataR(apiURL string) (structdata.RelationData, error) {
	var relations structdata.RelationData

	response, err := http.Get(apiURL)
	if err != nil {
		return relations, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&relations)
	if err != nil {
		return relations, err
	}
	return relations, nil

}
