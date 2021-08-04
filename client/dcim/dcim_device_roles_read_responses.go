package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceRolesReadReader is a Reader for the DcimDeviceRolesRead structure.
type DcimDeviceRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceRolesReadOK creates a DcimDeviceRolesReadOK with default headers values
func NewDcimDeviceRolesReadOK() *DcimDeviceRolesReadOK {
	return &DcimDeviceRolesReadOK{}
}

/* DcimDeviceRolesReadOK describes a response with status code 200, with default header values.

DcimDeviceRolesReadOK dcim device roles read o k
*/
type DcimDeviceRolesReadOK struct {
	Payload *models.DeviceRole
}

func (o *DcimDeviceRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/device-roles/{id}/][%d] dcimDeviceRolesReadOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceRolesReadOK) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
