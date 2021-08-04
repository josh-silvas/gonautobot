package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimInventoryItemsBulkPartialUpdateReader is a Reader for the DcimInventoryItemsBulkPartialUpdate structure.
type DcimInventoryItemsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInventoryItemsBulkPartialUpdateOK creates a DcimInventoryItemsBulkPartialUpdateOK with default headers values
func NewDcimInventoryItemsBulkPartialUpdateOK() *DcimInventoryItemsBulkPartialUpdateOK {
	return &DcimInventoryItemsBulkPartialUpdateOK{}
}

/* DcimInventoryItemsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemsBulkPartialUpdateOK dcim inventory items bulk partial update o k
*/
type DcimInventoryItemsBulkPartialUpdateOK struct {
	Payload *models.InventoryItem
}

func (o *DcimInventoryItemsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-items/][%d] dcimInventoryItemsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInventoryItemsBulkPartialUpdateOK) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
