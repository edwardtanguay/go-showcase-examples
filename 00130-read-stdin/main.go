
package main

import "bufio"
import "os"
import "fmt"

func main() {
	print("Enter choice [1,2,3]: ")
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	fmt.Printf("Choice was: %s", choice)
}
