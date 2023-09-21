package account

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	e "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"github.com/quocbang/data-flow-sync/server/utils/roles"
)

type service struct {
	pg *gorm.DB
}

func NewService(pg *gorm.DB) repositories.AccountServices {
	return service{
		pg: pg,
	}
}

// getAccount definition.
func (s service) GetAccount(ctx context.Context, Identifier string) (models.Account, error) {
	user := models.Account{}
	err := s.pg.Where(`user_id=?`, Identifier).Take(&user).Error
	if err != nil {
		err = s.pg.Where(`email=?`, Identifier).Take(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.Account{}, e.Error{
					Code:    e.Code_ACCOUNT_NOT_FOUND,
					Details: "account not found",
				}
			}
			return models.Account{}, err
		}
	}
	return user, nil
}

// CreateAccount is create new account with unspecified role.
func (s service) createAccount(ctx context.Context, req repositories.CreateAccountRequest) (repositories.CreateAccountReply, error) {
	// hash password
	pwd, err := models.ToHashPassword(req.Password)
	if err != nil {
		return repositories.CreateAccountReply{}, e.Error{
			Details: fmt.Sprintf("failed to generate password, error: %v", err.Error()),
		}
	}

	reply := s.pg.Create(&models.Account{
		UserID:   req.UserID,
		Email:    req.Email,
		Password: pwd,
		Roles:    roles.Roles_UNSPECIFIED,
	})

	return repositories.CreateAccountReply{RowsAffected: reply.RowsAffected}, reply.Error
}

// DeleteAccount is delete existing account
func (s service) DeleteAccount(ctx context.Context, req repositories.DeleteAccountRequest) (repositories.CommonUpdateAndDeleteReply, error) {
	reply := s.pg.Where(`user_id=?`, req.UserID).Delete(&models.Account{})
	return repositories.CommonUpdateAndDeleteReply{RowsAffected: reply.RowsAffected}, reply.Error
}

// UpdateRole updates the role for specified account
func (s service) UpdateToUserRole(ctx context.Context, email string) (repositories.CommonUpdateAndDeleteReply, error) {
	reply := s.pg.Model(&models.Account{}).Where(`email = ?`, email).Update("roles", roles.Roles_USER)
	return repositories.CommonUpdateAndDeleteReply{RowsAffected: reply.RowsAffected}, reply.Error
}

func (s service) SignUp(ctx context.Context, req repositories.SignUpAccountRequest) error {
	_, err := s.createAccount(ctx, repositories.CreateAccountRequest{
		UserID:   req.UserID,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return err
	}
	return nil
}
