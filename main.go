package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func main() {
	data, _ := json.Marshal(context.WithValue(context.Background(), "a", "b"))
	fmt.Println(string(data))
}
