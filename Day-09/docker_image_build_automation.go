package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("docker", "build", "-t", "myapp:latest", ".", "-f", "Dockerfile.production")
	cmd.Dir = "./go-rest-api-container-example" // Set the working directory

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
