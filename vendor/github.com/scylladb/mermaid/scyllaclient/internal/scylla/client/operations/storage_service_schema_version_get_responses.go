// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageServiceSchemaVersionGetReader is a Reader for the StorageServiceSchemaVersionGet structure.
type StorageServiceSchemaVersionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceSchemaVersionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceSchemaVersionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceSchemaVersionGetOK creates a StorageServiceSchemaVersionGetOK with default headers values
func NewStorageServiceSchemaVersionGetOK() *StorageServiceSchemaVersionGetOK {
	return &StorageServiceSchemaVersionGetOK{}
}

/*StorageServiceSchemaVersionGetOK handles this case with default header values.

StorageServiceSchemaVersionGetOK storage service schema version get o k
*/
type StorageServiceSchemaVersionGetOK struct {
	Payload string
}

func (o *StorageServiceSchemaVersionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/schema_version][%d] storageServiceSchemaVersionGetOK  %+v", 200, o.Payload)
}

func (o *StorageServiceSchemaVersionGetOK) GetPayload() string {
	return o.Payload
}

func (o *StorageServiceSchemaVersionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}