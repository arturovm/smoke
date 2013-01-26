package main

import (
	r "github.com/christopherhesse/rethinkgo"
)

func createUser(r *r.Session, name string, email string, password string) (*RESTError) {
	r.Close()
	return nil
}
