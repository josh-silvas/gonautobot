package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRackRolesPartialUpdateReader is a Reader for the DcimRackRolesPartialUpdate structure.
type DcimRackRolesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackRolesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackRolesPartialUpdateOK creates a DcimRackRolesPartialUpdateOK with default headers values
func NewDcimRackRolesPartialUpdateOK() *DcimRackRolesPartialUpdateOK {
	return &DcimRackRolesPartialUpdateOK{}
}

/* DcimRackRolesPartialUpdateOK describes a response with status code 200, with default header values.

DcimRackRolesPartialUpdateOK dcim rack roles partial update o k
*/
type DcimRackRolesPartialUpdateOK struct {
	Payload *models.RackRole
}

func (o *DcimRackRolesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rack-roles/{id}/][%d] dcimRackRolesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRackRolesPartialUpdateOK) GetPayload() *models.RackRole {
	return o.Payload
}

func (o *DcimRackRolesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
