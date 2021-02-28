// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package apierr

import (
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

const (
	Errors_EmailNoFound  = "Community_EmailNoFound"
	Errors_EmailExists   = "Community_EmailExists"
	Errors_UserNoFound   = "Community_UserNoFound"
	Errors_NameExists    = "Community_NameExists"
	Errors_PageInvalid   = "Community_PageInvalid"
	Errors_EmptyIcon     = "Community_EmptyIcon"
	Errors_EmptyPassword = "Community_EmptyPassword"
	Errors_EmptyName     = "Community_EmptyName"
	Errors_EmptyEmail    = "Community_EmptyEmail"
	Errors_MysqlErr      = "Community_MysqlErr"
	Errors_RedisErr      = "Community_RedisErr"
	Errors_CodeErr       = "Community_CodeErr"
	Errors_FieldInvalid  = "Community_FieldInvalid"
	Errors_NilPointer    = "Community_NilPointer"
)

func IsEmailNoFound(err error) bool {
	return errors.Reason(err) == Errors_EmailNoFound
}

func IsEmailExists(err error) bool {
	return errors.Reason(err) == Errors_EmailExists
}

func IsUserNoFound(err error) bool {
	return errors.Reason(err) == Errors_UserNoFound
}

func IsNameExists(err error) bool {
	return errors.Reason(err) == Errors_NameExists
}

func IsPageInvalid(err error) bool {
	return errors.Reason(err) == Errors_PageInvalid
}

func IsEmptyIcon(err error) bool {
	return errors.Reason(err) == Errors_EmptyIcon
}

func IsEmptyPassword(err error) bool {
	return errors.Reason(err) == Errors_EmptyPassword
}

func IsEmptyName(err error) bool {
	return errors.Reason(err) == Errors_EmptyName
}

func IsEmptyEmail(err error) bool {
	return errors.Reason(err) == Errors_EmptyEmail
}

func IsMysqlErr(err error) bool {
	return errors.Reason(err) == Errors_MysqlErr
}

func IsRedisErr(err error) bool {
	return errors.Reason(err) == Errors_RedisErr
}

func IsCodeErr(err error) bool {
	return errors.Reason(err) == Errors_CodeErr
}

func IsFieldInvalid(err error) bool {
	return errors.Reason(err) == Errors_FieldInvalid
}

func IsNilPointer(err error) bool {
	return errors.Reason(err) == Errors_NilPointer
}