package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVlanGroupsUpdateReader is a Reader for the IpamVlanGroupsUpdate structure.
type IpamVlanGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlanGroupsUpdateOK creates a IpamVlanGroupsUpdateOK with default headers values
func NewIpamVlanGroupsUpdateOK() *IpamVlanGroupsUpdateOK {
	return &IpamVlanGroupsUpdateOK{}
}

/* IpamVlanGroupsUpdateOK describes a response with status code 200, with default header values.

IpamVlanGroupsUpdateOK ipam vlan groups update o k
*/
type IpamVlanGroupsUpdateOK struct {
	Payload *models.VLANGroup
}

func (o *IpamVlanGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVlanGroupsUpdateOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
