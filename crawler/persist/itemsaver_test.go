package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"learngo/crawler/engine"
	"learngo/model"
	"testing"
)

func TestSaver(t *testing.T) {
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

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// save
	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// get
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
