package view

type AttachtagsLambdaPayload struct {
	Pin  *AttachTagsLambdaPayloadPin `json:"pin"`
	Tags []string                    `json:"tags"`
}
