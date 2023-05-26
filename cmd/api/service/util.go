package service

import (
	"encoding/json"
	"os"
)

type Data struct {
	AccountId int
	Id        int
	Advisor   string
	Body      string
}

type Payload struct {
	Data []Data
}

func raw() ([]Data, error) {

	r, err := os.ReadFile("data.json")

	if err != nil {
		panic(err)
		return nil, err
	}
	var payload Payload

	err = json.Unmarshal(r, &payload.Data)

	if err != nil {
		panic(err)
		return nil, err
	}
	return payload.Data, nil

}

func GetAll() ([]Data, error) {
	data, err := raw()
	if err != nil {
		panic(err)

	}
	return data, nil
}

func GetById(idx int) (any, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}
	if idx > len(data) {
		r := make([]string, 0)
		return r, nil
	}
	return data[idx], nil

}
