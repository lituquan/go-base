
//reference: https://www.cnblogs.com/lanyangsh/p/9806311.html

package main

import (
"encoding/json"
"fmt"
)

type Bird struct {
	A   map[string]string   `json:"a"`
}

func (bd *Bird) MarshalJSON() ([]byte, error) {
	l := []string{}
	for _,v := range bd.A {
		l = append(l,v)
	}

	return json.Marshal(l)
}

func (bd *Bird) UnmarshalJSON(b []byte) error {
	l := []string{}
	err := json.Unmarshal(b, &l)
	if err != nil {
		return err
	}

	for i,v := range l {
		k := fmt.Sprintf("%d", i)
		bd.A[k] = v
	}

	return nil
}


func main() {

	m := map[string]string{"1": "110", "2":"120", "3":"119"}
	xiQue := &Bird{A:m}

	xJson, err := json.Marshal(xiQue)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
	}

	fmt.Println("xJson:", string(xJson))

	b := `["apple", "orange", "banana"]`

	baoXiNiao := &Bird{A:map[string]string{}}
	err = json.Unmarshal([]byte(b), baoXiNiao)
	if err != nil {
		fmt.Println("json.Unmarshal failed:", err)
	}

	fmt.Println("baoXiNiao:", baoXiNiao)
}