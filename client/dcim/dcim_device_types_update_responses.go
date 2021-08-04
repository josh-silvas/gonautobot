package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceTypesUpdateReader is a Reader for the DcimDeviceTypesUpdate structure.
type DcimDeviceTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceTypesUpdateOK creates a DcimDeviceTypesUpdateOK with default headers values
func NewDcimDeviceTypesUpdateOK() *DcimDeviceTypesUpdateOK {
	return &DcimDeviceTypesUpdateOK{}
}

/* DcimDeviceTypesUpdateOK describes a response with status code 200, with default header values.

DcimDeviceTypesUpdateOK dcim device types update o k
*/
type DcimDeviceTypesUpdateOK struct {
	Payload *models.DeviceType
}

func (o *DcimDeviceTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-types/{id}/][%d] dcimDeviceTypesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceTypesUpdateOK) GetPayload() *models.DeviceType {
	return o.Payload
}

func (o *DcimDeviceTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
