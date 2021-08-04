package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsolePortsBulkUpdateReader is a Reader for the DcimConsolePortsBulkUpdate structure.
type DcimConsolePortsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortsBulkUpdateOK creates a DcimConsolePortsBulkUpdateOK with default headers values
func NewDcimConsolePortsBulkUpdateOK() *DcimConsolePortsBulkUpdateOK {
	return &DcimConsolePortsBulkUpdateOK{}
}

/* DcimConsolePortsBulkUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortsBulkUpdateOK dcim console ports bulk update o k
*/
type DcimConsolePortsBulkUpdateOK struct {
	Payload *models.ConsolePort
}

func (o *DcimConsolePortsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-ports/][%d] dcimConsolePortsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsolePortsBulkUpdateOK) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
