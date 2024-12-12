package main

import (
	"fmt"
	"log"
	"os/exec"
)

func checkResourceUsage() {
	// CPU usage
	cmd := exec.Command("top", "-bn1", "|", "grep", "\"Cpu(s)\"")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CPU Usage:", string(output))

	// Memory usage
	cmd = exec.Command("free", "-h")
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Memory Usage:", string(output))

	// Disk usage
	cmd = exec.Command("df", "-h")
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disk Usage:", string(output))
}

func main() {
	checkResourceUsage()
}
