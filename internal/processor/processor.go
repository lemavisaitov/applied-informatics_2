package processor

import (
	"github.com/lemavisaitov/applied-informatics_2/internal/fetcher"
	"github.com/lemavisaitov/applied-informatics_2/internal/validator"
)

type Processor struct {
	Fetcher   fetcher.DataFetcher
	Validator validator.Validator
}

func NewProcessor(fetcher fetcher.DataFetcher, validator validator.Validator) *Processor {
	return &Processor{Fetcher: fetcher, Validator: validator}
}

func (p *Processor) Process(source string) ([]string, error) {
	data, err := p.Fetcher.Fetch(source)
	if err != nil {
		return nil, err
	}
	return p.Validator.FindMatches(data), nil
}
