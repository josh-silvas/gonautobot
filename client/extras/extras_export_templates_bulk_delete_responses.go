package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasExportTemplatesBulkDeleteReader is a Reader for the ExtrasExportTemplatesBulkDelete structure.
type ExtrasExportTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasExportTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasExportTemplatesBulkDeleteNoContent creates a ExtrasExportTemplatesBulkDeleteNoContent with default headers values
func NewExtrasExportTemplatesBulkDeleteNoContent() *ExtrasExportTemplatesBulkDeleteNoContent {
	return &ExtrasExportTemplatesBulkDeleteNoContent{}
}

/* ExtrasExportTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasExportTemplatesBulkDeleteNoContent extras export templates bulk delete no content
*/
type ExtrasExportTemplatesBulkDeleteNoContent struct {
}

func (o *ExtrasExportTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/][%d] extrasExportTemplatesBulkDeleteNoContent ", 204)
}

func (o *ExtrasExportTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
