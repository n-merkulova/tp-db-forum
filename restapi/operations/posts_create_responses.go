// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/crueltycute/tech-db-forum/models"
)

// PostsCreateCreatedCode is the HTTP code returned for type PostsCreateCreated
const PostsCreateCreatedCode int = 201

/*PostsCreateCreated Посты успешно созданы.
Возвращает данные созданных постов в том же порядке, в котором их передали на вход метода.


swagger:response postsCreateCreated
*/
type PostsCreateCreated struct {

	/*
	  In: Body
	*/
	Payload models.Posts `json:"body,omitempty"`
}

// NewPostsCreateCreated creates PostsCreateCreated with default headers values
func NewPostsCreateCreated() *PostsCreateCreated {

	return &PostsCreateCreated{}
}

// WithPayload adds the payload to the posts create created response
func (o *PostsCreateCreated) WithPayload(payload models.Posts) *PostsCreateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the posts create created response
func (o *PostsCreateCreated) SetPayload(payload models.Posts) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostsCreateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Posts, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// PostsCreateNotFoundCode is the HTTP code returned for type PostsCreateNotFound
const PostsCreateNotFoundCode int = 404

/*PostsCreateNotFound Ветка обсуждения отсутствует в базе данных.


swagger:response postsCreateNotFound
*/
type PostsCreateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostsCreateNotFound creates PostsCreateNotFound with default headers values
func NewPostsCreateNotFound() *PostsCreateNotFound {

	return &PostsCreateNotFound{}
}

// WithPayload adds the payload to the posts create not found response
func (o *PostsCreateNotFound) WithPayload(payload *models.Error) *PostsCreateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the posts create not found response
func (o *PostsCreateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostsCreateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostsCreateConflictCode is the HTTP code returned for type PostsCreateConflict
const PostsCreateConflictCode int = 409

/*PostsCreateConflict Хотя бы один родительский пост отсутсвует в текущей ветке обсуждения.


swagger:response postsCreateConflict
*/
type PostsCreateConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostsCreateConflict creates PostsCreateConflict with default headers values
func NewPostsCreateConflict() *PostsCreateConflict {

	return &PostsCreateConflict{}
}

// WithPayload adds the payload to the posts create conflict response
func (o *PostsCreateConflict) WithPayload(payload *models.Error) *PostsCreateConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the posts create conflict response
func (o *PostsCreateConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostsCreateConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
