package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInterfaceTemplatesBulkUpdateReader is a Reader for the DcimInterfaceTemplatesBulkUpdate structure.
type DcimInterfaceTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfaceTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesBulkUpdateOK creates a DcimInterfaceTemplatesBulkUpdateOK with default headers values
func NewDcimInterfaceTemplatesBulkUpdateOK() *DcimInterfaceTemplatesBulkUpdateOK {
	return &DcimInterfaceTemplatesBulkUpdateOK{}
}

/* DcimInterfaceTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimInterfaceTemplatesBulkUpdateOK dcim interface templates bulk update o k
*/
type DcimInterfaceTemplatesBulkUpdateOK struct {
	Payload *models.InterfaceTemplate
}

func (o *DcimInterfaceTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/interface-templates/][%d] dcimInterfaceTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInterfaceTemplatesBulkUpdateOK) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
