package usecase

import (
	"context"

	"github.com/sidiqPratomo/DJKI-Pengaduan/apperror"
	"github.com/sidiqPratomo/DJKI-Pengaduan/dto"
	"github.com/sidiqPratomo/DJKI-Pengaduan/entity"
	"github.com/sidiqPratomo/DJKI-Pengaduan/repository"
	"github.com/sidiqPratomo/DJKI-Pengaduan/util"
)

type UserUsecase interface {
	IndexUser(ctx context.Context, params dto.UserQueryParams) (*dto.ResponseIndex[dto.PagedResult[dto.UserDetail]], error)
	ReadUser(ctx context.Context, userID int) (*dto.User, error)
	UpdateUser(ctx context.Context, input dto.UpdateUserRequest) (*dto.UserDetail, error)
	SoftDeleteUser(ctx context.Context, userID int, updatedBy string) error
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
	transaction    repository.Transaction
	JwtHelper      util.JwtAuthentication
}

type UserUsecaseImplOpts struct {
	UserRepository repository.UserRepository
	Transaction    repository.Transaction
	JwtHelper      util.JwtAuthentication
}

func NewUserUsecaseImpl(opts UserUsecaseImplOpts) userUsecaseImpl {
	return userUsecaseImpl{
		userRepository: opts.UserRepository,
		transaction:    opts.Transaction,
		JwtHelper:      opts.JwtHelper,
	}
}

func (u *userUsecaseImpl) IndexUser(ctx context.Context, params dto.UserQueryParams) (*dto.ResponseIndex[dto.PagedResult[dto.UserDetail]], error) {
	queryParam := dto.MapDTOQuerytoEntity(params)

	users, count, err := u.userRepository.FindAll(ctx, queryParam)
	if err != nil {
		return nil, apperror.NewAppError(500, err, "Failed to fetch users")
	}
	var dtoUsers []dto.UserDetail
	for _, user := range users {
		dtoUserDetail := dto.UserDetail{
			Id:              user.Id,
			Nik:             user.Nik,
			Photo:           user.Photo,
			FirstName:       user.FirstName,
			LastName:        user.LastName,
			Username:        user.Username,
			Email:           user.Email,
			Gender:          user.Gender,
			Address:         user.Address,
			PhoneNumber:     user.PhoneNumber,
			EmailVerifiedAt: user.EmailVerifiedAt,
			CreatedBy:       user.CreatedBy,
			UpdatedBy:       user.UpdatedBy,
			CreatedTime:     user.CreatedTime,
			UpdatedTime:     user.UpdatedTime,
			Status:          int(user.Status),
		}
		dtoUsers = append(dtoUsers, dtoUserDetail)
	}
	result := dto.ResponseIndex[dto.PagedResult[dto.UserDetail]]{
		Status: true,
		Data: dto.PagedResult[dto.UserDetail]{
			Result: dtoUsers,
			Count:  count,
		},
		Message: "Success",
		Code:    200,
	}

	return &result, nil
}

func (u *userUsecaseImpl) ReadUser(ctx context.Context, userID int) (*dto.User, error) {
	user, err := u.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userDetails := dto.User{
		Id:              user.Id,
		StatusOTP:       &user.StatusOTP,
		Nik:             user.Nik,
		Photo:           user.Photo,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Username:        user.Username,
		Email:           user.Email,
		Gender:          user.Gender,
		Address:         user.Address,
		PhoneNumber:     user.PhoneNumber,
		EmailVerifiedAt: user.EmailVerifiedAt,
		Status:          int(user.Status),
		CreatedBy:       user.CreatedBy,
		UpdatedBy:       user.UpdatedBy,
		CreatedTime:     user.CreatedTime,
		UpdatedTime:     user.UpdatedTime,
	}
	return &userDetails, nil
}

func (u *userUsecaseImpl) UpdateUser(ctx context.Context, input dto.UpdateUserRequest) (*dto.UserDetail, error) {
	userEntity := entity.User{
		Id:          input.Id,
		StatusOTP:   *input.StatusOTP,
		Nik:         input.Nik,
		Photo:       input.Photo,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Username:    input.Username,
		Email:       input.Email,
		Gender:      input.Gender,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		UpdatedBy:   input.UpdatedBy,
		Status:      input.Status,
	}

	tx, err := u.transaction.BeginTx()
	if err != nil {
		return nil, apperror.InternalServerError(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	updateUserTx := tx.UserRepository()
	err = updateUserTx.UpdateUser(ctx, userEntity)
	if err != nil {
		tx.Rollback()
		return nil, apperror.InternalServerError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, apperror.InternalServerError(err)
	}

	updatedUser, err := u.userRepository.FindByID(ctx, int(input.Id))
	if err != nil {
		return nil, err
	}

	// Mapping to DTO
	return &dto.UserDetail{
		Id:              updatedUser.Id,
		Nik:             updatedUser.Nik,
		Photo:           updatedUser.Photo,
		FirstName:       updatedUser.FirstName,
		LastName:        updatedUser.LastName,
		Username:        updatedUser.Username,
		Email:           updatedUser.Email,
		Gender:          updatedUser.Gender,
		Address:         updatedUser.Address,
		PhoneNumber:     updatedUser.PhoneNumber,
		EmailVerifiedAt: updatedUser.EmailVerifiedAt,
		CreatedBy:       updatedUser.CreatedBy,
		UpdatedBy:       updatedUser.UpdatedBy,
		CreatedTime:     updatedUser.CreatedTime,
		UpdatedTime:     updatedUser.UpdatedTime,
		Status:          int(updatedUser.Status),
	}, nil
}

func (u *userUsecaseImpl) SoftDeleteUser(ctx context.Context, userID int, updatedBy string) error {
	tx, err := u.transaction.BeginTx()
	if err != nil {
		return apperror.InternalServerError(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	err = tx.UserRepository().SoftDeleteUser(ctx, userID, updatedBy)
	if err != nil {
		tx.Rollback()
		return apperror.InternalServerError(err)
	}

	if err := tx.Commit(); err != nil {
		return apperror.InternalServerError(err)
	}

	return nil
}
