package services

import (
	"github.com/rapidstellar/gohexa/internal/core/domain"
	"github.com/rapidstellar/gohexa/internal/core/ports"
)

type GeneratorServiceImpls struct {
	flag domain.GeneratorFlagDomain
}

func NewGeneratorService(flag domain.GeneratorFlagDomain) ports.IGeneratorService {
	return &GeneratorServiceImpls{flag}
}
