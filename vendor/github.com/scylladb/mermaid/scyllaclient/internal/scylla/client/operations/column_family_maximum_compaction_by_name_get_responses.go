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

// ColumnFamilyMaximumCompactionByNameGetReader is a Reader for the ColumnFamilyMaximumCompactionByNameGet structure.
type ColumnFamilyMaximumCompactionByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMaximumCompactionByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMaximumCompactionByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMaximumCompactionByNameGetOK creates a ColumnFamilyMaximumCompactionByNameGetOK with default headers values
func NewColumnFamilyMaximumCompactionByNameGetOK() *ColumnFamilyMaximumCompactionByNameGetOK {
	return &ColumnFamilyMaximumCompactionByNameGetOK{}
}

/*ColumnFamilyMaximumCompactionByNameGetOK handles this case with default header values.

ColumnFamilyMaximumCompactionByNameGetOK column family maximum compaction by name get o k
*/
type ColumnFamilyMaximumCompactionByNameGetOK struct {
	Payload string
}

func (o *ColumnFamilyMaximumCompactionByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/maximum_compaction/{name}][%d] columnFamilyMaximumCompactionByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMaximumCompactionByNameGetOK) GetPayload() string {
	return o.Payload
}

func (o *ColumnFamilyMaximumCompactionByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}