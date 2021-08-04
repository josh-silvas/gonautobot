package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimFrontPortsPathsReader is a Reader for the DcimFrontPortsPaths structure.
type DcimFrontPortsPathsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsPathsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsPathsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortsPathsOK creates a DcimFrontPortsPathsOK with default headers values
func NewDcimFrontPortsPathsOK() *DcimFrontPortsPathsOK {
	return &DcimFrontPortsPathsOK{}
}

/* DcimFrontPortsPathsOK describes a response with status code 200, with default header values.

DcimFrontPortsPathsOK dcim front ports paths o k
*/
type DcimFrontPortsPathsOK struct {
	Payload *models.FrontPort
}

func (o *DcimFrontPortsPathsOK) Error() string {
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/paths/][%d] dcimFrontPortsPathsOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortsPathsOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsPathsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
