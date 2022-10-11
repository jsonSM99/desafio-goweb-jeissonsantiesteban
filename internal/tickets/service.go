package tickets

import (
	"context"
	"desafio-goweb-jeissonsantiesteban/internal/domain"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
	ts, err := s.repository.GetTicketByDestination(ctx, destination)

	if err != nil {
		return []domain.Ticket{}, err
	}

	return ts, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {

	totalTs, err := s.repository.GetAll(ctx)

	if err != nil {
		return 0, err
	}

	ts, err := s.repository.GetTicketByDestination(ctx, destination)

	if err != nil {
		return 0, err
	}

	if len(totalTs) == 0 {
		return 0, nil
	}

	return (float64(len(ts)) / float64(len(totalTs))) * 100, nil
}
