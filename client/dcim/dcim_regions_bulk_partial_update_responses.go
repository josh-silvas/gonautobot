package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRegionsBulkPartialUpdateReader is a Reader for the DcimRegionsBulkPartialUpdate structure.
type DcimRegionsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRegionsBulkPartialUpdateOK creates a DcimRegionsBulkPartialUpdateOK with default headers values
func NewDcimRegionsBulkPartialUpdateOK() *DcimRegionsBulkPartialUpdateOK {
	return &DcimRegionsBulkPartialUpdateOK{}
}

/* DcimRegionsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRegionsBulkPartialUpdateOK dcim regions bulk partial update o k
*/
type DcimRegionsBulkPartialUpdateOK struct {
	Payload *models.Region
}

func (o *DcimRegionsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/regions/][%d] dcimRegionsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRegionsBulkPartialUpdateOK) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
