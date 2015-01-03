package main

import (
	"log";
	"path/filepath";
	"os";
	"flag";
	"github.com/stathat/jconfig";
)

var (
	config_path = flag.String("config_path", "", "Path to config")
	source_root = flag.String("source_root", "", "Source root path")
	dest_root = flag.String("dest_root", "", "Dest root path")
	sms_params credertial
	notify_to string
	)

func main() {
	flag.Parse()

	if (*config_path != "") {
		config := jconfig.LoadConfig(*config_path)
		*source_root = config.GetString("source_root")
		*dest_root = config.GetString("dest_root")

		if (config.GetInt("api_id") != -1) {
			sms_params.api_id = config.GetInt("api_id")
			sms_params.api_key = config.GetString("api_key")
			sms_params.api_sender = config.GetString("api_sender")
			notify_to = config.GetString("to")
		}
	}

	if (*source_root == "" || *dest_root == "") {
		log.Fatal("source_root and dest_root must be defined")
	}
	
	if _, err := os.Stat(*source_root); err != nil {
		log.Panic(err)
	}
	
	err := filepath.Walk(*source_root, WalkFunc)

	if (err != nil) {
		log.Fatal(err)
	}

	message := "Files had copied"
	SendSMS(sms_params, notify_to, message)
}
