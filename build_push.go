package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jhoonb/archivex"
)

// BuildImage asdf
func buildImage() {
	// Run in directory where Dockerfile is found
	// os.Chdir("build-dir")

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err, " :unable to init client")
	}

	// Image Build requiresa tar file
	tar := new(archivex.TarFile)
	tar.Create("dockerfile.tar")
	tar.AddAll(".", true)
	tar.Close()

	// Use tar file as docker context
	dockerBuildContext, err := os.Open("dockerfile.tar")
	defer dockerBuildContext.Close()
	options := types.ImageBuildOptions{
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
		Tags:           []string{"latest"},
		Dockerfile:     "Dockerfile",
		Context:        tar,
	}
	buildResponse, err := cli.ImageBuild(context.Background(), dockerBuildContext, options)
	defer buildResponse.Body.Close()
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}

	// Copy out response of stream
	_, err = io.Copy(os.Stdout, buildResponse.Body)

	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}
}

//writes from the build response to the log
func writeToLog(reader io.ReadCloser) error {
	defer reader.Close()
	rd := bufio.NewReader(reader)
	for {
		n, _, err := rd.ReadLine()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		log.Println(string(n))
	}
	return nil
}

func main() {
	buildImage()
}
