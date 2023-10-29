package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Images []ImageConfig `yaml:"Images"`
}

type ImageConfig struct {
	ImageName string `yaml:"imageName"`
	ImageTag  string `yaml:"imageTag"`
}

func shaGetter() {
	// Open the configuration file.
	// change this to take variable input so cobra command can send file
	fmt.Print("Enter the path to the YAML file: ")
	var filePath string
	fmt.Scanln(&filePath)
	configFile, err := os.Open("approved_repos.yaml")
	if err != nil {
		fmt.Println("Error opening YAML file:", err)
		return
	}
	defer configFile.Close()

	// Read the configuration from the YAML file.
	config := Config{}
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding YAML:", err)
		return
	}

	// Process each Docker image in the list.
	for _, image := range config.Images {
		// Form the URL for the Docker Hub API using the values from the YAML file.
		apiURL := fmt.Sprintf("https://registry-1.docker.io/v2/%s/manifests/%s", image.ImageName, image.ImageTag)

		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			fmt.Println("Error creating HTTP request:", err)
			return
		}

		// Set the User-Agent header to avoid rate limiting issues.
		req.Header.Add("User-Agent", "Docker-Image-Info-Scraper")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error sending HTTP request:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error: Status code %d\n", resp.StatusCode)
			return
		}

		// Retrieve and print the Docker image SHA from the Docker-Content-Digest header.
		imageSHA := resp.Header.Get("Docker-Content-Digest")
		fmt.Printf("Docker Image SHA for %s:%s: %s\n", image.ImageName, image.ImageTag, imageSHA)
	}
}
