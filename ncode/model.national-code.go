package ncode

import (
	"strconv"
)

type NationalCode struct{
	StringFormat string `json:"id"`
	IsValid bool `json:"isValid"`
	City string	  `json:"city"`
	digitalFormat [10]int
}

const (
	_divider int = 11
	_controllerDigit int = 9
	_nationalCodeLength int = 10

)

func NewNationalCode(code string)  *NationalCode{
	nationalCode := &NationalCode{StringFormat: code}
	nationalCode.IsValid=nationalCode.isValid()
	if nationalCode.IsValid {
		nationalCode.City=nationalCode.whichCity()
	}
	return nationalCode
}

func (n *NationalCode) isValid() bool{
	if len(n.StringFormat) == _nationalCodeLength{
		for index, value := range n.StringFormat {
			integer, err := strconv.Atoi(string(value))
			if err != nil{
				return false
			}
			n.digitalFormat[index]=integer
		}
		if n.isCheckSumValid(){
			return true
		}

	}
	return false
}

func (n *NationalCode) isCheckSumValid() bool{
	elements := [10]int{10,9,8,7,6,5,4,3,2,1}
	sum := 0
	for step,value := range elements{
		if step == _controllerDigit{
			break
		}
		sum += n.digitalFormat[step] * value
	}
	remain := sum % _divider
	if remain < 2 {
		if remain == n.digitalFormat[_controllerDigit]{
			return true
		}
	} else{
		if _divider-remain == n.digitalFormat[_controllerDigit]{
			return true
		}
	}
	return false

}

func (n *NationalCode)whichCity() string{
	cityCode := n.StringFormat[:3]

	switch cityCode {
	case "412":
		return "Boroujerd"
	case "413":
		return "Boroujerd"
	case "406":
		return "Khoram Abad"
	case "407":
		return "Khoram Abad"
	case "421":
		return "Doroud"
	case "498":
		return "Babolsar"
	case "483":
		return "Chalous"
	case "227":
		return "Ramsar"
	case "627":
		return "Kelardasht"
	case "001":
		return "Tehran"
	default:
		return "Not Found"		
	}
}