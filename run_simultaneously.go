package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatalln("need more arguments")
	}

	var wg sync.WaitGroup

	for i := 2; i < len(os.Args); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cmd := os.Args[1]
			arg := os.Args[i]

			c := exec.Command(cmd, arg)
			if err := c.Start(); err != nil {
				fmt.Printf("%v %v ERROR: %v\n", cmd, arg, err)
				return
			}

			if err := c.Wait(); err != nil {
				fmt.Printf("%v %v ERROR: %v\n", cmd, arg, err)
			} else {
				fmt.Printf("%v %v done\n", cmd, arg)
			}

		}(i)
	}

	wg.Wait()
	fmt.Println("done")
}
