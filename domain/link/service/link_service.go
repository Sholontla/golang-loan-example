package service

import (
	"github.com/google/uuid"
	"github.com/projects/loans/domain/link"
	"github.com/projects/loans/domain/link/dto"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type LinkRepoService struct {
	repo link.LinkRepository
}

func (req LinkRepoService) GetLinkService(li link.Link) (*dto.LinkResponse, utilerrors.RestErr) {

	l := link.Link{
		LinkId:     uuid.New(),
		LinkCode:   li.LinkCode,
		CustomerId: li.CustomerId,
	}

	getLink, err := req.repo.GetLinkDao(l)
	if err != nil {
		return nil, utilerrors.NewBadRequestError("Error while trying to get Link")
	}

	response := getLink.ToNewLinkResponseDto()
	return &response, nil
}

func NewLinkService(repo link.LinkRepository) LinkRepoService {
	return LinkRepoService{repo: repo}
}
