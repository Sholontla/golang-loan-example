package service

import (
	"github.com/google/uuid"
	"github.com/projects/loans/domain/admin"
	dto "github.com/projects/loans/domain/admin/dto"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type RoleAdminRepoService struct {
	repo admin.RoleAdminRepository
}

func (req RoleAdminRepoService) RoleNewAdmin(u admin.AdminRoles) (*dto.RoleAdminResponse, *utilerrors.RestErr) {

	c := admin.AdminRoles{
		AdminRolesId: uuid.New(),
		AdminRoles:   u.AdminRoles,
	}

	newAdmin, err := req.repo.RoleAdminRegistration(c)
	if err != nil {
		return nil, err
	}

	response := newAdmin.ToNewRoleAdminResponseDto()

	return &response, nil
}

func (req RoleAdminRepoService) RoleAdminUpdateInfoService(u admin.AdminRoles) (*admin.AdminRoles, error) {

	c := admin.AdminRoles{
		AdminRolesId: uuid.New(),
		AdminRoles:   u.AdminRoles,
	}

	req.repo.RoleAdminUpdateInfoDao(c)

	return &c, nil
}

func NewRoleAdminService(repo admin.RoleAdminRepository) RoleAdminRepoService {
	return RoleAdminRepoService{repo: repo}
}
