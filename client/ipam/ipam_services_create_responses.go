package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamServicesCreateReader is a Reader for the IpamServicesCreate structure.
type IpamServicesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamServicesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamServicesCreateCreated creates a IpamServicesCreateCreated with default headers values
func NewIpamServicesCreateCreated() *IpamServicesCreateCreated {
	return &IpamServicesCreateCreated{}
}

/* IpamServicesCreateCreated describes a response with status code 201, with default header values.

IpamServicesCreateCreated ipam services create created
*/
type IpamServicesCreateCreated struct {
	Payload *models.Service
}

func (o *IpamServicesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/services/][%d] ipamServicesCreateCreated  %+v", 201, o.Payload)
}
func (o *IpamServicesCreateCreated) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
