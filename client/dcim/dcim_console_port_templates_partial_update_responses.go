package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsolePortTemplatesPartialUpdateReader is a Reader for the DcimConsolePortTemplatesPartialUpdate structure.
type DcimConsolePortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortTemplatesPartialUpdateOK creates a DcimConsolePortTemplatesPartialUpdateOK with default headers values
func NewDcimConsolePortTemplatesPartialUpdateOK() *DcimConsolePortTemplatesPartialUpdateOK {
	return &DcimConsolePortTemplatesPartialUpdateOK{}
}

/* DcimConsolePortTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortTemplatesPartialUpdateOK dcim console port templates partial update o k
*/
type DcimConsolePortTemplatesPartialUpdateOK struct {
	Payload *models.ConsolePortTemplate
}

func (o *DcimConsolePortTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsolePortTemplatesPartialUpdateOK) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
