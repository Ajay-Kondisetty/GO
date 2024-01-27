package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var inputFile string
	flag.StringVar(&inputFile, "input_json_file", "", "Input JSON file path")
	flag.Parse()

	if inputFile == "" {
		log.Fatal("Input JSON file path is empty.")
	}

	inputJSON, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal("Error occurred while reading from JSON file: ", err)
	}

	if len(inputJSON) == 0 {
		log.Fatal("Received empty JSON data.")
	}

	var convertedJSON interface{}
	outJSON := make(map[string]interface{})
	err = json.Unmarshal(inputJSON, &convertedJSON)
	if err != nil {
		log.Fatal("Error while unmarshalling JSON: ", err)
	}

	FlattenJSON(convertedJSON, "", outJSON)
	data, _ := json.Marshal(outJSON)
	_ = os.WriteFile("output.json", data, 777)
}

func FlattenJSON(inputJSON interface{}, prefix string, outJSON map[string]interface{}) {

	switch v := inputJSON.(type) {
	case []interface{}:
		for k, v := range v {
			key := fmt.Sprintf("%s.%d", prefix, k)

			switch v := v.(type) {
			case int, string, bool, float64, []byte:
				outJSON[key] = v
			case []interface{}, map[string]interface{}:
				FlattenJSON(v, key, outJSON)
			default:
				fmt.Println("Unsupported type")
			}
		}
	case map[string]interface{}:
		for k, v := range v {
			key := prefix
			if key != "" {
				key += "."
			}
			key += k

			switch v := v.(type) {
			case int, string, bool, float64, []byte:
				outJSON[key] = v
			case []interface{}, map[string]interface{}:
				FlattenJSON(v, key, outJSON)
			default:
				fmt.Println("Unsupported type")
			}
		}
	default:
		fmt.Println("Unsupported type: ", v)
	}

	//i := inputJSON.(map[string]interface{})

	//for k, v := range i {
	//	key := prefix
	//	if key == "" {
	//		key += k
	//	} else {
	//		key += "." + k
	//	}
	//	switch v.(type) {
	//	case int, string, bool, float32, float64, []byte:
	//		outJSON[key] = v
	//	case []interface{}:
	//		val := v.([]interface{})
	//		for k1, v1 := range val {
	//			key += "." + string(rune(k1))
	//			FlatternJSON(v1, key, outJSON)
	//		}
	//	case interface{}:
	//		FlatternJSON(v, key, outJSON)
	//	default:
	//		outJSON[key] = v
	//	}
	//}
}
