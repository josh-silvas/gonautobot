package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerPortTemplatesBulkPartialUpdateReader is a Reader for the DcimPowerPortTemplatesBulkPartialUpdate structure.
type DcimPowerPortTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortTemplatesBulkPartialUpdateOK creates a DcimPowerPortTemplatesBulkPartialUpdateOK with default headers values
func NewDcimPowerPortTemplatesBulkPartialUpdateOK() *DcimPowerPortTemplatesBulkPartialUpdateOK {
	return &DcimPowerPortTemplatesBulkPartialUpdateOK{}
}

/* DcimPowerPortTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesBulkPartialUpdateOK dcim power port templates bulk partial update o k
*/
type DcimPowerPortTemplatesBulkPartialUpdateOK struct {
	Payload *models.PowerPortTemplate
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/][%d] dcimPowerPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
