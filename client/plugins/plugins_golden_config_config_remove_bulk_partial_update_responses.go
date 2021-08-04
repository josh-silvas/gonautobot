package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigRemoveBulkPartialUpdateReader is a Reader for the PluginsGoldenConfigConfigRemoveBulkPartialUpdate structure.
type PluginsGoldenConfigConfigRemoveBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigRemoveBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigRemoveBulkPartialUpdateOK creates a PluginsGoldenConfigConfigRemoveBulkPartialUpdateOK with default headers values
func NewPluginsGoldenConfigConfigRemoveBulkPartialUpdateOK() *PluginsGoldenConfigConfigRemoveBulkPartialUpdateOK {
	return &PluginsGoldenConfigConfigRemoveBulkPartialUpdateOK{}
}

/* PluginsGoldenConfigConfigRemoveBulkPartialUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigRemoveBulkPartialUpdateOK plugins golden config config remove bulk partial update o k
*/
type PluginsGoldenConfigConfigRemoveBulkPartialUpdateOK struct {
	Payload *models.ConfigRemove
}

func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/golden-config/config-remove/][%d] pluginsGoldenConfigConfigRemoveBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateOK) GetPayload() *models.ConfigRemove {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigRemove)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
