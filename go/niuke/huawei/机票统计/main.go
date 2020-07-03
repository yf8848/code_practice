package main

import "fmt"

func main() {
	var n int
	var name string
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}
		ticket := make(map[string]int, n+1)
		seq := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&name)
			ticket[name] = 0
			seq[i] = name
		}
		fmt.Scan(&n)
		for i := 0; i < n; i++ {
			fmt.Scan(&name)
			if _, b := ticket[name]; b {
				ticket[name]++
			} else {
				ticket["Invalid"]++
			}
		}

		for _, name := range seq {
			fmt.Printf("%s : %d\n", name, ticket[name])
		}
		fmt.Printf("Invalid : %d\n", ticket["Invalid"])
	}
}
