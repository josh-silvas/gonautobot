package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamVrfsReadReader is a Reader for the IpamVrfsRead structure.
type IpamVrfsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVrfsReadOK creates a IpamVrfsReadOK with default headers values
func NewIpamVrfsReadOK() *IpamVrfsReadOK {
	return &IpamVrfsReadOK{}
}

/* IpamVrfsReadOK describes a response with status code 200, with default header values.

IpamVrfsReadOK ipam vrfs read o k
*/
type IpamVrfsReadOK struct {
	Payload *models.VRF
}

func (o *IpamVrfsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/vrfs/{id}/][%d] ipamVrfsReadOK  %+v", 200, o.Payload)
}
func (o *IpamVrfsReadOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
