package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamAggregatesBulkPartialUpdateReader is a Reader for the IpamAggregatesBulkPartialUpdate structure.
type IpamAggregatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAggregatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamAggregatesBulkPartialUpdateOK creates a IpamAggregatesBulkPartialUpdateOK with default headers values
func NewIpamAggregatesBulkPartialUpdateOK() *IpamAggregatesBulkPartialUpdateOK {
	return &IpamAggregatesBulkPartialUpdateOK{}
}

/* IpamAggregatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamAggregatesBulkPartialUpdateOK ipam aggregates bulk partial update o k
*/
type IpamAggregatesBulkPartialUpdateOK struct {
	Payload *models.Aggregate
}

func (o *IpamAggregatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/aggregates/][%d] ipamAggregatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamAggregatesBulkPartialUpdateOK) GetPayload() *models.Aggregate {
	return o.Payload
}

func (o *IpamAggregatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
