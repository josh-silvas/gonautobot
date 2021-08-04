package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerOutletsReadReader is a Reader for the DcimPowerOutletsRead structure.
type DcimPowerOutletsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletsReadOK creates a DcimPowerOutletsReadOK with default headers values
func NewDcimPowerOutletsReadOK() *DcimPowerOutletsReadOK {
	return &DcimPowerOutletsReadOK{}
}

/* DcimPowerOutletsReadOK describes a response with status code 200, with default header values.

DcimPowerOutletsReadOK dcim power outlets read o k
*/
type DcimPowerOutletsReadOK struct {
	Payload *models.PowerOutlet
}

func (o *DcimPowerOutletsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-outlets/{id}/][%d] dcimPowerOutletsReadOK  %+v", 200, o.Payload)
}
func (o *DcimPowerOutletsReadOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
