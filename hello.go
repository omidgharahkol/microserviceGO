package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var m = map[string]string{"omid": "omd"}
	a, err := json.Marshal(m)
	fmt.Println(err)
	fmt.Println(string(a))
}
