package user

import (
	"github.com/Kamva/mgm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	name             string `json:"name" bson:"name"`
	age              string `json:"age" bson:"age"`
}

type M map[string]interface{}

var internalError = M{"message": "internal error"}
var userNotFound = M{"message": "user not found"}

func Create(c echo.Context) error {
	requestData := &Request{}

	user := &User{
		name: requestData.name,
		age:  requestData.age,
	}

	err := mgm.Coll(user).Save(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	return c.JSON(http.StatusOK, user)
}

type Request struct {
	name string `json:"name"`
	age  string `json:"age"`
}

func (r *Request) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.name, validation.Required),
		validation.Field(&r.age, validation.Required),
	)
}
