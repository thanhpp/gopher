package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const path = "property/project/project_locations/[0]/geo_location/[0]"
	var jsonResp = `
	{
		"property": {
			"project": {
				"project_locations": [
					{
						"geo_location": [
							"Hanoi"
						]
					}
				]
			}
		}
	}`
	resp := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonResp), &resp); err != nil {
		panic(err)
	}

	paths := strings.Split(path, "/")
	var temp interface{} = resp

	for i := range paths {
		fmt.Printf("paths[%d] = %s\n", i, paths[i])
		switch {
		case strings.Contains(paths[i], "["): // [int]
			k := strings.ReplaceAll(paths[i], "[", "")
			k = strings.ReplaceAll(k, "]", "")
			iK, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}

			arrTemp, ok := temp.([]interface{})
			if !ok {
				panic(fmt.Sprintf("cast []string error %+v %T", temp, temp))
			}
			temp = arrTemp[iK]

		default:
			m, ok := temp.(map[string]interface{})
			if !ok {
				panic(fmt.Sprintf("cast map[string]interface{} error %+v %T", temp, temp))
			}
			v := m[paths[i]]
			temp = v
		}
	}

	fmt.Println(temp)
}

/*
paths[0] = property
paths[1] = project
paths[2] = project_locations
paths[3] = [0]
paths[4] = geo_location
paths[5] = [0]
Hanoi
*/
