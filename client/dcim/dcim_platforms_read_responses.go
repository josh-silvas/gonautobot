package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPlatformsReadReader is a Reader for the DcimPlatformsRead structure.
type DcimPlatformsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPlatformsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPlatformsReadOK creates a DcimPlatformsReadOK with default headers values
func NewDcimPlatformsReadOK() *DcimPlatformsReadOK {
	return &DcimPlatformsReadOK{}
}

/* DcimPlatformsReadOK describes a response with status code 200, with default header values.

DcimPlatformsReadOK dcim platforms read o k
*/
type DcimPlatformsReadOK struct {
	Payload *models.Platform
}

func (o *DcimPlatformsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/platforms/{id}/][%d] dcimPlatformsReadOK  %+v", 200, o.Payload)
}
func (o *DcimPlatformsReadOK) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
