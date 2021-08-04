package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRacksBulkUpdateReader is a Reader for the DcimRacksBulkUpdate structure.
type DcimRacksBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRacksBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRacksBulkUpdateOK creates a DcimRacksBulkUpdateOK with default headers values
func NewDcimRacksBulkUpdateOK() *DcimRacksBulkUpdateOK {
	return &DcimRacksBulkUpdateOK{}
}

/* DcimRacksBulkUpdateOK describes a response with status code 200, with default header values.

DcimRacksBulkUpdateOK dcim racks bulk update o k
*/
type DcimRacksBulkUpdateOK struct {
	Payload *models.Rack
}

func (o *DcimRacksBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/racks/][%d] dcimRacksBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRacksBulkUpdateOK) GetPayload() *models.Rack {
	return o.Payload
}

func (o *DcimRacksBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
