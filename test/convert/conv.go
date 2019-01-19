package main

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
)

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

// todo : YAMLToJSON

func main() {
	fmt.Printf("Input: %s\n", s)
	var body interface{}
	if err := yaml.Unmarshal([]byte(s), &body); err != nil {
		panic(err)
	}

	body = convert(body)

	if b, err := json.Marshal(body); err != nil {
		panic(err)
	} else {
		fmt.Printf("Output: %s\n", b)
	}

	j, err := yaml.YAMLToJSON([]byte(s))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("or: %s", string(j))

}

const s = `Services:
-   Orders:
    -   ID: $save ID1
        SupplierOrderCode: $SupplierOrderCode
    -   ID: $save ID2
        SupplierOrderCode: 111111
`
