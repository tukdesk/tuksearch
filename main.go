package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/tukdesk/tuksearch/app"
)

var confPath = flag.String("c", "./misc/conf.json", "path to the config file")

func main() {
	flag.Parse()

	cfg := app.Config{}
	b, err := ioutil.ReadFile(*confPath)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(b, &cfg); err != nil {
		log.Fatalln(err)
	}

	s, err := app.New(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	if err := s.Run(); err != nil {
		log.Fatalln(err)
	}
}
