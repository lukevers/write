package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

var (
	// Stderr is a standard logger prints out to stderr
	Stderr = log.New(os.Stderr, "", log.LstdFlags)

	// Stdout is a standard logger that prints to stdout
	Stdout = log.New(os.Stdout, "", log.LstdFlags)
)

func init() {
	logs := Storage.NewStorage("logs")

	Stdout.SetOutput(&lumberjack.Logger{
		Filename: logs.GetPath() + "/stdout.log",
		MaxSize:  10,
	})

	Stderr.SetOutput(&lumberjack.Logger{
		Filename: logs.GetPath() + "/stderr.log",
		MaxSize:  10,
	})
}
