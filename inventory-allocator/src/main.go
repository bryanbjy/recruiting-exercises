package main
import (
  "parser"
  "log"
  inventoryAPI "inventory"
  "sync"
)

var wg sync.WaitGroup

func main(){
  order,inventory,err := parser.Parse("test0.json")
  if err != nil {
    log.Fatal(err)
  }

  wg := &sync.WaitGroup{}
  inventoryAllocator := &inventoryAPI.InventoryAllocator{}
  for k, v := range(order){
     wg.Add(1)
    go inventoryAllocator.SetOrder(k,v,inventory,wg)
  }

  wg.Wait()

}
