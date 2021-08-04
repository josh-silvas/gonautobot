package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamPrefixesAvailablePrefixesCreateReader is a Reader for the IpamPrefixesAvailablePrefixesCreate structure.
type IpamPrefixesAvailablePrefixesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailablePrefixesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamPrefixesAvailablePrefixesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamPrefixesAvailablePrefixesCreateCreated creates a IpamPrefixesAvailablePrefixesCreateCreated with default headers values
func NewIpamPrefixesAvailablePrefixesCreateCreated() *IpamPrefixesAvailablePrefixesCreateCreated {
	return &IpamPrefixesAvailablePrefixesCreateCreated{}
}

/* IpamPrefixesAvailablePrefixesCreateCreated describes a response with status code 201, with default header values.

IpamPrefixesAvailablePrefixesCreateCreated ipam prefixes available prefixes create created
*/
type IpamPrefixesAvailablePrefixesCreateCreated struct {
	Payload *models.Prefix
}

func (o *IpamPrefixesAvailablePrefixesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-prefixes/][%d] ipamPrefixesAvailablePrefixesCreateCreated  %+v", 201, o.Payload)
}
func (o *IpamPrefixesAvailablePrefixesCreateCreated) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesAvailablePrefixesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
