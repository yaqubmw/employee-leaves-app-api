package repository

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/utils/common"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(payload model.UserCredential) error
	GetByUsername(username string) (model.UserCredential, error)
	GetByUsernamePassword(username, password string) (model.UserCredential, error)
	Paging(requestPaging dto.PaginationParam) ([]model.UserCredential, dto.Paging, error)
}

type userRepository struct {
	db *gorm.DB
}

// Create implements UserRepository.
func (u *userRepository) Create(payload model.UserCredential) error {
	return u.db.Create(&payload).Error
}

// GetByUsername implements UserRepository.
func (u *userRepository) GetByUsername(username string) (model.UserCredential, error) {
	var user model.UserCredential
	err := u.db.Where("username = ?", username).First(&user).Error
	return user, err
}

// GetByUsernamePassword implements UserRepository.
func (u *userRepository) GetByUsernamePassword(username, password string) (model.UserCredential, error) {
	user, err := u.GetByUsername(username)
	if err != nil {
		return model.UserCredential{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.UserCredential{}, err
	}
	return user, nil
}

// Paging implements UserRepository.
func (u *userRepository) Paging(requestPaging dto.PaginationParam) ([]model.UserCredential, dto.Paging, error) {
	var users []model.UserCredential
	var totalRows int64

	pagination := common.GetPaginationParams(requestPaging)
	result := u.db.Model(&model.UserCredential{}).Count(&totalRows)
	if result.Error != nil {
		return nil, dto.Paging{}, result.Error
	}

	query := u.db.Model(&model.UserCredential{}).Limit(pagination.Take).Offset(pagination.Skip)
	result = query.Find(&users)
	if result.Error != nil {
		return nil, dto.Paging{}, result.Error
	}

	return users, common.Paginate(pagination.Page, pagination.Take, int(totalRows)), nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
