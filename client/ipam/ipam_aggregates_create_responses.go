package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamAggregatesCreateReader is a Reader for the IpamAggregatesCreate structure.
type IpamAggregatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamAggregatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamAggregatesCreateCreated creates a IpamAggregatesCreateCreated with default headers values
func NewIpamAggregatesCreateCreated() *IpamAggregatesCreateCreated {
	return &IpamAggregatesCreateCreated{}
}

/* IpamAggregatesCreateCreated describes a response with status code 201, with default header values.

IpamAggregatesCreateCreated ipam aggregates create created
*/
type IpamAggregatesCreateCreated struct {
	Payload *models.Aggregate
}

func (o *IpamAggregatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/aggregates/][%d] ipamAggregatesCreateCreated  %+v", 201, o.Payload)
}
func (o *IpamAggregatesCreateCreated) GetPayload() *models.Aggregate {
	return o.Payload
}

func (o *IpamAggregatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
