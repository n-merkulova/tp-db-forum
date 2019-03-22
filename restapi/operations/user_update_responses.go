// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/crueltycute/tech-db-forum/models"
)

// UserUpdateOKCode is the HTTP code returned for type UserUpdateOK
const UserUpdateOKCode int = 200

/*UserUpdateOK Актуальная информация о пользователе после изменения профиля.


swagger:response userUpdateOK
*/
type UserUpdateOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewUserUpdateOK creates UserUpdateOK with default headers values
func NewUserUpdateOK() *UserUpdateOK {

	return &UserUpdateOK{}
}

// WithPayload adds the payload to the user update o k response
func (o *UserUpdateOK) WithPayload(payload *models.User) *UserUpdateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user update o k response
func (o *UserUpdateOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserUpdateNotFoundCode is the HTTP code returned for type UserUpdateNotFound
const UserUpdateNotFoundCode int = 404

/*UserUpdateNotFound Пользователь отсутсвует в системе.


swagger:response userUpdateNotFound
*/
type UserUpdateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUserUpdateNotFound creates UserUpdateNotFound with default headers values
func NewUserUpdateNotFound() *UserUpdateNotFound {

	return &UserUpdateNotFound{}
}

// WithPayload adds the payload to the user update not found response
func (o *UserUpdateNotFound) WithPayload(payload *models.Error) *UserUpdateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user update not found response
func (o *UserUpdateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserUpdateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserUpdateConflictCode is the HTTP code returned for type UserUpdateConflict
const UserUpdateConflictCode int = 409

/*UserUpdateConflict Новые данные профиля пользователя конфликтуют с имеющимися пользователями.


swagger:response userUpdateConflict
*/
type UserUpdateConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUserUpdateConflict creates UserUpdateConflict with default headers values
func NewUserUpdateConflict() *UserUpdateConflict {

	return &UserUpdateConflict{}
}

// WithPayload adds the payload to the user update conflict response
func (o *UserUpdateConflict) WithPayload(payload *models.Error) *UserUpdateConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user update conflict response
func (o *UserUpdateConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserUpdateConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
