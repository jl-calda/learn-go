package main

import "fmt"

func keys(m map[string]string) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
func maps() {
	websites := map[string]string{
		"Google": "https://www.google.com",
		"Azure":  "https://www.azure.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])
	fmt.Println(keys(websites))
}
