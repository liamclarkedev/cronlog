package main

import (
	"github.com/clarke94/cronlog/cmd"
	"github.com/clarke94/cronlog/pkg/cron"
)

func main() {
	parser := cron.New()
	rootCmd := cmd.New(parser)
	rootCmd.Execute()
}
