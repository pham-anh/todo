package todo

import (
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	var items []Item
	if err = json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	return items, nil
}
