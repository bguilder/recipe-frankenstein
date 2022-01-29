package handler

import (
	"io/ioutil"
	"os"
)

const (
	feelingHungryLocation = "./ingredients_fixtures/feeling_hungry.json"
)

func OpenFeelingHungry() ([]byte, error) {
	jsonFile, err := os.Open(feelingHungryLocation)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	result, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	return result, nil
}
