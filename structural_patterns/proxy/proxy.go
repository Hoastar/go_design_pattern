package proxy

import (
	"log"
	"time"
)

// IUser 接口
type IUser interface {
	Login(username, password string) error
}

// User 用户
type User struct{}

// Login 用户登录
func (u User) Login(username, password string) error {
	return nil
}

// UserProxy 代理类
type UserProxy struct {
	user *User
}

// NewUserProxy
func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

// Login登录，和user实现相同的接口
func (p *UserProxy) Login(username, password string) error {
	start := time.Now()

	if err := p.user.Login(username, password); err != nil {
		return nil
	}

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return nil
}
