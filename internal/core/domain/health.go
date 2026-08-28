package domain

import "time"

// HealthStatus representa a saúde e os metadados de tempo de execução da aplicação.
type HealthStatus struct {
	Status    string    `json:"status"`
	Version   string    `json:"version"`
	Commit    string    `json:"commit"`
	Uptime    string    `json:"uptime"`
	StartTime time.Time `json:"start_time"`
}
