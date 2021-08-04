package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInterfaceTemplatesUpdateReader is a Reader for the DcimInterfaceTemplatesUpdate structure.
type DcimInterfaceTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfaceTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesUpdateOK creates a DcimInterfaceTemplatesUpdateOK with default headers values
func NewDcimInterfaceTemplatesUpdateOK() *DcimInterfaceTemplatesUpdateOK {
	return &DcimInterfaceTemplatesUpdateOK{}
}

/* DcimInterfaceTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimInterfaceTemplatesUpdateOK dcim interface templates update o k
*/
type DcimInterfaceTemplatesUpdateOK struct {
	Payload *models.InterfaceTemplate
}

func (o *DcimInterfaceTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInterfaceTemplatesUpdateOK) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
