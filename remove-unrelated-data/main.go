package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	delete(dataMap, key)
	return dataMap
}

func main() {
	data := map[string]any{
		"satu":  "ichi",
		"dua":   "ni",
		"tiga":  "san",
		"empat": "yon",
		"lima":  "go",
	}
	fmt.Println("Data awal:", data)

	keyToRemove := "tiga"
	result := removeUnrelated(data, keyToRemove)
	fmt.Printf("Data setelah key '%s' dihapus: %v\n", keyToRemove, result)
}
