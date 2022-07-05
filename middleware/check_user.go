package middleware

import (
	"practice/models"
	"practice/pkg/errors"
)

func isActive(user *models.User) bool {
	return user.IsActive
}

func isBlocked(user *models.User) bool {
	return user.IsBlocked
}

func isDeleted(user *models.User) bool {
	return user.DeletedAt != nil
}

func CheckUser(user *models.User) int {
	if isDeleted(user) {
		return errors.USER_DELETED
	}

	if !isActive(user) {
		return errors.INACTIVE_USER
	}

	if isBlocked(user) {
		return errors.USER_BLOCKED
	}

	return errors.SUCCESS

}
