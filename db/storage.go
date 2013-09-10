// Package db implements storage functions to mittun
// #TODO
// - reconnection
// - connection pool
package db

import (
    "labix.org/v2/mgo"
)

type Storage struct {
    session *mgo.Session
    dbname  string
}

// Conn connects with mongodb in the given address
// and returns instance of Storage.
func Conn(addr, dbname string) (*Storage, error) {
    sess, err := mgo.DialWithTimeout(addr, 1e9)
    return &Storage{session: sess, dbname: dbname}, err
}

func (s *Storage) Close() {
    s.session.Close()
}
