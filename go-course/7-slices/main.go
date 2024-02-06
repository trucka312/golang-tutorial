package main

import "fmt"

func main() {
	linguagens := []string{"python", "ruby", "javascript"}
	linguagens = append(linguagens, "go")

	fmt.Println("nà ní: ", linguagens)
	fmt.Println("nà ní: ", linguagens[1:3])

	for i, linguagem := range linguagens {
		fmt.Printf(
			"ơ kìa %d ???, var: %s\n",
			i,
			linguagem,
		)
	}
}
