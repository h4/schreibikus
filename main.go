package main

import (
	"log";
	"path/filepath";
	"os";
	"flag";
	"github.com/stathat/jconfig";
	)

var (
	config = jconfig.LoadConfig("config.json")
	source_root = flag.String("source_root", config.GetString("source_root"), "Source root path")
	dest_root = flag.String("dest_root", config.GetString("dest_root"), "Dest root path")
	)

func main() {
	flag.Parse()
	
	if _, err := os.Stat(*source_root); err != nil {
		log.Panic(err)
	}
	
	filepath.Walk(*source_root, WalkFunc)
}
