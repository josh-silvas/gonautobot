package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamRirsCreateReader is a Reader for the IpamRirsCreate structure.
type IpamRirsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamRirsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRirsCreateCreated creates a IpamRirsCreateCreated with default headers values
func NewIpamRirsCreateCreated() *IpamRirsCreateCreated {
	return &IpamRirsCreateCreated{}
}

/* IpamRirsCreateCreated describes a response with status code 201, with default header values.

IpamRirsCreateCreated ipam rirs create created
*/
type IpamRirsCreateCreated struct {
	Payload *models.RIR
}

func (o *IpamRirsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/rirs/][%d] ipamRirsCreateCreated  %+v", 201, o.Payload)
}
func (o *IpamRirsCreateCreated) GetPayload() *models.RIR {
	return o.Payload
}

func (o *IpamRirsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RIR)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
