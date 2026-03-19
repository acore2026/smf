package processor

import (
	"github.com/acore2026/smf/internal/sbi/consumer"
	"github.com/acore2026/smf/pkg/app"
)

const (
	CONTEXT_NOT_FOUND = "CONTEXT_NOT_FOUND"
)

type ProcessorSmf interface {
	app.App

	Consumer() *consumer.Consumer
}

type Processor struct {
	ProcessorSmf
}

func NewProcessor(smf ProcessorSmf) (*Processor, error) {
	p := &Processor{
		ProcessorSmf: smf,
	}
	return p, nil
}
