package request

import "crmeb_go/internal/validation"

type CreateProductParams struct {
	CreateProduct *validation.CreateProductParam
	BaseServiceParams
}
