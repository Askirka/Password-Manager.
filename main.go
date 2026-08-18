package main

import "time"

func main() {

}

type Password struct {
	Name         string
	Value        string
	Category     string
	CreatedAt    time.Time
	LastModified time.Time
}

func NewPassword(name, value, category string) Password
