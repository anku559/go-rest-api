package category

import (
	"github.com/go-playground/validator/v10"
)

type CreateCategoryValidation struct {
	Name string `json:"name" validate:"alphanum,min=3,required" error:"Name should be alphanumeric and have a minimum length of 3 characters"`
	Slug string `json:"slug" validate:"alphanum,min=3,omitempty" error:"Slug should be alphanumeric and have a minimum length of 3 characters"`
}

func (c *CreateCategoryValidation) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

// func (c *CreateCategoryValidation) ValidateCreateCategory() (string, error) {
// 	validate := validator.New()
// 	err := validate.Struct(c)

// 	if err != nil {
// 		log.Fatal("Validation failed")

// 		for _, e := range err.(validator.ValidationErrors) {
// 			fmt.Printf("Field: %s, Error: %s\n", e.Field(), e.Tag())
// 		}
// 		return "", err
// 	} else {
// 		fmt.Println("Validation succeeded")
// 		return "", nil
// 	}
// }
