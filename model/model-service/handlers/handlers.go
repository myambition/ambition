package handlers

import (
	"golang.org/x/net/context"

	pb "github.com/adamryman/ambition/model/model-service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.ModelServer {
	return modelService{}
}

type modelService struct{}

// CreateAction implements Service.
func (s modelService) CreateAction(ctx context.Context, in *pb.Action) (*pb.Action, error) {
	var resp pb.Action
	resp = pb.Action{
	// ID:
	// Name:
	// UserID:
	}
	return &resp, nil
}

// CreateOccurrence implements Service.
func (s modelService) CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.Occurrence, error) {
	var resp pb.Occurrence
	resp = pb.Occurrence{
	// ID:
	// ActionID:
	// Datetime:
	// Data:
	}
	return &resp, nil
}

// ReadAction implements Service.
func (s modelService) ReadAction(ctx context.Context, in *pb.Action) (*pb.Action, error) {
	var resp pb.Action
	resp = pb.Action{
	// ID:
	// Name:
	// UserID:
	}
	return &resp, nil
}

// ReadActions implements Service.
func (s modelService) ReadActions(ctx context.Context, in *pb.User) (*pb.ActionsResponse, error) {
	var resp pb.ActionsResponse
	resp = pb.ActionsResponse{
	// Actions:
	}
	return &resp, nil
}

// ReadOccurrencesByDate implements Service.
func (s modelService) ReadOccurrencesByDate(ctx context.Context, in *pb.OccurrencesByDateReq) (*pb.OccurrencesResponse, error) {
	var resp pb.OccurrencesResponse
	resp = pb.OccurrencesResponse{
	// Occurrences:
	}
	return &resp, nil
}

// ReadOccurrences implements Service.
func (s modelService) ReadOccurrences(ctx context.Context, in *pb.Action) (*pb.OccurrencesResponse, error) {
	var resp pb.OccurrencesResponse
	resp = pb.OccurrencesResponse{
	// Occurrences:
	}
	return &resp, nil
}
