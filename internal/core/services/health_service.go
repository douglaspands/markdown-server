package services

import (
	"context"
	"time"

	"github.com/douglas/markdown-server/internal/core/domain"
	"github.com/douglas/markdown-server/internal/core/ports"
	"github.com/douglas/markdown-server/internal/version"
)

// HealthCheckService implementa a porta ports.HealthServicePort para diagnóstico da aplicação.
type HealthCheckService struct {
	startTime time.Time
}

// NewHealthCheckService instancia um novo HealthCheckService com o tempo de início fixado.
func NewHealthCheckService(startTime time.Time) ports.HealthServicePort {
	if startTime.IsZero() {
		startTime = time.Now()
	}
	return &HealthCheckService{
		startTime: startTime,
	}
}

// CheckHealth retorna o estado de integridade operacional, tempo de atividade e versão da aplicação.
func (s *HealthCheckService) CheckHealth(ctx context.Context) (*domain.HealthStatus, error) {
	if ctx != nil && ctx.Err() != nil {
		return nil, ctx.Err()
	}

	uptime := time.Since(s.startTime).Round(time.Second).String()
	verInfo := version.GetInfo()

	return &domain.HealthStatus{
		Status:    "UP",
		Version:   verInfo.Version,
		Commit:    verInfo.Commit,
		Uptime:    uptime,
		StartTime: s.startTime,
	}, nil
}
