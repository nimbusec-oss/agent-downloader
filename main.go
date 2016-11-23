package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cumulodev/nimbusec"
)

var api *nimbusec.API

// api reference: https://kb.nimbusec.com/API/API#agents
func main() {
	key := flag.String("key", "", "API key")
	secret := flag.String("secret", "", "API secret")
	agentArch := flag.String("arch", "64bit", "64bit or 32bit")
	agentOS := flag.String("os", "linux", "linux or windows")
	agentVersion := flag.Int("version", 13, "version")
	agentFormat := flag.String("format", "bin", "tar.gz | zip | bin")
	headless := flag.Bool("headless", false, "true if you know what you do")
	flag.Parse()

	// setup nimbusec api
	var err error
	api, err = nimbusec.NewAPI(nimbusec.DefaultAPI, *key, *secret)
	if err != nil {
		log.Fatal(err)
	}

	// download agent if you don't need guidance
	if *headless {
		download(nimbusec.Agent{
			OS:      *agentOS,
			Arch:    *agentArch,
			Version: *agentVersion,
			Format:  *agentFormat,
		})
		return
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

	err = download(agents[i])
	if err != nil {
		log.Fatal(err)
	}
}

// downloads the agent as requested to the current folder
func download(agent nimbusec.Agent) error {
	fmt.Printf("Downloading nimbusagent-%s-%s-%d.%s", agent.OS, agent.Arch, agent.Version, agent.Format)
	b, err := api.DownloadAgent(agent)
	if err != nil {
		return err
	}
	f, err := os.Create(fmt.Sprintf("nimbusagent-%s-%s-%d.%s", agent.OS, agent.Arch, agent.Version, agent.Format))
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}
