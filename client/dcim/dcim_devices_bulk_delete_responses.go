package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDevicesBulkDeleteReader is a Reader for the DcimDevicesBulkDelete structure.
type DcimDevicesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDevicesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDevicesBulkDeleteNoContent creates a DcimDevicesBulkDeleteNoContent with default headers values
func NewDcimDevicesBulkDeleteNoContent() *DcimDevicesBulkDeleteNoContent {
	return &DcimDevicesBulkDeleteNoContent{}
}

/* DcimDevicesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimDevicesBulkDeleteNoContent dcim devices bulk delete no content
*/
type DcimDevicesBulkDeleteNoContent struct {
}

func (o *DcimDevicesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/devices/][%d] dcimDevicesBulkDeleteNoContent ", 204)
}

func (o *DcimDevicesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
