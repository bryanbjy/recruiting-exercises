package inventory

import (
  "fmt"
  "sync"
)

type InventoryAllocator struct {

}

type Inventory struct {
  Name string `json:"name"`
  Inventory map[string]int `json:"inventory"`
}

func (i *InventoryAllocator) SetOrder(orderName string, orderNumber int, inventory []Inventory, wg *sync.WaitGroup) {
  index := 0
  execOrder := make([]Inventory,0)

  for orderNumber > 0 && index < len(inventory) {
    target := inventory[index].Inventory
    if val, ok := target[orderName]; ok {
      if orderNumber - val < 0 {
        val = orderNumber
        orderNumber = 0
        newOrder := &Inventory{Name: inventory[index].Name, Inventory: map[string]int{orderName:val}}
        execOrder = append(execOrder, *newOrder)
      } else {
        orderNumber -= val
        newOrder := &Inventory{Name: inventory[index].Name, Inventory: map[string]int{orderName:val}}
        execOrder = append(execOrder, *newOrder)
      }
    }
    index += 1
  }

  if orderNumber > 0 {
    fmt.Println(make([]Inventory,0))
    wg.Done()
    return
  }

  fmt.Println(execOrder)
  wg.Done()
  return
}
