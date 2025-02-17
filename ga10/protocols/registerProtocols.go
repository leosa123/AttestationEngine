package protocols

import (
	"a10/operations"
	
	"a10/protocols/a10httprestv2"
	"a10/protocols/nullprotocol"
	"a10/protocols/netconfprotocol"

)


func RegisterProtocols() {
	operations.AddProtocol(a10httprestv2.Registration())
	operations.AddProtocol(nullprotocol.Registration())
	operations.AddProtocol(netconfprotocol.Registration())

}
