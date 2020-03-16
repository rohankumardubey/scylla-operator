// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla_v2/models"
)

// FindConfigCommitlogSyncReader is a Reader for the FindConfigCommitlogSync structure.
type FindConfigCommitlogSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCommitlogSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCommitlogSyncOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCommitlogSyncDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCommitlogSyncOK creates a FindConfigCommitlogSyncOK with default headers values
func NewFindConfigCommitlogSyncOK() *FindConfigCommitlogSyncOK {
	return &FindConfigCommitlogSyncOK{}
}

/*FindConfigCommitlogSyncOK handles this case with default header values.

Config value
*/
type FindConfigCommitlogSyncOK struct {
	Payload string
}

func (o *FindConfigCommitlogSyncOK) Error() string {
	return fmt.Sprintf("[GET /config/commitlog_sync][%d] findConfigCommitlogSyncOK  %+v", 200, o.Payload)
}

func (o *FindConfigCommitlogSyncOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigCommitlogSyncOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCommitlogSyncDefault creates a FindConfigCommitlogSyncDefault with default headers values
func NewFindConfigCommitlogSyncDefault(code int) *FindConfigCommitlogSyncDefault {
	return &FindConfigCommitlogSyncDefault{
		_statusCode: code,
	}
}

/*FindConfigCommitlogSyncDefault handles this case with default header values.

unexpected error
*/
type FindConfigCommitlogSyncDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config commitlog sync default response
func (o *FindConfigCommitlogSyncDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCommitlogSyncDefault) Error() string {
	return fmt.Sprintf("[GET /config/commitlog_sync][%d] find_config_commitlog_sync default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigCommitlogSyncDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCommitlogSyncDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}