package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimSitesPartialUpdateReader is a Reader for the DcimSitesPartialUpdate structure.
type DcimSitesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimSitesPartialUpdateOK creates a DcimSitesPartialUpdateOK with default headers values
func NewDcimSitesPartialUpdateOK() *DcimSitesPartialUpdateOK {
	return &DcimSitesPartialUpdateOK{}
}

/* DcimSitesPartialUpdateOK describes a response with status code 200, with default header values.

DcimSitesPartialUpdateOK dcim sites partial update o k
*/
type DcimSitesPartialUpdateOK struct {
	Payload *models.Site
}

func (o *DcimSitesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/sites/{id}/][%d] dcimSitesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimSitesPartialUpdateOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *DcimSitesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
