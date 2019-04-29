package logic

import "fmt"

func checkTable(result []int) int {

	for i :=range result {
		fmt.Println(result[i])
	}
	if len(result) != 0 {
		return 100
	} else {
		return 0
	}

}
