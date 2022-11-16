package main

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

func main() {
	const path = "/property/project/project_locations/[1]/geo_location/[1]"

	elems := strings.Split(path, "/")
	elems = elems[1:] // remove empty

	var (
		obj any = "Hanoi"
	)

	for i := len(elems) - 1; i >= 0; i-- {
		if strings.Contains(elems[i], "[") {
			strIdx := strings.ReplaceAll(strings.ReplaceAll(elems[i], "[", ""), "]", "")
			idx, err := strconv.Atoi(strIdx)
			if err != nil {
				panic(err)
			}
			tmp := make([]any, idx+1)
			tmp[idx] = obj
			obj = tmp
			continue
		}
		tmp := make(map[string]any)
		tmp[elems[i]] = obj
		obj = tmp
	}

	data, _ := json.MarshalIndent(obj, "", "    ")
	log.Printf("\n%s\n", string(data))

	/*
		{
			"property": {
				"project": {
					"project_locations": [
						null,
						{
							"geo_location": [
								null,
								"Hanoi"
							]
						}
					]
				}
			}
		}
	*/
}

/*
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
*/
/*
paths[0] = property
paths[1] = project
paths[2] = project_locations
paths[3] = [0]
paths[4] = geo_location
paths[5] = [0]
Hanoi
*/
