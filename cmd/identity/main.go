package main

import (
	"Identity/cmd/config"
	"Identity/internal/store"
	"log"
	"os"

	"github.com/bwmarrin/snowflake"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

func run() error {
	conf, err := config.Load()
	if err != nil {
		return err
	}

	_, err = store.NewMSSQLStorage(conf)
	if err != nil {
		return err
	}

	node, err := snowflake.NewNode(1)

	node.Generate()

	return nil
}
