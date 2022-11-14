package repositories

import (
	"context"
	"ecommerce/entities"
	"ecommerce/helpers"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserRepository interface {
	GetUser(id string) (entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
	FindByIdUser(id string) (entities.User, error)
	Update(user entities.User) (entities.User, error)
	GetUsersWithCache() ([]entities.User, error)
	FindAllUsersPaginate(pagination utils.Pagination, queryFilter string) ([]*entities.User, utils.Pagination)
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUser(id string) (entities.User, error) {
	var user entities.User
	findUser := r.db.Where("id =?", id).First(&user).Error
	if findUser != nil {
		return user, findUser
	}
	return user, nil
}
func (r *userRepository) FindByEmail(email string) (entities.User, error) {

	var user entities.User
	findUser := r.db.Where("email =?", email).First(&user).Error
	if findUser != nil {
		return user, findUser
	}
	return user, nil
}
func (r *userRepository) FindByIdUser(id string) (entities.User, error) {

	var user entities.User
	findUser := r.db.Preload("Role").Preload("Role.MenuAccess").Preload("Role.MenuAccess.Menu").Where("id =?", id).First(&user).Error
	if findUser != nil {
		return user, findUser
	}
	return user, nil
}
func (r *userRepository) CreateUser(user entities.User) (entities.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *userRepository) Update(user entities.User) (entities.User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *userRepository) GetUsersWithCache() ([]entities.User, error) {
	var (
		users []entities.User
	)

	rdb := helpers.NewRedisClient()
	fmt.Println("redis client initialized", rdb)
	ctx := context.Background()

	getUsers, errGet := rdb.Get(ctx, "users").Bytes()
	fmt.Println(getUsers, "[][]][]get users dari redis")

	if errGet != nil {
		fmt.Println("unable to GET data. error: %v", errGet)
		err := r.db.Find(&users).Error
		if err != nil {
			return users, err
		}
		cachedProducts, errCache := json.Marshal(users)
		if errCache != nil {
			return users, errCache
		}
		log.Println("larinya kesini ke db")
		errSet := rdb.Set(ctx, "users", cachedProducts, 10*time.Second).Err()
		if errSet != nil {

			return nil, errSet

		}
		return users, nil
		fmt.Println("ke cachhe")
	}

	//

	return users, nil
}

func (r *userRepository) FindAllUsersPaginate(pagination utils.Pagination, queryFilter string) ([]*entities.User, utils.Pagination) {
	var users []*entities.User

	r.db.Scopes(utils.Paginate(&entities.User{}, &pagination, r.db)).Where(queryFilter).Preload("Role").Find(&users)

	return users, pagination
}
