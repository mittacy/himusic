// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package apierr

import (
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

const (
	Errors_UserNoFound  = "Community_UserNoFound"
	Errors_MysqlErr     = "Community_MysqlErr"
	Errors_FieldInvalid = "Community_FieldInvalid"
)

func IsUserNoFound(err error) bool {
	return errors.Reason(err) == Errors_UserNoFound
}

func IsMysqlErr(err error) bool {
	return errors.Reason(err) == Errors_MysqlErr
}

func IsFieldInvalid(err error) bool {
	return errors.Reason(err) == Errors_FieldInvalid
}
