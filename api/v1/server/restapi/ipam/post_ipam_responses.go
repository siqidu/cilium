// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// PostIpamCreatedCode is the HTTP code returned for type PostIpamCreated
const PostIpamCreatedCode int = 201

/*PostIpamCreated Success

swagger:response postIpamCreated
*/
type PostIpamCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IPAMResponse `json:"body,omitempty"`
}

// NewPostIpamCreated creates PostIpamCreated with default headers values
func NewPostIpamCreated() *PostIpamCreated {

	return &PostIpamCreated{}
}

// WithPayload adds the payload to the post ipam created response
func (o *PostIpamCreated) WithPayload(payload *models.IPAMResponse) *PostIpamCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post ipam created response
func (o *PostIpamCreated) SetPayload(payload *models.IPAMResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIpamCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostIpamFailureCode is the HTTP code returned for type PostIpamFailure
const PostIpamFailureCode int = 502

/*PostIpamFailure Allocation failure

swagger:response postIpamFailure
*/
type PostIpamFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPostIpamFailure creates PostIpamFailure with default headers values
func NewPostIpamFailure() *PostIpamFailure {

	return &PostIpamFailure{}
}

// WithPayload adds the payload to the post ipam failure response
func (o *PostIpamFailure) WithPayload(payload models.Error) *PostIpamFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post ipam failure response
func (o *PostIpamFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIpamFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
