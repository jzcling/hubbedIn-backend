package database

import (
	"fmt"

	"github.com/pkg/errors"
)

func nilErr(s string) error {
	return errors.New(fmt.Sprintf("Input parameter %s is nil", s))
}

func failedToInsertErr(err error, s string, m interface{}) error {
	msg := "Failed to insert %s %v"
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf(msg, s, m))
	}
	return errors.New(fmt.Sprintf(msg, s, m))
}

func updateErr(err error, s string, id uint64) error {
	msg := "Cannot update %s with id %v"
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf(msg, s, id))
	}
	return errors.New(fmt.Sprintf(msg, s, id))
}

func deleteErr(err error, s string, id uint64) error {
	msg := "Cannot delete %s with id %v"
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf(msg, s, id))
	}
	return errors.New(fmt.Sprintf(msg, s, id))
}

func candidateIDErr(err error, cid uint64) error {
	msg := "Error getting candidate joblistings for candidate id %v"
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf(msg, cid))
	}
	return errors.New(fmt.Sprintf(msg, cid))
}
