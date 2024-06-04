package employeService

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/andrei1998Front/grpc_workers_server/internal/domain/models"
)

type EmployeeService struct {
	log        *slog.Logger
	infoGetter InfoGetter
}

type InfoGetter interface {
	GetInfoByName(ctx context.Context, title string) ([]*models.EmployeeInfo, error)
}

func New(
	log *slog.Logger,
	infoGetter InfoGetter,
) *EmployeeService {
	return &EmployeeService{
		log:        log,
		infoGetter: infoGetter,
	}
}

func (s EmployeeService) GetInfoByName(ctx context.Context, title string) ([]*models.EmployeeInfo, error) {
	const op = "EmployeeService.GetInfoByName"

	log := s.log.With(
		slog.String("op", op),
		slog.String("title", title),
	)

	log.Info("get info from employee " + title)

	empInfo, err := s.infoGetter.GetInfoByName(ctx, title)

	if err != nil {
		log.Error("employee "+title+" upload failed", err)
		return []*models.EmployeeInfo{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("employee " + title + " upload was successful")

	return empInfo, nil
}
