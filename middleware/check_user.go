package middleware

import (
	"practice/models"
	"practice/pkg/errors"
)

func CheckUser(user *models.User) int {
	if user.DeletedAt != nil {
		return errors.USER_DELETED
	}

	if !user.IsActive {
		return errors.INACTIVE_USER
	}

	if user.IsBlocked {
		return errors.USER_BLOCKED
	}

	return errors.SUCCESS

}

func CheckAdmin(admin *models.Admin) int {
	if admin.DeletedAt != nil {
		return errors.ADMIN_DELETED
	}

	if !admin.IsActive {
		return errors.INACTIVE_ADMIN
	}

	return errors.SUCCESS

}
