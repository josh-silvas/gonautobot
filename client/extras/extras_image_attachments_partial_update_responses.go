package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasImageAttachmentsPartialUpdateReader is a Reader for the ExtrasImageAttachmentsPartialUpdate structure.
type ExtrasImageAttachmentsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasImageAttachmentsPartialUpdateOK creates a ExtrasImageAttachmentsPartialUpdateOK with default headers values
func NewExtrasImageAttachmentsPartialUpdateOK() *ExtrasImageAttachmentsPartialUpdateOK {
	return &ExtrasImageAttachmentsPartialUpdateOK{}
}

/* ExtrasImageAttachmentsPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsPartialUpdateOK extras image attachments partial update o k
*/
type ExtrasImageAttachmentsPartialUpdateOK struct {
	Payload *models.ImageAttachment
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extrasImageAttachmentsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasImageAttachmentsPartialUpdateOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
