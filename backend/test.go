package main

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
)

func main() {
	var document string = `{
	  "cobrandId": 10010352,
	  "channelId": -1,
	  "locale": "en_US",
	  "tncVersion": 2,
	  "categories":["rabbit","bunny","frog"],
	  "people": [
	 	{
			"name": "jack",
			"age": {
				"birth":10,
				"year":2000,
				"animals": [
					{
						"barks":"yes",
						"tail":"yes"
					},
					{
						"barks":"no",
						"tail":"yes"
					}
				]
			}
		},
		{
			"name": "jill",
			"age": {
				"birth":11,
				"year":2001
			}
		}
	  ]
}`

	type Animal struct {
		Barks string `jpath:"barks"`
	}

	type People struct {
		Age     int      `jpath:"age.birth"` // jpath is relative to the array
		Animals []Animal `jpath:"age.animals"`
	}

	type Items struct {
		Categories []string `jpath:"categories"`
		Peoples    []People `jpath:"people"` // Specify the location of the array
	}

	docScript := []byte(document)
	var docMap map[string]interface{}
	json.Unmarshal(docScript, &docMap)
	fmt.Println(docMap)
	var items Items
	mapstructure.DecodePath(docMap, &items)

	fmt.Printf("%#v", items)
	// Output:
	// mapstructure.Items{Categories:[]string{"rabbit", "bunny", "frog"}, Peoples:[]mapstructure.People{mapstructure.People{Age:10, Animals:[]mapstructure.Animal{mapstructure.Animal{Barks:"yes"}, mapstructure.Animal{Barks:"no"}}}, mapstructure.People{Age:11, Animals:[]mapstructure.Animal(nil)}}}
}

