package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationClustersPartialUpdateReader is a Reader for the VirtualizationClustersPartialUpdate structure.
type VirtualizationClustersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClustersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClustersPartialUpdateOK creates a VirtualizationClustersPartialUpdateOK with default headers values
func NewVirtualizationClustersPartialUpdateOK() *VirtualizationClustersPartialUpdateOK {
	return &VirtualizationClustersPartialUpdateOK{}
}

/* VirtualizationClustersPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationClustersPartialUpdateOK virtualization clusters partial update o k
*/
type VirtualizationClustersPartialUpdateOK struct {
	Payload *models.Cluster
}

func (o *VirtualizationClustersPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/clusters/{id}/][%d] virtualizationClustersPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationClustersPartialUpdateOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
