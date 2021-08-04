package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersGroupsUpdateReader is a Reader for the UsersGroupsUpdate structure.
type UsersGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersGroupsUpdateOK creates a UsersGroupsUpdateOK with default headers values
func NewUsersGroupsUpdateOK() *UsersGroupsUpdateOK {
	return &UsersGroupsUpdateOK{}
}

/* UsersGroupsUpdateOK describes a response with status code 200, with default header values.

UsersGroupsUpdateOK users groups update o k
*/
type UsersGroupsUpdateOK struct {
	Payload *models.Group
}

func (o *UsersGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/groups/{id}/][%d] usersGroupsUpdateOK  %+v", 200, o.Payload)
}
func (o *UsersGroupsUpdateOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *UsersGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
