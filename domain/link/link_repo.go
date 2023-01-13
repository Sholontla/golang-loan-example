package link

import utilerrors "github.com/projects/loans/utils/util_errors"

type LinkRepository interface {
	GetLinkDao(Link) (*Link, utilerrors.RestErr)
}
