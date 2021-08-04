package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerPortTemplatesReadReader is a Reader for the DcimPowerPortTemplatesRead structure.
type DcimPowerPortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortTemplatesReadOK creates a DcimPowerPortTemplatesReadOK with default headers values
func NewDcimPowerPortTemplatesReadOK() *DcimPowerPortTemplatesReadOK {
	return &DcimPowerPortTemplatesReadOK{}
}

/* DcimPowerPortTemplatesReadOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesReadOK dcim power port templates read o k
*/
type DcimPowerPortTemplatesReadOK struct {
	Payload *models.PowerPortTemplate
}

func (o *DcimPowerPortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesReadOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPortTemplatesReadOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
