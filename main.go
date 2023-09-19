package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
    // Define a handler function for your HTTP server
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Set the response header
        w.Header().Set("Content-Type", "text/plain")

        // Write the response content
        fmt.Fprintln(w, "Hello, World!")
    })

	//Version of Docker
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		// Set the response header
		w.Header().Set("Content-Type", "text/plain")
		// Write the response content
		fmt.Fprintln(w, DockerVerion())
	})

    // Start the HTTP server on port 8080
    fmt.Println("Server is running on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error:", err)
    }
}


//Create Docker client and return version 
func DockerVerion() string {
	// Create a Docker client
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        panic(err)
    }

    // List containers
    containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
    if err != nil {
        panic(err)
    }

    fmt.Println("Containers:")
    for _, container := range containers {
        fmt.Printf("ID: %s, Name: %s\n", container.ID[:10], container.Names[0])
    }

    // Pull an image (e.g., Alpine)
    image := "alpine:latest"
    reader, err := cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
    if err != nil {
        panic(err)
    }
    defer reader.Close()

    // Print the pull output
    // This is optional but can be useful to see the download progress
    _, err = io.Copy(os.Stdout, reader)
    if err != nil {
        panic(err)
    }

    fmt.Println("Image pulled:", image)

	return "Docker version: " + cli.ClientVersion()
}