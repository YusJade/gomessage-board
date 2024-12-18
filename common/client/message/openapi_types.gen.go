// Package message provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package message

// Error defines model for Error.
type Error struct {
	Message *string `json:"message,omitempty"`
}

// Message defines model for Message.
type Message struct {
	Content  *string `json:"content,omitempty"`
	Datetime *string `json:"datetime,omitempty"`
	Id       *string `json:"id,omitempty"`
}

// PostMessageBoardJSONRequestBody defines body for PostMessageBoard for application/json ContentType.
type PostMessageBoardJSONRequestBody = Message
