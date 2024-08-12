package main

type User struct {
	id        uint
	name      string
	surname   string
	nickname  string
	phone     string
	email     string
	instagram string
	telegram  string
	status_id string
	blocked   bool
}

func (u *User) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}
