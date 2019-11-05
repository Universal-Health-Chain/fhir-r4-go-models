package models

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestObservationR4(t *testing.T) {
	path := "../examplesForTesting/observation-hemoglobin.json"
	observation, err := readObservationJson(path)
	assert.Equal(t, err, nil, "OK observation loaded")
	assert.NotEmpty(t, observation.Interpretation, "OK observation loaded")

}

func readObservationJson(path string) (observation Observation, err error) {
	observation = Observation{}
	file, err1 := ioutil.ReadFile(path)
	if err1 != nil {
		fmt.Println(err1)
		return observation, err1
	}
	err2 := json.Unmarshal([]byte(file), &observation)
	if err2 != nil {
		fmt.Println(err2)
		return observation, err2
	}
	return observation, nil
}
