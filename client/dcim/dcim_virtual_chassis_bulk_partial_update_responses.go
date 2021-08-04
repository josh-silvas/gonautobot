package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimVirtualChassisBulkPartialUpdateReader is a Reader for the DcimVirtualChassisBulkPartialUpdate structure.
type DcimVirtualChassisBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualChassisBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimVirtualChassisBulkPartialUpdateOK creates a DcimVirtualChassisBulkPartialUpdateOK with default headers values
func NewDcimVirtualChassisBulkPartialUpdateOK() *DcimVirtualChassisBulkPartialUpdateOK {
	return &DcimVirtualChassisBulkPartialUpdateOK{}
}

/* DcimVirtualChassisBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimVirtualChassisBulkPartialUpdateOK dcim virtual chassis bulk partial update o k
*/
type DcimVirtualChassisBulkPartialUpdateOK struct {
	Payload *models.VirtualChassis
}

func (o *DcimVirtualChassisBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-chassis/][%d] dcimVirtualChassisBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimVirtualChassisBulkPartialUpdateOK) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
