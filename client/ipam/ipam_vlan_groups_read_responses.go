package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVlanGroupsReadReader is a Reader for the IpamVlanGroupsRead structure.
type IpamVlanGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlanGroupsReadOK creates a IpamVlanGroupsReadOK with default headers values
func NewIpamVlanGroupsReadOK() *IpamVlanGroupsReadOK {
	return &IpamVlanGroupsReadOK{}
}

/* IpamVlanGroupsReadOK describes a response with status code 200, with default header values.

IpamVlanGroupsReadOK ipam vlan groups read o k
*/
type IpamVlanGroupsReadOK struct {
	Payload *models.VLANGroup
}

func (o *IpamVlanGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsReadOK  %+v", 200, o.Payload)
}
func (o *IpamVlanGroupsReadOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
