// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageServiceForceRemoveCompletionPostReader is a Reader for the StorageServiceForceRemoveCompletionPost structure.
type StorageServiceForceRemoveCompletionPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceForceRemoveCompletionPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceForceRemoveCompletionPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceForceRemoveCompletionPostOK creates a StorageServiceForceRemoveCompletionPostOK with default headers values
func NewStorageServiceForceRemoveCompletionPostOK() *StorageServiceForceRemoveCompletionPostOK {
	return &StorageServiceForceRemoveCompletionPostOK{}
}

/*StorageServiceForceRemoveCompletionPostOK handles this case with default header values.

StorageServiceForceRemoveCompletionPostOK storage service force remove completion post o k
*/
type StorageServiceForceRemoveCompletionPostOK struct {
}

func (o *StorageServiceForceRemoveCompletionPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/force_remove_completion][%d] storageServiceForceRemoveCompletionPostOK ", 200)
}

func (o *StorageServiceForceRemoveCompletionPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}