package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimVirtualChassisReadReader is a Reader for the DcimVirtualChassisRead structure.
type DcimVirtualChassisReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualChassisReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimVirtualChassisReadOK creates a DcimVirtualChassisReadOK with default headers values
func NewDcimVirtualChassisReadOK() *DcimVirtualChassisReadOK {
	return &DcimVirtualChassisReadOK{}
}

/* DcimVirtualChassisReadOK describes a response with status code 200, with default header values.

DcimVirtualChassisReadOK dcim virtual chassis read o k
*/
type DcimVirtualChassisReadOK struct {
	Payload *models.VirtualChassis
}

func (o *DcimVirtualChassisReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisReadOK  %+v", 200, o.Payload)
}
func (o *DcimVirtualChassisReadOK) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
