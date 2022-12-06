package repository

import (
	"tokoku/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Role     string
	Email    string `gorm:"unique"`
	Password string
	Token    string `gorm:"-:migration;<-:false"`
}

func FromCore(du user.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID, CreatedAt: du.CreatedAt, UpdatedAt: du.UpdatedAt},
		Name:     du.Name,
		Role:     du.Role,
		Email:    du.Email,
		Password: du.Password,
		Token:    du.Token,
	}
}

func ToCore(u User) user.Core {
	return user.Core{
		ID:        u.ID,
		Name:      u.Name,
		Role:      u.Role,
		Email:     u.Email,
		Password:  u.Password,
		Token:     u.Token,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToCoreArray(ua []User) []user.Core {
	var arr []user.Core
	for _, val := range ua {
		arr = append(arr, user.Core{
			ID:        val.ID,
			Name:      val.Name,
			Role:      val.Role,
			Email:     val.Email,
			Password:  val.Password,
			Token:     val.Token,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}

	return arr
}
