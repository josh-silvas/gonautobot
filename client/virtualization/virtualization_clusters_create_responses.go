package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationClustersCreateReader is a Reader for the VirtualizationClustersCreate structure.
type VirtualizationClustersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVirtualizationClustersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClustersCreateCreated creates a VirtualizationClustersCreateCreated with default headers values
func NewVirtualizationClustersCreateCreated() *VirtualizationClustersCreateCreated {
	return &VirtualizationClustersCreateCreated{}
}

/* VirtualizationClustersCreateCreated describes a response with status code 201, with default header values.

VirtualizationClustersCreateCreated virtualization clusters create created
*/
type VirtualizationClustersCreateCreated struct {
	Payload *models.Cluster
}

func (o *VirtualizationClustersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /virtualization/clusters/][%d] virtualizationClustersCreateCreated  %+v", 201, o.Payload)
}
func (o *VirtualizationClustersCreateCreated) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
