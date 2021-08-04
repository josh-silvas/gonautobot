package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasImageAttachmentsReadReader is a Reader for the ExtrasImageAttachmentsRead structure.
type ExtrasImageAttachmentsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasImageAttachmentsReadOK creates a ExtrasImageAttachmentsReadOK with default headers values
func NewExtrasImageAttachmentsReadOK() *ExtrasImageAttachmentsReadOK {
	return &ExtrasImageAttachmentsReadOK{}
}

/* ExtrasImageAttachmentsReadOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsReadOK extras image attachments read o k
*/
type ExtrasImageAttachmentsReadOK struct {
	Payload *models.ImageAttachment
}

func (o *ExtrasImageAttachmentsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/image-attachments/{id}/][%d] extrasImageAttachmentsReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasImageAttachmentsReadOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
