package parser

import(
  "encoding/json"
  "io/ioutil"
  "log"
  "fmt"
  "inventory"
)

func Parse(fileName string) (map[string]int, []inventory.Inventory, error){

  var data struct{
    Order map[string]int `json:"order"`
    Inventory []inventory.Inventory `json:"data"`
  }
  content, err := ioutil.ReadFile(fileName)
  if err != nil {
		log.Fatal(err)
	}
  if err := json.Unmarshal(content, &data); err != nil {
		return nil, nil, fmt.Errorf("cannot parse json: %v, %v", string(content), err)
	}
  return data.Order, data.Inventory, nil
}
