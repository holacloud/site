package api

import (
	"github.com/fulldump/box"

	"github.com/holacloud/site/statics"
)

type JSON = map[string]any

func Build(staticsDir, version string) *box.B {

	b := box.NewBox()

	b.WithInterceptors(
		NaiveAccessLog,
		PrettyErrorInterceptor,
	)

	// Version
	b.Handle("GET", "/version", Write(version))

	// Static files
	b.Handle("GET", "/*", statics.ServeStatics(staticsDir)).
		WithInterceptors(
			IfModifiedSince(),
		)

	return b
}
