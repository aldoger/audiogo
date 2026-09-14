package service

const NUCLEAR_BASE = "http://192.168.1.42:4120/api"

type NuclearPlayer struct{}

type HealthCheckJSON struct {
	Status string `json:"status"`
}

func (np *NuclearPlayer) HealthCheck() HealthCheckJSON {
	return HealthCheckJSON{}
}
