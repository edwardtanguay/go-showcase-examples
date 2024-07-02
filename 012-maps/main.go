package main

import "fmt"

func display(m map[string]string) {
	fmt.Println("---")
	for key, value := range m {
		fmt.Printf("%v: \"%v\"\n", key, value)
	}
}

func main() {
	frameworks := map[string]string{
		"svelte":  "Svelte",
		"angular": "Angular",
		"solid": "SolidJS",
	}
	frameworks["react"] = "React"
	frameworks["nextjs"] = "Next.js"
	frameworks["vue"] = "Vue.js"
	
	display(frameworks)

	delete(frameworks, "solid")

	display(frameworks)

	
}
