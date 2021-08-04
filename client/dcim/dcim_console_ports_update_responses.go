package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimConsolePortsUpdateReader is a Reader for the DcimConsolePortsUpdate structure.
type DcimConsolePortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortsUpdateOK creates a DcimConsolePortsUpdateOK with default headers values
func NewDcimConsolePortsUpdateOK() *DcimConsolePortsUpdateOK {
	return &DcimConsolePortsUpdateOK{}
}

/* DcimConsolePortsUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortsUpdateOK dcim console ports update o k
*/
type DcimConsolePortsUpdateOK struct {
	Payload *models.ConsolePort
}

func (o *DcimConsolePortsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-ports/{id}/][%d] dcimConsolePortsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsolePortsUpdateOK) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
