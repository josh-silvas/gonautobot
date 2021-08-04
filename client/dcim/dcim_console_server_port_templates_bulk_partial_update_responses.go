package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsoleServerPortTemplatesBulkPartialUpdateReader is a Reader for the DcimConsoleServerPortTemplatesBulkPartialUpdate structure.
type DcimConsoleServerPortTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesBulkPartialUpdateOK creates a DcimConsoleServerPortTemplatesBulkPartialUpdateOK with default headers values
func NewDcimConsoleServerPortTemplatesBulkPartialUpdateOK() *DcimConsoleServerPortTemplatesBulkPartialUpdateOK {
	return &DcimConsoleServerPortTemplatesBulkPartialUpdateOK{}
}

/* DcimConsoleServerPortTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortTemplatesBulkPartialUpdateOK dcim console server port templates bulk partial update o k
*/
type DcimConsoleServerPortTemplatesBulkPartialUpdateOK struct {
	Payload *models.ConsoleServerPortTemplate
}

func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-port-templates/][%d] dcimConsoleServerPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
