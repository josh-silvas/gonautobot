package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRearPortTemplatesUpdateReader is a Reader for the DcimRearPortTemplatesUpdate structure.
type DcimRearPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortTemplatesUpdateOK creates a DcimRearPortTemplatesUpdateOK with default headers values
func NewDcimRearPortTemplatesUpdateOK() *DcimRearPortTemplatesUpdateOK {
	return &DcimRearPortTemplatesUpdateOK{}
}

/* DcimRearPortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimRearPortTemplatesUpdateOK dcim rear port templates update o k
*/
type DcimRearPortTemplatesUpdateOK struct {
	Payload *models.RearPortTemplate
}

func (o *DcimRearPortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRearPortTemplatesUpdateOK) GetPayload() *models.RearPortTemplate {
	return o.Payload
}

func (o *DcimRearPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
