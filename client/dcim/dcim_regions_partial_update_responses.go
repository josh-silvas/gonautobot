package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRegionsPartialUpdateReader is a Reader for the DcimRegionsPartialUpdate structure.
type DcimRegionsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRegionsPartialUpdateOK creates a DcimRegionsPartialUpdateOK with default headers values
func NewDcimRegionsPartialUpdateOK() *DcimRegionsPartialUpdateOK {
	return &DcimRegionsPartialUpdateOK{}
}

/* DcimRegionsPartialUpdateOK describes a response with status code 200, with default header values.

DcimRegionsPartialUpdateOK dcim regions partial update o k
*/
type DcimRegionsPartialUpdateOK struct {
	Payload *models.Region
}

func (o *DcimRegionsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/regions/{id}/][%d] dcimRegionsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRegionsPartialUpdateOK) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
