package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"syscall"
	"time"

	"github.com/google/shlex"

	"github.com/Lucino772/envelop/internal/config"
	"github.com/Lucino772/envelop/internal/modules/core"
	"github.com/Lucino772/envelop/internal/modules/minecraft"
	"github.com/Lucino772/envelop/internal/wrapper"
)

func loadConfig(configPath string) (*config.Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	conf, err := config.LoadConfig(data)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working directory")
	}
	conf, err := loadConfig(path.Join(workingDir, "envelop.yaml"))
	if err != nil {
		log.Fatal("Failed to load config")
	}
	command, err := shlex.Split(conf.Process.Command)
	if err != nil {
		log.Fatal("Failed to parse command")
	}

	var options []wrapper.WrapperOptFunc
	options = append(
		options,
		wrapper.WithWorkingDirectory(workingDir),
		wrapper.WithGracefulTimeout(time.Duration(conf.Process.Graceful.Timeout)*time.Second),
		wrapper.WithForwardLogToEvent(),
		wrapper.WithForwardLogToStdout(),
	)
	if conf.Process.Graceful.Type == "cmd" {
		options = append(
			options,
			wrapper.WithGracefulStopCommand(conf.Process.Graceful.Options["cmd"].(string)),
		)
	} else if conf.Process.Graceful.Type == "signal" {
		options = append(
			options,
			wrapper.WithGracefulStopSignal(conf.Process.Graceful.Options["signal"].(syscall.Signal)),
		)
	}

	modules := map[string]wrapper.WrapperModule{
		"envelop.core":      core.NewCoreModule(),
		"envelop.minecraft": minecraft.NewMinecraftModule(),
	}
	for _, mod := range conf.Modules {
		if module, ok := modules[mod.Uses]; ok {
			options = append(options, wrapper.WithModule(module))
		} else {
			fmt.Printf("Failed to load module '%s'", mod.Uses)
		}
	}

	wp, err := wrapper.NewWrapper(command[0], command[1:], options...)
	if err != nil {
		log.Fatal("Error while creating wrapper")
	}
	wp.Run(context.Background())
}
