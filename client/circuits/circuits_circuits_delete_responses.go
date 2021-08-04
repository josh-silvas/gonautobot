package circuits

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CircuitsCircuitsDeleteReader is a Reader for the CircuitsCircuitsDelete structure.
type CircuitsCircuitsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCircuitsCircuitsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitsDeleteNoContent creates a CircuitsCircuitsDeleteNoContent with default headers values
func NewCircuitsCircuitsDeleteNoContent() *CircuitsCircuitsDeleteNoContent {
	return &CircuitsCircuitsDeleteNoContent{}
}

/* CircuitsCircuitsDeleteNoContent describes a response with status code 204, with default header values.

CircuitsCircuitsDeleteNoContent circuits circuits delete no content
*/
type CircuitsCircuitsDeleteNoContent struct {
}

func (o *CircuitsCircuitsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/circuits/{id}/][%d] circuitsCircuitsDeleteNoContent ", 204)
}

func (o *CircuitsCircuitsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
