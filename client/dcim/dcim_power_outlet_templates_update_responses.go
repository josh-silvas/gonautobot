package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerOutletTemplatesUpdateReader is a Reader for the DcimPowerOutletTemplatesUpdate structure.
type DcimPowerOutletTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletTemplatesUpdateOK creates a DcimPowerOutletTemplatesUpdateOK with default headers values
func NewDcimPowerOutletTemplatesUpdateOK() *DcimPowerOutletTemplatesUpdateOK {
	return &DcimPowerOutletTemplatesUpdateOK{}
}

/* DcimPowerOutletTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletTemplatesUpdateOK dcim power outlet templates update o k
*/
type DcimPowerOutletTemplatesUpdateOK struct {
	Payload *models.PowerOutletTemplate
}

func (o *DcimPowerOutletTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-outlet-templates/{id}/][%d] dcimPowerOutletTemplatesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerOutletTemplatesUpdateOK) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
