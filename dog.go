package animals

import (
    "fmt"
    "github.com/liubinpy/food"
)

func Bark() string {
    fmt.Println("吃了" + food.GetFood()) 
    return "旺旺旺"
}
