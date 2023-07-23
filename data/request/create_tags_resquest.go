package request

type CreateTagsRequest struct {
	Name string `validate:"required,min=2,max=200" json:"name"`
}
