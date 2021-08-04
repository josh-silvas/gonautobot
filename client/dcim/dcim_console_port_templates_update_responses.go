package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsolePortTemplatesUpdateReader is a Reader for the DcimConsolePortTemplatesUpdate structure.
type DcimConsolePortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortTemplatesUpdateOK creates a DcimConsolePortTemplatesUpdateOK with default headers values
func NewDcimConsolePortTemplatesUpdateOK() *DcimConsolePortTemplatesUpdateOK {
	return &DcimConsolePortTemplatesUpdateOK{}
}

/* DcimConsolePortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortTemplatesUpdateOK dcim console port templates update o k
*/
type DcimConsolePortTemplatesUpdateOK struct {
	Payload *models.ConsolePortTemplate
}

func (o *DcimConsolePortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsolePortTemplatesUpdateOK) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
