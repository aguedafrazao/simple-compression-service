package main 

import "fmt"
import "os/exec"

func main() {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

}
