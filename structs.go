package main

/***************** estructuras para parametos ****************************/
type PayParams struct{
	Name string `json:"name"`
	Email string `json:"email"`
	CourseName string `json:"course_name"`
	Price int64 `json:"price"`
	PaymentSources PaymentStruct `json:"payment_sources"`
}

type PaymentStruct struct{
	TokenId string `json:"token_id"`
	PaymentType string `json:"payment_type"`
}