/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/types"
    "github.com/apimatic/go-core-runtime/utilities"
    "net/http"
)

func ToPointer[T any](value T) *T {
    return &value
}

// FileWrapper is a struct that represents a file along with its metadata such as the
// file content, file name, and file headers.
type FileWrapper = https.FileWrapper

// GetFile retrieves a file from the given fileUrl and returns it as a FileWrapper.
// It makes an HTTP GET request to the fileUrl to fetch the file's content and metadata.
var GetFile = https.GetFile 

// Optional is a generic struct that allows any type to be used as optional and nullable.
// Optional.set is true when Optional.value is to be used.
type Optional[T any] struct {
    types.Optional[T]
}

// NewOptional creates and returns an Optional instance with the given value set.
func NewOptional[T any](value *T) Optional[T] {
    return Optional[T]{
    	Optional: types.NewOptional(value),
    }
}

// Value returns the value stored in the Optional. It returns nil if no value is set.
func (o *Optional[T]) Value() *T {
    return o.Optional.Value()
}

// SetValue sets the value of the Optional.
func (o *Optional[T]) SetValue(value *T) {
    o.Optional.SetValue(value)
}

// IsValueSet returns true if a value is set in the Optional, false otherwise.
func (o *Optional[T]) IsValueSet() bool {
    return o.Optional.IsValueSet()
}

// ShouldSetValue sets whether the value should be used or not.
func (o *Optional[T]) ShouldSetValue(set bool) {
    o.Optional.ShouldSetValue(set)
}

// ApiResponse is a generic struct that represents an API response containing data and the HTTP response.
// The `Data` field holds the data of any type `T` returned by the API.
// The `Response` field contains the underlying HTTP response associated with the API call.
type ApiResponse[T any] struct {
    https.ApiResponse[T]
}

// NewApiResponse creates a new instance of ApiResponse.
// It takes the `data` of type `T` and the `response` as parameters, and returns an ApiResponse[T] struct.
func NewApiResponse[T any](
    data T,
    response *http.Response) ApiResponse[T] {
    return ApiResponse[T]{
    	ApiResponse: https.NewApiResponse(data, response),
    }
}

// DEFAULT_DATE is a utility.
var DEFAULT_DATE = utilities.DEFAULT_DATE 

// TimeToStringMap is a utility.
var TimeToStringMap = utilities.TimeToStringMap 

// TimeToStringSlice is a utility.
var TimeToStringSlice = utilities.TimeToStringSlice 

// ToTimeSlice is a utility.
var ToTimeSlice = utilities.ToTimeSlice 

// ToTimeMap is a utility.
var ToTimeMap = utilities.ToTimeMap 
