package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamRirsBulkUpdateReader is a Reader for the IpamRirsBulkUpdate structure.
type IpamRirsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRirsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRirsBulkUpdateOK creates a IpamRirsBulkUpdateOK with default headers values
func NewIpamRirsBulkUpdateOK() *IpamRirsBulkUpdateOK {
	return &IpamRirsBulkUpdateOK{}
}

/* IpamRirsBulkUpdateOK describes a response with status code 200, with default header values.

IpamRirsBulkUpdateOK ipam rirs bulk update o k
*/
type IpamRirsBulkUpdateOK struct {
	Payload *models.RIR
}

func (o *IpamRirsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/rirs/][%d] ipamRirsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamRirsBulkUpdateOK) GetPayload() *models.RIR {
	return o.Payload
}

func (o *IpamRirsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RIR)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
