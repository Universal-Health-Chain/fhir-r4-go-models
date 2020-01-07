package models

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestEncounterR4(t *testing.T) {
	path := "../examplesForTesting/encounter-example-f203.json"
	encounter, err := readEncounterJson(path)
	assert.Equal(t, err, nil, "OK encounter loaded")
	assert.NotEmpty(t, encounter.Identifier, "OK encounter loaded")
}

func readEncounterJson(path string) (encounter Encounter, err error) {
	encounter = Encounter{}
	file, err1 := ioutil.ReadFile(path)
	if err1 != nil {
		fmt.Println(err1)
		return encounter, err1
	}
	err2 := json.Unmarshal([]byte(file), &encounter)
	if err2 != nil {
		fmt.Println(err2)
		return encounter, err2
	}
	return encounter, nil
}
