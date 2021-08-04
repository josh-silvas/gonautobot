package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVlanGroupsBulkUpdateReader is a Reader for the IpamVlanGroupsBulkUpdate structure.
type IpamVlanGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlanGroupsBulkUpdateOK creates a IpamVlanGroupsBulkUpdateOK with default headers values
func NewIpamVlanGroupsBulkUpdateOK() *IpamVlanGroupsBulkUpdateOK {
	return &IpamVlanGroupsBulkUpdateOK{}
}

/* IpamVlanGroupsBulkUpdateOK describes a response with status code 200, with default header values.

IpamVlanGroupsBulkUpdateOK ipam vlan groups bulk update o k
*/
type IpamVlanGroupsBulkUpdateOK struct {
	Payload *models.VLANGroup
}

func (o *IpamVlanGroupsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/][%d] ipamVlanGroupsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVlanGroupsBulkUpdateOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
