package db

import (
    "launchpad.net/gocheck"
    "testing"
)

var _ = gocheck.Suite(&S{})

type S struct {}

func Test(t *testing.T) { gocheck.TestingT(t) }

func (s *S) TestConnConnectsToDatabase(c *gocheck.C) {
    storage, err := Conn("127.0.0.1:27017", "mittun_storage_test")
    c.Assert(err, gocheck.IsNil)
    defer storage.session.Close()
    err = storage.session.Ping()
    c.Assert(err, gocheck.IsNil)
}

func (s *S) TestCloseClosesTheConnection(c *gocheck.C) {
    defer func() {
        r := recover()
        c.Assert(r, gocheck.NotNil)
    }()
    storage, err := Conn("127.0.0.1:27017", "mittun_storage_test")
    c.Assert(err, gocheck.IsNil)
    storage.Close()
    err = storage.session.Ping()
    c.Assert(err, gocheck.NotNil)
}
