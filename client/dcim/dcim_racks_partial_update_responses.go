package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRacksPartialUpdateReader is a Reader for the DcimRacksPartialUpdate structure.
type DcimRacksPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRacksPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRacksPartialUpdateOK creates a DcimRacksPartialUpdateOK with default headers values
func NewDcimRacksPartialUpdateOK() *DcimRacksPartialUpdateOK {
	return &DcimRacksPartialUpdateOK{}
}

/* DcimRacksPartialUpdateOK describes a response with status code 200, with default header values.

DcimRacksPartialUpdateOK dcim racks partial update o k
*/
type DcimRacksPartialUpdateOK struct {
	Payload *models.Rack
}

func (o *DcimRacksPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/racks/{id}/][%d] dcimRacksPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRacksPartialUpdateOK) GetPayload() *models.Rack {
	return o.Payload
}

func (o *DcimRacksPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
