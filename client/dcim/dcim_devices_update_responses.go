package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDevicesUpdateReader is a Reader for the DcimDevicesUpdate structure.
type DcimDevicesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDevicesUpdateOK creates a DcimDevicesUpdateOK with default headers values
func NewDcimDevicesUpdateOK() *DcimDevicesUpdateOK {
	return &DcimDevicesUpdateOK{}
}

/* DcimDevicesUpdateOK describes a response with status code 200, with default header values.

DcimDevicesUpdateOK dcim devices update o k
*/
type DcimDevicesUpdateOK struct {
	Payload *models.DeviceWithConfigContext
}

func (o *DcimDevicesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/devices/{id}/][%d] dcimDevicesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDevicesUpdateOK) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
