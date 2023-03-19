package main

import "fmt"

func main() {
	arr := []string{"Palak", "Sabahat", "Bindu", "Supriya"}
	freq := make(map[string]int)
	for _, num := range arr {
		freq[num] ++
	}

	fmt.Println("Frequency of the Array is:", freq)

	for j := 0; j <= len(arr); j++ {
		name := arr[j]
		letco := make(map[string]int)
		for _, char := range name {
			letco[string(char)]++
		}
		fmt.Println(letco)
	}
}
