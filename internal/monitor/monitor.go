package monitor

import "github.com/free5gc/amf/internal/context"

func GetAmountOfSuccessRegistration() float64 {
	return context.AMF_Self().RegisSuccess
}

func GetAmountOfReceivedRegistration() float64 {
	return context.AMF_Self().RegisTry
}

func IncAmountOfSuccessRegistration() {
	context.AMF_Self().Lock.Lock()
	context.AMF_Self().RegisSuccess++
	context.AMF_Self().Lock.Unlock()
}

func IncAmountOfReceivedRegistration() {
	context.AMF_Self().Lock.Lock()
	context.AMF_Self().RegisTry++
	context.AMF_Self().Lock.Unlock()
}
