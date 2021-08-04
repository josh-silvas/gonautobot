package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRearPortTemplatesReadReader is a Reader for the DcimRearPortTemplatesRead structure.
type DcimRearPortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortTemplatesReadOK creates a DcimRearPortTemplatesReadOK with default headers values
func NewDcimRearPortTemplatesReadOK() *DcimRearPortTemplatesReadOK {
	return &DcimRearPortTemplatesReadOK{}
}

/* DcimRearPortTemplatesReadOK describes a response with status code 200, with default header values.

DcimRearPortTemplatesReadOK dcim rear port templates read o k
*/
type DcimRearPortTemplatesReadOK struct {
	Payload *models.RearPortTemplate
}

func (o *DcimRearPortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesReadOK  %+v", 200, o.Payload)
}
func (o *DcimRearPortTemplatesReadOK) GetPayload() *models.RearPortTemplate {
	return o.Payload
}

func (o *DcimRearPortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
