package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInterfacesReadReader is a Reader for the DcimInterfacesRead structure.
type DcimInterfacesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfacesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfacesReadOK creates a DcimInterfacesReadOK with default headers values
func NewDcimInterfacesReadOK() *DcimInterfacesReadOK {
	return &DcimInterfacesReadOK{}
}

/* DcimInterfacesReadOK describes a response with status code 200, with default header values.

DcimInterfacesReadOK dcim interfaces read o k
*/
type DcimInterfacesReadOK struct {
	Payload *models.Interface
}

func (o *DcimInterfacesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/interfaces/{id}/][%d] dcimInterfacesReadOK  %+v", 200, o.Payload)
}
func (o *DcimInterfacesReadOK) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
