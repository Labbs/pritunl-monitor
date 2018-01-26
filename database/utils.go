package database

import (
	"log"

	"gopkg.in/mgo.v2"
)

// Get mongodb error code from error
func GetErrorCode(err error) (errCode int) {
	switch err := err.(type) {
	case *mgo.LastError:
		errCode = err.Code
	case *mgo.QueryError:
		errCode = err.Code
	}

	return
}

// Parse database error data and return error type
func ParseError(err error) (newErr error) {
	if err == mgo.ErrNotFound {
		log.Println(err)
		return
	}

	errCode := GetErrorCode(err)

	switch errCode {
	case 11000, 11001, 12582, 16460:
		log.Println(err)
	default:
		log.Println(err)
	}

	return
}
