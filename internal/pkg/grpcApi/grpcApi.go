package grpcApi

import (
	context "context"
	"fmt"
	"log"
	"mkn-backend/internal/pkg/config"
	"mkn-backend/internal/pkg/repository"
)

type GRPCServer struct {
	repo   *repository.Repository
	config *config.Config
}

func New(ctx context.Context) (*GRPCServer, error) {
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DataBase.Host, cfg.DataBase.User, cfg.DataBase.Password, cfg.DataBase.Name, cfg.DataBase.Port)
	repo, err := repository.New(dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &GRPCServer{
		repo:   repo,
		config: cfg,
	}, nil
}

func (s *GRPCServer) UpdateNotificationStatus(context.Context, *UpdateNotificationStatusRequest) (*UpdateNotificationStatusResponse, error) {
	return nil, nil
}

func (s *GRPCServer) GetFullNotificationInfo(ctx context.Context, req *NotificationInfoRequest) (*NotificationInfoResponse, error) {
	notification, err := s.repo.GetNotificationById(req.NotificationId)
	if err != nil {
		return nil, err
	}

	section, err := s.repo.GetSectionById(notification.SectionId.String())
	if err != nil {
		return nil, err
	}

	project, err := s.repo.GetProjectById(section.ProjectId.String())
	if err != nil {
		return nil, err
	}

	emails := []string{}
	collabs, err := s.repo.GetAllCollaborators(project.Id.String())

	for i := range collabs {
		emails = append(emails, collabs[i].Email)
	}

	user, err := s.repo.GetUserById(project.OwnerId.String())
	if err != nil {
		return nil, err
	}

	emails = append(emails, user.Email)

	return &NotificationInfoResponse{
		NotificationId:    notification.Id.String(),
		ProjectTitle:      project.Title,
		SectionTitle:      section.Title,
		NotificationTitle: notification.Title,
		Description:       notification.Description,
		Email:             emails,
	}, nil
}
