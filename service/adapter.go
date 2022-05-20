package service

import (
	"fmt"

	"github.com/arifwidiasan/todo-app/config"
	"github.com/arifwidiasan/todo-app/domain"
	"github.com/golang-jwt/jwt"
)

type svcUser struct {
	c    config.Config
	repo domain.AdapterRepository
}

func (s *svcUser) ClaimToken(bearer *jwt.Token) string {
	claim := bearer.Claims.(jwt.MapClaims)
	username := fmt.Sprintf("%v", claim["username"])

	return username
}

func NewServiceUser(repo domain.AdapterRepository, c config.Config) domain.AdapterService {
	return &svcUser{
		repo: repo,
		c:    c,
	}
}
