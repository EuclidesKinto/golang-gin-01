package request

type UpdateTagsRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required, min=2,max=10" json:"name"`
}