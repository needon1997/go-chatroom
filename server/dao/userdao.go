package dao

import (
	"encoding/json"
	"fmt"

	"../Entity"
	"../util"
	"github.com/garyburd/redigo/redis"
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao() (userDao *UserDao) {
	userDao = new(UserDao)
	userDao.pool = util.NewPool()
	return
}

func (this *UserDao) FindUser(user *Entity.User) (result *Entity.User) {
	result = new(Entity.User)
	conn := this.pool.Get()
	queryResult, err := redis.String(conn.Do("get", user.Username))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	json.Unmarshal([]byte(queryResult), result)
	return
}

func (this *UserDao) AddUser(user *Entity.User) bool {
	userByte, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return false
	}
	conn := this.pool.Get()
	_, err = conn.Do("set", user.Username, string(userByte))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
