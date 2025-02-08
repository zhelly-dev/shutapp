package main

import (
	"fmt"
	"log"
	"time"
	"bufio"
	"os"
	"strings"
	"strconv"
	"flag"

	"github.com/shirou/gopsutil/process"
)

type BanEntry struct {
	Name string
	BanTime int
}

func loadBanList(filename string) ([]BanEntry, error) {
	var banList []BanEntry
	file, err := os.Open(filename)
	if err != nil {return nil, err}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		
		processName := parts[0]
		banTime, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Printf("Invalid ban time for process %s: %v",  processName, err)
			continue
		}
		banList = append(banList, BanEntry{Name: processName, BanTime: banTime * 60})
	}
	if err := scanner.Err(); err != nil {return nil, err}
	return banList, nil
}

func main() {
	banlistFile := flag.String("file", "", "Banlist")
	flag.Parse()
	banList, err := loadBanList(*banlistFile)
	if err != nil {log.Fatalf("Error loading banlist file: %v", err)}

	for _, entry := range banList {
		go func(entry BanEntry) {
			blockProcess(entry.Name, entry.BanTime)
		}(entry)
	}

	select {}
}

func blockProcess(processName string, banTime int) {
	for banTime > 0 {
		processes, err := process.Processes()
		if err != nil {
			log.Printf("Error retrieving processes: %v", err)
			return
		}

		for _, p := range processes {
			aname, err := p.Name()
			if err != nil {continue}

			if aname == processName {
				pid := p.Pid
				fmt.Printf("Blocking process: %s (PID: %d) for %d seconds\n", processName, pid, banTime)
				if err := p.Kill(); err != nil {
					log.Printf("Failed to kill process %d: %v", pid, err)
				} else {
					fmt.Printf("Successfully blocked process %d\n", pid)
				}
			}
		}

		time.Sleep(1 * time.Second)
		banTime--
		fmt.Printf("Time remaining for %s: %d seconds\n", processName, banTime)
	}
}
