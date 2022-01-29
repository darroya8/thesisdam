package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for control the food
type SmartContract struct {
	contractapi.Contract
}

//Food describes basic details of what makes up a food
type Shipping struct {
	Adate string `json:"adate"`
	Bbuyer  string `json:"bbuyer"`
	Cseller string `json:"cseller"`
	Dterms string `json:"dterms"`
	Ematerial string `json:"ematerial"`
	Fhscode string `json:"fhscode"`
	Gquantity string `json:"gquantity"`
	Hunitprice string `json:"hunitprice"`
	Ipricetotal string `json:"ipricetotal"`
	Jpaymentterms string `json:"jpaymentterms"`
	Kinvoice string `json:"kinvoice"`
	Lbl string `json:"lbl"`
	Mstatus string `json:"mstatus"`
	Nexitfactorydate string `json:"nexitfactorydate"`
	Oetd string `json:"oetd"`
	Prtd string `json:"prtd"`
	Qeta string `json:"qeta"`
	Rrta string `json:"rrta"`
	Scustomsdate string `json:"scustomsdate"`
	Tarrivaldate string `json:"tarrivaldate"`
}

func (s *SmartContract) Set(ctx contractapi.TransactionContextInterface, shippingId string, adate string, bbuyer string, cseller string, dterms string, ematerial string, fhscode string, gquantity string, hunitprice string, ipricetotal string, jpaymentterms string, kinvoice string, lbl string, mstatus string, nexitfactorydate string, oetd string, prtd string, qeta string, rrta string, scustomsdate string, tarrivaldate string) error {

	//Validaciones de sintaxis

	//validaciones de negocio

	shipping := Shipping{
		Adate: adate,
		Bbuyer:  bbuyer,
		Cseller: cseller,
		Dterms: dterms,
		Ematerial: ematerial,
		Fhscode: fhscode, 
		Gquantity: gquantity,
		Hunitprice: hunitprice,
		Ipricetotal: ipricetotal,
		Jpaymentterms: jpaymentterms,
		Kinvoice: kinvoice,
		Lbl: lbl, 
		Mstatus: mstatus,
		Nexitfactorydate: nexitfactorydate,
		Oetd: oetd,
		Prtd: prtd,
		Qeta: qeta,
		Rrta: rrta, 
		Scustomsdate: scustomsdate,
		Tarrivaldate: tarrivaldate,
	}

	shippingAsBytes, err := json.Marshal(shipping)
	if err != nil {
		fmt.Printf("Marshal error: %s", err.Error())
		return err
	}

	return ctx.GetStub().PutState(shippingId, shippingAsBytes)
}

func (s *SmartContract) Query(ctx contractapi.TransactionContextInterface, shippingId string) (*Shipping, error) {

	shippingAsBytes, err := ctx.GetStub().GetState(shippingId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if shippingAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", shippingId)
	}

	shipping := new(Shipping)

	err = json.Unmarshal(shippingAsBytes, shipping)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal error. %s", err.Error())
	}

	return shipping, nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create shipping chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting shipping chaincode: %s", err.Error())
	}
}
