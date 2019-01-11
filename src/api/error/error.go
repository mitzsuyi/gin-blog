package error

import (
 "net/http"
)

type Error struct {
  Code int
  Err  error
  Internal bool
}

func NewError(code int, err error, internal bool) * Error {
  if err != nil{  
   return &Error{Code:code, Err:err, Internal: internal} 
  } 
  return nil
}
func ApiError(code int, err error) *Error {
  return NewError(code, err, false)
}
func BadRequest(err error) *Error {
  return ApiError(http.StatusBadRequest, err)
}
func InternalError(err error) * Error {
 return NewError(http.StatusInternalServerError, err, true)
}