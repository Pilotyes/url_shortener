/*
 * Swagger URL Shortener
 *
 * URL Shortener API
 *
 * API version: 1.0.0
 * Contact: admin@urlshortener.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package model

// ErrorMessage is a ...
type ErrorMessage struct {
	Code int32 `json:"code,omitempty"`

	Error_ string `json:"error,omitempty"`
}
