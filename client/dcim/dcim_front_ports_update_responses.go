package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimFrontPortsUpdateReader is a Reader for the DcimFrontPortsUpdate structure.
type DcimFrontPortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortsUpdateOK creates a DcimFrontPortsUpdateOK with default headers values
func NewDcimFrontPortsUpdateOK() *DcimFrontPortsUpdateOK {
	return &DcimFrontPortsUpdateOK{}
}

/* DcimFrontPortsUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortsUpdateOK dcim front ports update o k
*/
type DcimFrontPortsUpdateOK struct {
	Payload *models.FrontPort
}

func (o *DcimFrontPortsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-ports/{id}/][%d] dcimFrontPortsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortsUpdateOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
