package main

import (
	"net/http"
	conekta "github.com/conekta/conekta-go"
    "github.com/conekta/conekta-go/order"
    "log"
    "encoding/json"
 	//"io/ioutil"
 	"fmt"
)


func PayOrder(w http.ResponseWriter, r *http.Request){
	conekta.APIKey = "key_YaDNzhudzCgf9Hw5jxz5gQ"
	var params PayParams
	error_response := ""

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil{
		log.Fatalln(err)
	}

	defer r.Body.Close()//cerrar la peticion

	/************ CustomerParams para conekta *******************/
	Address := &conekta.Address{
        Street1:    "Street1",
        Street2:    "Street2",
        City:       "City",
        State:      "State",
        Country:    "Country",
        PostalCode: "PostalCode", 
    }

	customerParams := &conekta.CustomerParams{
        Name:  params.Name,
        Email: params.Email,
        Phone: "6182611723",  
    }

    lineItemParams := &conekta.LineItemsParams{
        Name:      params.CourseName,
        UnitPrice: params.Price,
        Quantity:  1, 
    }

    shippingParams := &conekta.ShippingLinesParams{
        Amount:         0,
        TrackingNumber: "123",
        Carrier:        "Test Carrier",
        Method:         "method", 
    }

    shippingContactParams := &conekta.ShippingContactParams{
        Phone:          "6182611723",
        Receiver:       params.Name,
        BetweenStreets: "BetweenStreets",
        Address:        Address,  
    }

	chargeParams := &conekta.ChargeParams{
        PaymentMethod: &conekta.PaymentMethodParams{
            Type:    params.PaymentSources.PaymentType,
            TokenID: params.PaymentSources.TokenId,  
            //ExpiresAt: time.Now().AddDate(0, 0, 90).Unix(),
            //ExpiresAt es para usar en caso de enviar cargo tipo oxxo. Reemplazar por TokenID
        },
    }

    orderParams := &conekta.OrderParams{}
    orderParams.Currency = "MXN"
    orderParams.CustomerInfo = customerParams
    orderParams.PreAuth =  true //en caso de no pasar este parametro, procesar la orden no podrá ser posible
    orderParams.LineItems = append(orderParams.LineItems, lineItemParams)
    orderParams.ShippingLines = append(orderParams.ShippingLines, shippingParams)
    orderParams.ShippingContact = shippingContactParams
    orderParams.Charges = append(orderParams.Charges, chargeParams)

    ord, err := order.Create(orderParams)
	if err != nil {
		error_response = err.(conekta.Error).Details[0].Code
		//do something
	} else {
		orderId := ord.ID
		chargeId := ord.Charges.Data[0].ID
		//do something
		fmt.Println(orderId)
		fmt.Println(chargeId)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Endcoding, Content-Type, Content-Length")
	w.Header().Set("Accept", "application/json"); 
	w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(error_response)
}