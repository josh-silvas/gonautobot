package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVlansCreateReader is a Reader for the IpamVlansCreate structure.
type IpamVlansCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamVlansCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlansCreateCreated creates a IpamVlansCreateCreated with default headers values
func NewIpamVlansCreateCreated() *IpamVlansCreateCreated {
	return &IpamVlansCreateCreated{}
}

/* IpamVlansCreateCreated describes a response with status code 201, with default header values.

IpamVlansCreateCreated ipam vlans create created
*/
type IpamVlansCreateCreated struct {
	Payload *models.VLAN
}

func (o *IpamVlansCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/vlans/][%d] ipamVlansCreateCreated  %+v", 201, o.Payload)
}
func (o *IpamVlansCreateCreated) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
