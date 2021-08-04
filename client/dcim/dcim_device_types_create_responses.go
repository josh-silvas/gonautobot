package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceTypesCreateReader is a Reader for the DcimDeviceTypesCreate structure.
type DcimDeviceTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimDeviceTypesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceTypesCreateCreated creates a DcimDeviceTypesCreateCreated with default headers values
func NewDcimDeviceTypesCreateCreated() *DcimDeviceTypesCreateCreated {
	return &DcimDeviceTypesCreateCreated{}
}

/* DcimDeviceTypesCreateCreated describes a response with status code 201, with default header values.

DcimDeviceTypesCreateCreated dcim device types create created
*/
type DcimDeviceTypesCreateCreated struct {
	Payload *models.DeviceType
}

func (o *DcimDeviceTypesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/device-types/][%d] dcimDeviceTypesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimDeviceTypesCreateCreated) GetPayload() *models.DeviceType {
	return o.Payload
}

func (o *DcimDeviceTypesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
