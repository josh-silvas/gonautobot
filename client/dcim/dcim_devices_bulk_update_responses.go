package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDevicesBulkUpdateReader is a Reader for the DcimDevicesBulkUpdate structure.
type DcimDevicesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDevicesBulkUpdateOK creates a DcimDevicesBulkUpdateOK with default headers values
func NewDcimDevicesBulkUpdateOK() *DcimDevicesBulkUpdateOK {
	return &DcimDevicesBulkUpdateOK{}
}

/* DcimDevicesBulkUpdateOK describes a response with status code 200, with default header values.

DcimDevicesBulkUpdateOK dcim devices bulk update o k
*/
type DcimDevicesBulkUpdateOK struct {
	Payload *models.DeviceWithConfigContext
}

func (o *DcimDevicesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/devices/][%d] dcimDevicesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDevicesBulkUpdateOK) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
