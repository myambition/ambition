package handlers

import (
	"fmt"
	"golang.org/x/net/context"
	"os"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	modelSVC "github.com/myambition/ambition/services/model/model-service"
	pb "github.com/myambition/ambition/services/rello/rello-service"
	usersSVC "github.com/myambition/ambition/services/users/users-service"

	modelClient "github.com/myambition/ambition/services/model/model-service/svc/client/grpc"
	usersClient "github.com/myambition/ambition/services/users/users-service/svc/client/grpc"

	//"github.com/davecgh/go-spew/spew"

	"github.com/myambition/ambition/internal/logger"
)

const trelloTime = "2006-01-02T15:04:05.999Z07:00"

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.RelloServer {
	logger.Init("service", "rello")
	logger.Info().Log("msg", "new_rello_service")
	// TODO: Create ambitionSVC and usersSVC clients
	// TODO switch to grpc clients
	modelConnStr := os.Getenv("MODEL_HOST") + ":" + os.Getenv("MODEL_PORT")
	//fmt.Println(modelConnStr)
	modelConn, err := grpc.Dial(modelConnStr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	m, err := modelClient.New(modelConn)
	if err != nil {
		panic(err)
	}
	usersConnStr := os.Getenv("USERS_HOST") + ":" + os.Getenv("USERS_PORT")
	//fmt.Println(usersConnStr)
	usersConn, err := grpc.Dial(usersConnStr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	u, err := usersClient.New(usersConn)
	if err != nil {
		panic(err)
	}
	logger.Info().Log("msg", "done connection")

	return relloService{
		model: m,
		users: u,
	}
}

type relloService struct {
	users usersSVC.UsersServer
	model modelSVC.ModelServer
}

// CheckListWebhook implements Service.
// TODO: Auth middleware using user trello webhookcallback
func (s relloService) CheckListWebhook(ctx context.Context, in *pb.ChecklistUpdate) (*pb.Empty, error) {
	logger.Debug().Log("action", in.GetAction())
	logger.Debug().Log("member_creator_id", in.GetAction().GetIdMemberCreator())

	logger.Debug().Log("msg", "Contacting users service")
	user, err := s.users.ReadUser(ctx,
		&usersSVC.User{
			Trello: &usersSVC.TrelloInfo{
				ID: in.GetAction().GetIdMemberCreator(),
			},
			Info: &usersSVC.UserInfo{},
		})
	if err != nil {
		return nil, err
	}

	logger.Debug().Log("switch ", in.Action.Type)
	cItem := in.Action.Data.CheckItem
	switch in.Action.Type {
	case "createCheckItem":
		// TODO: Map CheckItem.Id to Action.Id
		// TODO: Map u.Action.MemberCreator.Id to Action.UserId
		// UserId hardcoded to 1
		_ = in.Action.MemberCreator.Id

		logger.Debug().Log("userid", user.GetID())
		action, err := s.model.CreateAction(ctx,
			&modelSVC.Action{
				Name:   cItem.GetName(),
				UserID: user.GetID(),
			})
		if err != nil {
			logger.Error().Log("err", errors.Wrap(err, "cannot create action"))
			break
		}
		logger.Debug().Log("msg", "action created", "action", action.Name)
	case "updateCheckItemStateOnCard":
		if cItem.State == "incomplete" {
			logger.Debug().Log("msg", "uncheck", "check_item", cItem.Name)
			break
		}
		action, err := s.model.ReadAction(ctx,
			&modelSVC.Action{
				Name:   cItem.GetName(),
				UserID: user.GetID(),
			})
		dateString := &in.Action.Date
		date, err := time.Parse(trelloTime, *dateString)
		_ = date
		if err != nil {
			logger.Error().Log("err", errors.Wrap(err, "time parsing failed"))
			break
		}
		// TODO: log occurrence
		_, err = s.model.CreateOccurrence(ctx,
			&modelSVC.CreateOccurrenceRequest{
				UserID: user.GetID(),
				Occurrence: &modelSVC.Occurrence{
					ActionID: action.GetID(),
					// TODO: make sure this time format is fine
					Datetime: date.String(),
				},
			})
		if err != nil {
			fmt.Println(errors.Wrap(err, "error creating occurrence"))
			break
		}

		// TODO: Create Occurrence with the ActionId from CheckItem.Id
		// Maps to action id
	default:
		logger.Debug().Log("action_type ", in.Action.Type)
	}
	return nil, nil
}
