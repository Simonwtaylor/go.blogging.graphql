// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package graph

import "github.com/jinzhu/gorm"

type Resolver struct {
	DB *gorm.DB
}
