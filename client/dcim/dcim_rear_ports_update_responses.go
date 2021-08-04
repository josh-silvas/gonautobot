package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRearPortsUpdateReader is a Reader for the DcimRearPortsUpdate structure.
type DcimRearPortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortsUpdateOK creates a DcimRearPortsUpdateOK with default headers values
func NewDcimRearPortsUpdateOK() *DcimRearPortsUpdateOK {
	return &DcimRearPortsUpdateOK{}
}

/* DcimRearPortsUpdateOK describes a response with status code 200, with default header values.

DcimRearPortsUpdateOK dcim rear ports update o k
*/
type DcimRearPortsUpdateOK struct {
	Payload *models.RearPort
}

func (o *DcimRearPortsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rear-ports/{id}/][%d] dcimRearPortsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRearPortsUpdateOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
