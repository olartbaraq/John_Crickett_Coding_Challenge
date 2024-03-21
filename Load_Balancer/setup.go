package main

var config LoadBalancerConfig

func setupLoadBalancer() *LoadBalancer {
	_ = config.LoadConfig()

	lb := LoadBalancer{
		Port:     config.Port,
		Count:    0,
		LastPort: config.Port,
		Config:   config,
	}

	if config.Env == "dev" {
		devSetup(&lb)
	} else if config.Env == "prod" {
		ProdSetup(&lb)
	} else {
		panic("Invalid ENV value")
	}

	return &lb
}

func ProdSetup(loadBalancer *LoadBalancer) {
	panic("unimplemented")
}

func devSetup(loadBalancer *LoadBalancer) {
	panic("unimplemented")
}
