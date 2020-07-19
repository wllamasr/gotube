package custom_validations

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/wllamasr/golangtube/config/db"
	"strings"
)

/*
Validates if the value exist in the given table.field

usage:
```
type User struct{
	Email `validate:"unrepeated=users.email"`
}
```
*/
func ValidateUnrepeated(fl validator.FieldLevel) bool {
	param := strings.Split(fl.Param(), ".")

	if len(param) != 2 {
		return false
	}

	table := param[0]
	field := param[1]
	query := fmt.Sprintf("%s=?", field)

	var count int64

	if err := db.Client.Table(table).Where(query, fl.Field().String()).Count(&count).Error; err != nil {
		panic(fmt.Sprintf("Erorr in your SQL: %s", err.Error()))
		return false
	}

	if count > 0 {
		return false
	}

	return true
}
