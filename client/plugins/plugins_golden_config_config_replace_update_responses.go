package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigReplaceUpdateReader is a Reader for the PluginsGoldenConfigConfigReplaceUpdate structure.
type PluginsGoldenConfigConfigReplaceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigReplaceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigReplaceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigReplaceUpdateOK creates a PluginsGoldenConfigConfigReplaceUpdateOK with default headers values
func NewPluginsGoldenConfigConfigReplaceUpdateOK() *PluginsGoldenConfigConfigReplaceUpdateOK {
	return &PluginsGoldenConfigConfigReplaceUpdateOK{}
}

/* PluginsGoldenConfigConfigReplaceUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigReplaceUpdateOK plugins golden config config replace update o k
*/
type PluginsGoldenConfigConfigReplaceUpdateOK struct {
	Payload *models.ConfigReplace
}

func (o *PluginsGoldenConfigConfigReplaceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/config-replace/{id}/][%d] pluginsGoldenConfigConfigReplaceUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigReplaceUpdateOK) GetPayload() *models.ConfigReplace {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigReplaceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigReplace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
