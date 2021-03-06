package student

import (
	"source/my-clean-arch/entity"
)

type studentRepository interface {
	GetAllById(ids []uint64) (*entity.ListStudent, error)
	Create(s *entity.Student) error
	Update(s *entity.Student) error
	Delete(id uint64) error
}
