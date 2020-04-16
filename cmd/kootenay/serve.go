package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/QuaererePlatform/go-kootenay/internal/server"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Short: "Start",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	c := new(server.Config)

	if err := viper.Unmarshal(c); err != nil {
		log.Fatal(err)
	}

	s, err := server.New(c)
	if err != nil {
		log.Fatal(err)
	}

	go func() { log.Fatal(s.Start())}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	log.Fatal(s.Shutdown(ctx))
}
