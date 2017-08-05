package handlers

import (
	"context"
	"time"

	sql "github.com/myambition/ambition/services/model/internal/mysql"
	pb "github.com/myambition/ambition/services/model/model-service"

	"github.com/myambition/ambition/internal/logger"

	"github.com/adamryman/kit/dbconn"
	"github.com/pkg/errors"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.ModelServer {
	logger.Init("service", "model")

	database, err := sql.Open(dbconn.FromENV("MYSQL").MySQL())
	if err != nil {
		// TODO: Do not panic, start something to try connection over and over.
		// Maybe 100 times?
		// DEBUG_SVC=1 then do like 3.
		// There will also need to be retry logic for the database methods
		panic(err)
	}
	return modelService{
		db: database,
	}
}

type modelService struct {
	db *sql.Database
}

// CreateAction implements Service.
func (s modelService) CreateAction(ctx context.Context, in *pb.Action) (*pb.Action, error) {
	// TODO: Input validation
	a, err := s.db.CreateAction(in)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create action")
	}
	return a, nil
}

// CreateOccurrence implements Service.
func (s modelService) CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.Occurrence, error) {
	utc7, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return nil, errors.Wrap(err, "cannot create time location UTC-7")
	}
	nowutc := time.Now()
	logger.Debug().Log("time", nowutc)
	now := nowutc.In(utc7)
	logger.Debug().Log("time", now)

	occurrence := in.GetOccurrence()
	if occurrence == nil {
		return nil, errors.New("cannot create nil occurrence")
	}
	//if occurrence.GetDatetime() == "" {
	//occurrence.Datetime = now
	//}

	action, err := s.db.ReadActionByID(occurrence.GetActionID())
	if err != nil {
		return nil, errors.Wrap(err, "cannot read action")
	}
	if action.GetUserID() != in.GetUserID() {
		logger.Debug().Log("action-user-id", action.GetUserID(), "user-id", in.GetUserID())
		return nil, errors.New("cannot create occurrence for action not owned by user")
	}

	o, err := s.db.CreateOccurrence(action.GetID(), now, in.GetOccurrence().GetData())
	if err != nil {
		return nil, errors.Wrap(err, "cannot create occurrence")
	}
	return o, nil
}

// ReadAction implements Service.
func (s modelService) ReadAction(ctx context.Context, in *pb.Action) (*pb.Action, error) {
	if in.GetID() != 0 {
		a, err := s.db.ReadActionByID(in.GetID())
		if err != nil {

			return nil, errors.Wrap(err, "cannot read action")
		}
		return a, nil
	}
	if name, userID := in.GetName(), in.GetUserID(); name != "" && userID != 0 {
		a, err := s.db.ReadActionByNameAndUserID(name, userID)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read action")
		}
		return a, nil
	}
	return nil, errors.New("cannot read action, need ID or BOTH UserID and Name")
}

// ReadActions implements Service.
// TODO: Implement
func (s modelService) ReadActions(ctx context.Context, in *pb.User) (*pb.ActionsResponse, error) {
	// TODO: Input validation
	var resp pb.ActionsResponse
	resp = pb.ActionsResponse{
	// Actions:
	}
	return &resp, nil
}

// ReadOccurrences implements Service.
// TODO: Implement
func (s modelService) ReadOccurrences(ctx context.Context, in *pb.Action) (*pb.OccurrencesResponse, error) {
	// TODO: Input validation
	var resp pb.OccurrencesResponse
	resp = pb.OccurrencesResponse{
	// Occurrences:
	}
	return &resp, nil
}

// ReadOccurrencesByDate implements Service.
func (s modelService) ReadOccurrencesByDate(ctx context.Context, in *pb.OccurrencesByDateReq) (*pb.OccurrencesResponse, error) {
	var resp pb.OccurrencesResponse
	if in.GetActionID() != 0 {

	}
	resp = pb.OccurrencesResponse{
	// Occurrences:
	}
	return &resp, nil
}
