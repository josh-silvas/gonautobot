package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRearPortsBulkUpdateReader is a Reader for the DcimRearPortsBulkUpdate structure.
type DcimRearPortsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortsBulkUpdateOK creates a DcimRearPortsBulkUpdateOK with default headers values
func NewDcimRearPortsBulkUpdateOK() *DcimRearPortsBulkUpdateOK {
	return &DcimRearPortsBulkUpdateOK{}
}

/* DcimRearPortsBulkUpdateOK describes a response with status code 200, with default header values.

DcimRearPortsBulkUpdateOK dcim rear ports bulk update o k
*/
type DcimRearPortsBulkUpdateOK struct {
	Payload *models.RearPort
}

func (o *DcimRearPortsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rear-ports/][%d] dcimRearPortsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRearPortsBulkUpdateOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
