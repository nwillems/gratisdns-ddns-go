package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"gopkg.in/yaml.v2"
)

// Configuration for the dns updater
type Configuration struct {
	Username      string   `yaml:"username"`
	Password      string   `yaml:"password"`
	DomainAccount string   `yaml:"domain_account"`
	Domains       []string `yaml:"domains"`
}

func SendRequest(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	decodedBody := string(body[0:2])

	if decodedBody != "OK" {
		return fmt.Errorf("Failed requesting URL %s.\nError: %s", url, string(body[:]))
	}

	return nil
}

func ReadConfiguration() Configuration {
	config := Configuration{}

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Could not read configuration file")
	}
	err = yaml.Unmarshal(yamlFile, &config)

	return config
}

func main() {
	// Read and deserialize
	config := ReadConfiguration()

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Could not read configuration file")
	}
	err = yaml.Unmarshal(yamlFile, &config)

	// TODO: Lookup IP using the myip.opendns.com method
	baseUrl := "https://admin.gratisdns.com/ddns.php"
	// For each domain
	for _, domain := range config.Domains {
		params := url.Values{}
		params.Add("u", config.Username)
		params.Add("p", config.Password)
		params.Add("d", config.DomainAccount)
		params.Add("h", domain)

		query := params.Encode()
		requestUrl := fmt.Sprintf("%s?%s", baseUrl, query)

		//   Update DNS - aka fire http request
		//  https://admin.gratisdns.com/ddns.php?u=johnson&p=password&d=example.com&h=fooo.example.com&i=1.1.1.1
		err := SendRequest(requestUrl)
		if err != nil {
			log.Fatalf("Failed updating Record for domain: %s", domain)
		}
	}

	print("All domains updated sucesfully\n")
	// Report status
}
