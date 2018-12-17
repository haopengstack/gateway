package proxy

import (
	"github.com/fagongzi/gateway/pkg/filter"
)

// XForwardForFilter XForwardForFilter
type XForwardForFilter struct {
	filter.BaseFilter
}

func newXForwardForFilter() filter.Filter {
	return &XForwardForFilter{}
}

// Init init filter
func (f *XForwardForFilter) Init(cfg string) error {
	return nil
}

// Name return name of this filter
func (f *XForwardForFilter) Name() string {
	return FilterXForward
}

// Pre execute before proxy
func (f *XForwardForFilter) Pre(c filter.Context) (statusCode int, err error) {
	c.ForwardRequest().Header.Add("X-Forwarded-For", c.OriginRequest().RemoteIP().String())
	return f.BaseFilter.Pre(c)
}
