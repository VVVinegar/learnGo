package parser

import (
	"io/ioutil"
	"learngo/crawler/engine"
	"learngo/model"
	"log"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile(
		"profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := parseProfile(contents,
		"http://album,zhenai.com/u/108906739",
		"张一")
	log.Printf("test: %v", result.Items[0])

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+
			"element; but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url:  "http://album,zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Age:        28,
			Height:     0,
			Weight:     0,
			Income:     "",
			Gender:     "",
			Name:       "张一",
			Xinzuo:     "",
			Occupation: "",
			Marriage:   "",
			House:      "",
			Hokou:      "",
			Education:  "",
			Car:        "",
		},
	}

	if actual != expected {
		t.Errorf("expected %v; but was %v",
			expected, actual)
	}
}
