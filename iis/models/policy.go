package models

import "github.com/mehrdadnekopour/go-tools/mypes"

// Policy ...
type Policy struct {
	mypes.Model
	Name       string      `json:"name"`
	Statements *Statements `json:"statements"`
}

// Statement ...
type Statement struct {
	mypes.Model
}

// Policies ...
type Policies []*Policy

// Statements ...
type Statements []*Statement
