package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cumulodev/nimbusec"
)

// api reference: https://kb.nimbusec.com/API/API#agents
func main() {
	key := flag.String("key", "", "API key")
	secret := flag.String("secret", "", "API secret")
	flag.Parse()

	// setup nimbusec api
	api, err := nimbusec.NewAPI(nimbusec.DefaultAPI, *key, *secret)
	if err != nil {
		log.Fatal(err)
	}

	// fetch available nimbusec agents
	agents, err := api.FindAgents(nimbusec.EmptyFilter)
	if err != nil {
		log.Fatal(err)
	}
	for idx, agent := range agents {
		fmt.Printf("%d: nimbusagent-%s-%s-%d.%s\n", idx, agent.OS, agent.Arch, agent.Version, agent.Format)
	}

	fmt.Printf("Please select the Agent to download (0 - %d): ", len(agents)-1)
	idx := ""
	fmt.Scanln(&idx)
	i, err := strconv.Atoi(idx)
	if err != nil {
		log.Fatal(err)
	}
	if i >= len(agents) {
		log.Fatal("Index out of range")
	}

	// prepare download over api
	b, err := api.DownloadAgent(agents[i])
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(fmt.Sprintf("nimbusagent-%s-%s-%d.%s", agents[i].OS, agents[i].Arch, agents[i].Version, agents[i].Format))
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}
