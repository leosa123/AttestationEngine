package netconfprotocol

import(
	"fmt"
	"log"

	"a10/structures"
)


func Registration() (structures.Protocol) {
	intents := []string{"null/good","null/test"}

	return structures.Protocol{"A10NETCONF","POC protocol module for NetConf",Call, intents}
}

func Call(e structures.Element, p structures.Policy, s structures.Session, cps map[string]interface{}) (map[string]interface{}, string) {

	// Create a test body

	rtn := map[string]interface{}{
		 "foo":"bar",
		 "calling": fmt.Sprintf("with protocol %v I would send an intent to %v",e.Protocol,p.Intent),
		 "aNumber": 42,
	}
	
	// Check if the policy intent was null/null, if so then return with the bodytype being set to null/test
	// or error if the above is false.
	//
	// Claim bodytype should be set to error and a ClaimError structure returned in an error field

	if p.Intent=="null/null" {
		log.Println(" null call worked ")
		rtn["worked"] = true
		return rtn,"null/test"
	} else {
		log.Println(" null call bad error ")	
		rtn["error"] = "Error here"
		return rtn,structures.CLAIMERROR
	}
}