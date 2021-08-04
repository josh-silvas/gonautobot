package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsoleServerPortTemplatesReadReader is a Reader for the DcimConsoleServerPortTemplatesRead structure.
type DcimConsoleServerPortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesReadOK creates a DcimConsoleServerPortTemplatesReadOK with default headers values
func NewDcimConsoleServerPortTemplatesReadOK() *DcimConsoleServerPortTemplatesReadOK {
	return &DcimConsoleServerPortTemplatesReadOK{}
}

/* DcimConsoleServerPortTemplatesReadOK describes a response with status code 200, with default header values.

DcimConsoleServerPortTemplatesReadOK dcim console server port templates read o k
*/
type DcimConsoleServerPortTemplatesReadOK struct {
	Payload *models.ConsoleServerPortTemplate
}

func (o *DcimConsoleServerPortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesReadOK  %+v", 200, o.Payload)
}
func (o *DcimConsoleServerPortTemplatesReadOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
