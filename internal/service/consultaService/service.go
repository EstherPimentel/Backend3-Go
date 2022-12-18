package consultaService

import (
	"checkpoint-2/internal/domain"
	"checkpoint-2/internal/repository/consulta"
)

type Service interface {
	Post(c domain.Consulta) (domain.ConsultaDTO, error)
	GetById(id int) (domain.ConsultaDTO, error)
	Update(id int, c domain.Consulta) (domain.ConsultaDTO, error)
	Delete(id int) error
}

type service struct {
	r consulta.Repository
}

func NewService(r consulta.Repository) Service {
	return &service{r}
}

func (s *service) Post(c domain.Consulta) (domain.ConsultaDTO, error) {
	return s.r.Post(c)
}

func (s *service) GetById(id int) (domain.ConsultaDTO, error) {
	return s.r.GetById(id)
}
func (s *service) Update(id int, c domain.Consulta) (domain.ConsultaDTO, error) {
	consulta, err := s.r.GetById(id)

	if err != nil {
		return domain.ConsultaDTO{}, err
	}

	if c.Descricao == "" {
		c.Descricao = consulta.Descricao
	}

	if c.DataConsulta == "" {
		c.DataConsulta = consulta.DataConsulta
	}

	if c.DentistaId > 0 {
		c.DentistaId = consulta.DentistaId
	}

	if c.PacienteId > 0 {
		c.PacienteId = consulta.PacienteId
	}

	return s.r.Update(id, c)
}

func (s *service) Delete(id int) error {
	return s.r.Delete(id)
}
