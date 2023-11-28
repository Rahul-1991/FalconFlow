package utility

import (
	"encoding/json"
	"fmt"
)

func JsonDump(data map[string]interface{}) string {
	// Marshal the map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(jsonData)
}
