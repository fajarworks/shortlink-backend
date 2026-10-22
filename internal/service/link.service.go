package service

import (
	"context"

	"github.com/fajarworks/shortlink-backend/internal/dto"
	"github.com/fajarworks/shortlink-backend/internal/model"
	"github.com/fajarworks/shortlink-backend/internal/repository"
	"github.com/fajarworks/shortlink-backend/internal/utils"
)

type LinkService struct {
	repo *repository.LinkRepo
}

func NewLinkService(repo *repository.LinkRepo) *LinkService {
	return &LinkService{
		repo: repo,
	}
}

func (s *LinkService) CreateLink(ctx context.Context, userID int, link dto.CreateShortlinkRequest) (dto.ShortlinkResponse, error) {
	slug := link.Slug

	if slug != "" {
		valid, err := utils.ValidateSlug(slug)
		if err != nil {
			return dto.ShortlinkResponse{}, err
		}

		if !valid {
			return dto.ShortlinkResponse{}, utils.ErrFormatSlug
		}
	} else {
		slug = utils.GenerateSlug(8)
	}

	newLink := model.Link{
		UserId:      userID,
		Slug:        slug,
		OriginalURL: link.DestinationURL,
	}

	data, err := s.repo.CreateLink(ctx, newLink)

	if err != nil {
		return dto.ShortlinkResponse{}, err
	}

	return dto.ShortlinkResponse{
		ID:             newLink.Id,
		Slug:           newLink.Slug,
		DestinationURL: newLink.OriginalURL,
		ClickCount:     0,
		CreatedAt:      data.CreatedAt,
	}, nil
}

func (s *LinkService) GetLinksByUserID(ctx context.Context, userID int) ([]dto.ShortlinkResponse, error) {
	links, err := s.repo.FindLinksByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ShortlinkResponse, 0, len(links))

	for _, link := range links {
		responses = append(responses, dto.ShortlinkResponse{
			ID:             link.Id,
			Slug:           link.Slug,
			DestinationURL: link.OriginalURL,
			ClickCount:     link.ClickCount,
			CreatedAt:      link.CreatedAt,
		})
	}

	return responses, nil
}
