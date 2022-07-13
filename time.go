package time

import (
	"time"
)

type DataSource interface {
	Now() time.Time
}

func NewDefaultDataSource() DefaultDataSource {
	return DefaultDataSource{}
}

type DefaultDataSource struct{}

func (d DefaultDataSource) Now() time.Time {
	return time.Now().UTC()
}
