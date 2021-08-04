package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerOutletsPartialUpdateReader is a Reader for the DcimPowerOutletsPartialUpdate structure.
type DcimPowerOutletsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletsPartialUpdateOK creates a DcimPowerOutletsPartialUpdateOK with default headers values
func NewDcimPowerOutletsPartialUpdateOK() *DcimPowerOutletsPartialUpdateOK {
	return &DcimPowerOutletsPartialUpdateOK{}
}

/* DcimPowerOutletsPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletsPartialUpdateOK dcim power outlets partial update o k
*/
type DcimPowerOutletsPartialUpdateOK struct {
	Payload *models.PowerOutlet
}

func (o *DcimPowerOutletsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlets/{id}/][%d] dcimPowerOutletsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerOutletsPartialUpdateOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
