package service

import (
	"github.com/projects/loans/domain/link"
	"github.com/projects/loans/domain/link/dto"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type ILinkService interface {
	GetLinkService(link.Link) (*dto.LinkResponse, utilerrors.RestErr)
}
