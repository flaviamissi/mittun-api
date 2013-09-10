package speakers

import (
    "launchpad.net/gocheck"
    "testing"
)

var _ = gocheck.Suite(&S{})

type S struct {}

func Test(t *testing.T) {
	gocheck.TestingT(t)
}

func (s *S) Test(c *gocheck.C) {
}
