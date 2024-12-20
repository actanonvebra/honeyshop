package services

import "github.com/actanonvebra/honeyshop/internal/repositories"

type LogService interface {
	LogAttack(attackType, details, ip string) error
}
type DefaultLogService struct {
	Repo repositories.LogRepository
}

func NewLogService(repo repositories.LogRepository) *DefaultLogService {
	return &DefaultLogService{Repo: repo}
}
func (s *DefaultLogService) LogAttack(attackType, details, ip string) error {
	return s.Repo.LogAttack(attackType, details, ip)
}
