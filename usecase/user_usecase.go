package usecase

import (
	"context"
	"fmt"

	"github.com/sidiqPratomo/DJKI-Pengaduan/apperror"
	"github.com/sidiqPratomo/DJKI-Pengaduan/dto"
	"github.com/sidiqPratomo/DJKI-Pengaduan/repository"
	"github.com/sidiqPratomo/DJKI-Pengaduan/util"
)

type UserUsecase interface {
	IndexUser(ctx context.Context, params dto.UserQueryParams) (*dto.ResponseIndex[dto.PagedResult[dto.UserDetail]], error)
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
	fmt.Println("params", params)
	queryParam := dto.MapDTOQuerytoEntity(params)
	fmt.Println("queryParam", queryParam.Offset)
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

// func (u *userUsecaseImpl) ReadUser(userID int64) (*dto.User, error) {
// 	user, err := u.userRepository.FindByID(userID)
// 	if err != nil {
// 		return dto.User{}, err
// 	}
// 	return user, nil
// }

// func (u *userUsecaseImpl) UpdateUser(userID int64, input dto.UpdateUserRequest) (dto.UserDetail, error) {
// 	err := u.transaction.WithTransaction(func() error {
// 		if err := u.userRepository.Update(userID, input); err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return dto.UserDetail{}, err
// 	}

// 	return u.userRepository.FindByID(userID)
// }

// func (u *userUsecaseImpl) SoftDeleteUser(userID int64) error {
// 	return u.userRepository.SoftDelete(userID)
// }
