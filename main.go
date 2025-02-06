package main

import (
	"fmt"
	"log"
	"flag"
	"time"

	"github.com/shirou/gopsutil/process"
)

func main() {
	name := flag.String("name", "", "Process name")
	blockTime := flag.Int("time", 1, "Time to block the process")
	flag.Parse()
	*blockTime *= 60

	for *blockTime <= 0 {
		for {
			processes, err := process.Processes()
			if err != nil {
				log.Fatalf("Error retrieving processes: %v", err)
			}

			for _, p := range processes {
				aname, err := p.Name()
				if err != nil {
					continue
				}

				if aname == *name {
					pid := p.Pid
					fmt.Printf("Killing process: %s (PID: %d)\n", *name, pid)
					if err := p.Kill(); err != nil {
						log.Printf("Failed to kill process %d: %v", pid, err)
					} else {
						fmt.Printf("Successfully killed process %d\n", pid)
					}
				}
			}
			time.Sleep(1 * time.Second)
			*blockTime--;
			fmt.Println(*blockTime)
		}
	}
}
