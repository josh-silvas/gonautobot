package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamPrefixesBulkPartialUpdateReader is a Reader for the IpamPrefixesBulkPartialUpdate structure.
type IpamPrefixesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamPrefixesBulkPartialUpdateOK creates a IpamPrefixesBulkPartialUpdateOK with default headers values
func NewIpamPrefixesBulkPartialUpdateOK() *IpamPrefixesBulkPartialUpdateOK {
	return &IpamPrefixesBulkPartialUpdateOK{}
}

/* IpamPrefixesBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamPrefixesBulkPartialUpdateOK ipam prefixes bulk partial update o k
*/
type IpamPrefixesBulkPartialUpdateOK struct {
	Payload *models.Prefix
}

func (o *IpamPrefixesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/prefixes/][%d] ipamPrefixesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamPrefixesBulkPartialUpdateOK) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
