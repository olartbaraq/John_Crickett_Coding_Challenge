package main

import (
	"errors"
	"strings"

	"github.com/olartbaraq/load_balancer/utils"
)

const PATHTOCONFIGFILE = "./files/config.json"

type ServerConf struct {
	Address     string `json:"address" validate:"required"`
	HealthCheck string `json:"healthcheck" validate:"required"`
}

type LoadBalancerConfig struct {
	Port int    `json:"port" validate:"required,hostname_port"`
	Env  string `json:"env" validate:"required,oneof=dev prod"`

	// Development Specific Configs
	NoOfServers      int  `json:"no_of_servers"`
	RandomServerDown bool `json:"randomServerDown"`

	// Production Specific Configs
	Servers []ServerConf `json:"servers"`
}

func (lbConfig *LoadBalancerConfig) LoadConfig() error {
	err := utils.LoadFile[LoadBalancerConfig](PATHTOCONFIGFILE, lbConfig)

	utils.OnPanicError(err, "A problem occurred while loading config")

	err = validate.Struct(lbConfig)
	utils.OnPanicError(err, "error occurred validating config")

	if (strings.ToLower(config.Env) == "dev") && (lbConfig.NoOfServers <= 0) {
		utils.OnPanicError(errors.New("number of servers in dev env must be greater than zero"), "")
	}

	if (strings.ToLower(config.Env) == "prod") && (len(lbConfig.Servers) == 0) {
		utils.OnPanicError(errors.New("at least one server is required to redirect traffic"), "")
	}

	return nil
}
