package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInterfaceTemplatesPartialUpdateReader is a Reader for the DcimInterfaceTemplatesPartialUpdate structure.
type DcimInterfaceTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfaceTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesPartialUpdateOK creates a DcimInterfaceTemplatesPartialUpdateOK with default headers values
func NewDcimInterfaceTemplatesPartialUpdateOK() *DcimInterfaceTemplatesPartialUpdateOK {
	return &DcimInterfaceTemplatesPartialUpdateOK{}
}

/* DcimInterfaceTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimInterfaceTemplatesPartialUpdateOK dcim interface templates partial update o k
*/
type DcimInterfaceTemplatesPartialUpdateOK struct {
	Payload *models.InterfaceTemplate
}

func (o *DcimInterfaceTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInterfaceTemplatesPartialUpdateOK) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
