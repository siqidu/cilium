// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/health/models"
)

// GetHealthzReader is a Reader for the GetHealthz structure.
type GetHealthzReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHealthzReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHealthzOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetHealthzFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHealthzOK creates a GetHealthzOK with default headers values
func NewGetHealthzOK() *GetHealthzOK {
	return &GetHealthzOK{}
}

/*GetHealthzOK handles this case with default header values.

Success
*/
type GetHealthzOK struct {
	Payload *models.HealthResponse
}

func (o *GetHealthzOK) Error() string {
	return fmt.Sprintf("[GET /healthz][%d] getHealthzOK  %+v", 200, o.Payload)
}

func (o *GetHealthzOK) GetPayload() *models.HealthResponse {
	return o.Payload
}

func (o *GetHealthzOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HealthResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHealthzFailed creates a GetHealthzFailed with default headers values
func NewGetHealthzFailed() *GetHealthzFailed {
	return &GetHealthzFailed{}
}

/*GetHealthzFailed handles this case with default header values.

Failed to contact local Cilium daemon
*/
type GetHealthzFailed struct {
	Payload models.Error
}

func (o *GetHealthzFailed) Error() string {
	return fmt.Sprintf("[GET /healthz][%d] getHealthzFailed  %+v", 500, o.Payload)
}

func (o *GetHealthzFailed) GetPayload() models.Error {
	return o.Payload
}

func (o *GetHealthzFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
