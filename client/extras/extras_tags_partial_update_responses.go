package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasTagsPartialUpdateReader is a Reader for the ExtrasTagsPartialUpdate structure.
type ExtrasTagsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasTagsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasTagsPartialUpdateOK creates a ExtrasTagsPartialUpdateOK with default headers values
func NewExtrasTagsPartialUpdateOK() *ExtrasTagsPartialUpdateOK {
	return &ExtrasTagsPartialUpdateOK{}
}

/* ExtrasTagsPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasTagsPartialUpdateOK extras tags partial update o k
*/
type ExtrasTagsPartialUpdateOK struct {
	Payload *models.Tag
}

func (o *ExtrasTagsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/tags/{id}/][%d] extrasTagsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasTagsPartialUpdateOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *ExtrasTagsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
