package ncode

import (
	"strconv"
)

var placeValue = [10]int{10,9,8,7,6,5,4,3,2,1}

const (
	residualConst = 11
	controlDigitIndex = 9 
)


type NationalID struct{
	stringFormat string
	digitFormat [10]int
	Validity bool
}

func NewNationalID(id string) *NationalID{
	nid := &NationalID{stringFormat: id}
	nid.Validity = nid.isValid()
	return nid
}

func (id *NationalID) isValid () bool{
	if len(id.stringFormat) == 10{
		for index, val := range id.stringFormat{
			digit,err := strconv.Atoi(string(val))
			if err != nil{
				return false
			}
			id.digitFormat[index] = digit
		}
		if id.isCheckSumValid(){
			id.Validity = true
			return true
		}
	}
	return false
}

func (id *NationalID) isCheckSumValid() bool{
	var (
		sum int
		resid int
	)
	
	for index := 0 ; index < 9 ; index++{
		sum = sum + id.digitFormat[index]*placeValue[index]
	
	}
	resid = sum % residualConst

	controlDigit := id.digitFormat[controlDigitIndex]

	if resid < 2 {
		if resid == controlDigit{
			return true
		}
	}else{
		if controlDigit == (11 -resid){
			return true
		}
	}
	return false
}