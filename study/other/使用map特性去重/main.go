package main

import "fmt"

func main() {
	slice := []string{"1", "2", "3", "3"}
	result := []string{}
	m := make(map[string]bool)

	for _, v := range slice {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	fmt.Println(result)
}