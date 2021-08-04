package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimFrontPortTemplatesReadReader is a Reader for the DcimFrontPortTemplatesRead structure.
type DcimFrontPortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortTemplatesReadOK creates a DcimFrontPortTemplatesReadOK with default headers values
func NewDcimFrontPortTemplatesReadOK() *DcimFrontPortTemplatesReadOK {
	return &DcimFrontPortTemplatesReadOK{}
}

/* DcimFrontPortTemplatesReadOK describes a response with status code 200, with default header values.

DcimFrontPortTemplatesReadOK dcim front port templates read o k
*/
type DcimFrontPortTemplatesReadOK struct {
	Payload *models.FrontPortTemplate
}

func (o *DcimFrontPortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesReadOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortTemplatesReadOK) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
