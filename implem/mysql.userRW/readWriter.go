package userRW

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"log"

	"time"

	"github.com/saeidraei/go-jwt-auth/domain"
	"github.com/saeidraei/go-jwt-auth/uc"
	"github.com/spf13/viper"
)

type rw struct {
	db *sql.DB
}

func New() uc.UserRW {
	db, err := sql.Open("mysql", viper.GetString("mysql.user")+":"+viper.GetString("mysql.password")+"@tcp("+viper.GetString("mysql.host")+":"+viper.GetString("mysql.port")+")/"+viper.GetString("mysql.database"))
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	return rw{
		db: db,
	}
}

func (rw rw) Create(username, email, password string) (*domain.User, error) {
	if _, err := rw.GetByName(username); err == nil {
		return nil, uc.ErrAlreadyInUse
	}

	ins, err := rw.db.Query("insert into url(ID,Address) values(?,?)", url.ID, url.Address)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer ins.Close()

	return rw.GetByName(url.ID)
}

func (rw rw) GetByName(userName string) (*domain.User, error) {
	value, ok := rw.store.Load(userName)
	if !ok {
		return nil, uc.ErrNotFound
	}

	user, ok := value.(domain.User)
	if !ok {
		return nil, errors.New("not a user stored at key")
	}

	return &user, nil
}

func (rw rw) GetByEmailAndPassword(email, password string) (*domain.User, error) {
	var err error
	var foundUser domain.User

	rw.store.Range(func(key, value interface{}) bool {
		user, ok := value.(domain.User)
		if !ok {
			err = errors.New("failed to assert to domain.User")
			return false
		}

		if user.Email == email && user.Password == password {
			foundUser = user
			return false // stop range
		}

		return true // keep iterating
	})

	return &foundUser, err
}

func (rw rw) Save(user domain.User) error {
	if user, _ := rw.GetByName(user.Name); user == nil {
		return uc.ErrNotFound
	}

	user.UpdatedAt = time.Now()
	rw.store.Store(user.Name, user)

	return nil
}
