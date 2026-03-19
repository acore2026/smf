package app

import (
	smf_context "github.com/acore2026/smf/internal/context"
	"github.com/acore2026/smf/pkg/factory"
)

type App interface {
	SetLogEnable(enable bool)
	SetLogLevel(level string)
	SetReportCaller(reportCaller bool)

	Start()
	Terminate()

	Context() *smf_context.SMFContext
	Config() *factory.Config
}
