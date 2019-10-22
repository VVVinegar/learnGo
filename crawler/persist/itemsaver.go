package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"learngo/crawler/engine"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error)  {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: Got item #%d: %v", itemCount, item)

			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
			Do(context.Background())
	if err != nil  {
		return err
	}

	return nil
}
