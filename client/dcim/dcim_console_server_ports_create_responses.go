package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsoleServerPortsCreateReader is a Reader for the DcimConsoleServerPortsCreate structure.
type DcimConsoleServerPortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimConsoleServerPortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortsCreateCreated creates a DcimConsoleServerPortsCreateCreated with default headers values
func NewDcimConsoleServerPortsCreateCreated() *DcimConsoleServerPortsCreateCreated {
	return &DcimConsoleServerPortsCreateCreated{}
}

/* DcimConsoleServerPortsCreateCreated describes a response with status code 201, with default header values.

DcimConsoleServerPortsCreateCreated dcim console server ports create created
*/
type DcimConsoleServerPortsCreateCreated struct {
	Payload *models.ConsoleServerPort
}

func (o *DcimConsoleServerPortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/console-server-ports/][%d] dcimConsoleServerPortsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimConsoleServerPortsCreateCreated) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
