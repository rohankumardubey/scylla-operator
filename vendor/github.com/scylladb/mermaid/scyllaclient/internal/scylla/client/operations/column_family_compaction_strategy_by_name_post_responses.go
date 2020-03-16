// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ColumnFamilyCompactionStrategyByNamePostReader is a Reader for the ColumnFamilyCompactionStrategyByNamePost structure.
type ColumnFamilyCompactionStrategyByNamePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyCompactionStrategyByNamePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyCompactionStrategyByNamePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyCompactionStrategyByNamePostOK creates a ColumnFamilyCompactionStrategyByNamePostOK with default headers values
func NewColumnFamilyCompactionStrategyByNamePostOK() *ColumnFamilyCompactionStrategyByNamePostOK {
	return &ColumnFamilyCompactionStrategyByNamePostOK{}
}

/*ColumnFamilyCompactionStrategyByNamePostOK handles this case with default header values.

ColumnFamilyCompactionStrategyByNamePostOK column family compaction strategy by name post o k
*/
type ColumnFamilyCompactionStrategyByNamePostOK struct {
}

func (o *ColumnFamilyCompactionStrategyByNamePostOK) Error() string {
	return fmt.Sprintf("[POST /column_family/compaction_strategy/{name}][%d] columnFamilyCompactionStrategyByNamePostOK ", 200)
}

func (o *ColumnFamilyCompactionStrategyByNamePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}