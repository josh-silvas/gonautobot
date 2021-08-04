package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceRolesPartialUpdateReader is a Reader for the DcimDeviceRolesPartialUpdate structure.
type DcimDeviceRolesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceRolesPartialUpdateOK creates a DcimDeviceRolesPartialUpdateOK with default headers values
func NewDcimDeviceRolesPartialUpdateOK() *DcimDeviceRolesPartialUpdateOK {
	return &DcimDeviceRolesPartialUpdateOK{}
}

/* DcimDeviceRolesPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceRolesPartialUpdateOK dcim device roles partial update o k
*/
type DcimDeviceRolesPartialUpdateOK struct {
	Payload *models.DeviceRole
}

func (o *DcimDeviceRolesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-roles/{id}/][%d] dcimDeviceRolesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceRolesPartialUpdateOK) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
