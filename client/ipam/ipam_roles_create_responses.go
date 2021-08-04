package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamRolesCreateReader is a Reader for the IpamRolesCreate structure.
type IpamRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRolesCreateCreated creates a IpamRolesCreateCreated with default headers values
func NewIpamRolesCreateCreated() *IpamRolesCreateCreated {
	return &IpamRolesCreateCreated{}
}

/* IpamRolesCreateCreated describes a response with status code 201, with default header values.

IpamRolesCreateCreated ipam roles create created
*/
type IpamRolesCreateCreated struct {
	Payload *models.Role
}

func (o *IpamRolesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/roles/][%d] ipamRolesCreateCreated  %+v", 201, o.Payload)
}
func (o *IpamRolesCreateCreated) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
