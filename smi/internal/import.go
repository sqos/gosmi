package internal

import (
	"github.com/sqos/gosmi/types"
)

var importConversions map[types.SmiImport]types.SmiImport = map[types.SmiImport]types.SmiImport{
	// From libsmi
	types.SmiImport{"RFC1155-SMI", "internet"}:      types.SmiImport{"SNMPv2-SMI", "internet"},
	types.SmiImport{"RFC1155-SMI", "directory"}:     types.SmiImport{"SNMPv2-SMI", "directory"},
	types.SmiImport{"RFC1155-SMI", "mgmt"}:          types.SmiImport{"SNMPv2-SMI", "mgmt"},
	types.SmiImport{"RFC1155-SMI", "experimental"}:  types.SmiImport{"SNMPv2-SMI", "experimental"},
	types.SmiImport{"RFC1155-SMI", "private"}:       types.SmiImport{"SNMPv2-SMI", "private"},
	types.SmiImport{"RFC1155-SMI", "enterprises"}:   types.SmiImport{"SNMPv2-SMI", "enterprises"},
	types.SmiImport{"RFC1155-SMI", "IpAddress"}:     types.SmiImport{"SNMPv2-SMI", "IpAddress"},
	types.SmiImport{"RFC1155-SMI", "Counter"}:       types.SmiImport{"SNMPv2-SMI", "Counter32"},
	types.SmiImport{"RFC1155-SMI", "Gauge"}:         types.SmiImport{"SNMPv2-SMI", "Gauge32"},
	types.SmiImport{"RFC1155-SMI", "TimeTicks"}:     types.SmiImport{"SNMPv2-SMI", "TimeTicks"},
	types.SmiImport{"RFC1155-SMI", "Opaque"}:        types.SmiImport{"SNMPv2-SMI", "Opaque"},
	types.SmiImport{"RFC1065-SMI", "internet"}:      types.SmiImport{"SNMPv2-SMI", "internet"},
	types.SmiImport{"RFC1065-SMI", "directory"}:     types.SmiImport{"SNMPv2-SMI", "directory"},
	types.SmiImport{"RFC1065-SMI", "mgmt"}:          types.SmiImport{"SNMPv2-SMI", "mgmt"},
	types.SmiImport{"RFC1065-SMI", "experimental"}:  types.SmiImport{"SNMPv2-SMI", "experimental"},
	types.SmiImport{"RFC1065-SMI", "private"}:       types.SmiImport{"SNMPv2-SMI", "private"},
	types.SmiImport{"RFC1065-SMI", "enterprises"}:   types.SmiImport{"SNMPv2-SMI", "enterprises"},
	types.SmiImport{"RFC1065-SMI", "IpAddress"}:     types.SmiImport{"SNMPv2-SMI", "IpAddress"},
	types.SmiImport{"RFC1065-SMI", "Counter"}:       types.SmiImport{"SNMPv2-SMI", "Counter32"},
	types.SmiImport{"RFC1065-SMI", "Gauge"}:         types.SmiImport{"SNMPv2-SMI", "Gauge32"},
	types.SmiImport{"RFC1065-SMI", "TimeTicks"}:     types.SmiImport{"SNMPv2-SMI", "TimeTicks"},
	types.SmiImport{"RFC1065-SMI", "Opaque"}:        types.SmiImport{"SNMPv2-SMI", "Opaque"},
	types.SmiImport{"RFC1213-MIB", "mib-2"}:         types.SmiImport{"SNMPv2-SMI", "mib-2"},
	types.SmiImport{"RFC1213-MIB", "DisplayString"}: types.SmiImport{"SNMPv2-TC", "DisplayString"},

	// Custom
	types.SmiImport{"RFC1213-MIB", "PhysAddress"}:             types.SmiImport{"SNMPv2-TC", "PhysAddress"},
	types.SmiImport{"RFC1213-MIB", "ifAdminStatus"}:           types.SmiImport{"IF-MIB", "ifAdminStatus"},
	types.SmiImport{"RFC1213-MIB", "ifDescr"}:                 types.SmiImport{"IF-MIB", "ifDescr"},
	types.SmiImport{"RFC1213-MIB", "ifEntry"}:                 types.SmiImport{"IF-MIB", "ifEntry"},
	types.SmiImport{"RFC1213-MIB", "ifInErrors"}:              types.SmiImport{"IF-MIB", "ifInErrors"},
	types.SmiImport{"RFC1213-MIB", "ifIndex"}:                 types.SmiImport{"IF-MIB", "ifIndex"},
	types.SmiImport{"RFC1213-MIB", "ifOperStatus"}:            types.SmiImport{"IF-MIB", "ifOperStatus"},
	types.SmiImport{"RFC1213-MIB", "ifOutDiscards"}:           types.SmiImport{"IF-MIB", "ifOutDiscards"},
	types.SmiImport{"RFC1213-MIB", "ifOutErrors"}:             types.SmiImport{"IF-MIB", "ifOutErrors"},
	types.SmiImport{"RFC1213-MIB", "ifPhysAddress"}:           types.SmiImport{"IF-MIB", "ifPhysAddress"},
	types.SmiImport{"RFC1213-MIB", "interfaces"}:              types.SmiImport{"IF-MIB", "interfaces"},
	types.SmiImport{"RFC1213-MIB", "ip"}:                      types.SmiImport{"IP-MIB", "ip"},
	types.SmiImport{"RFC1213-MIB", "ipAdEntAddr"}:             types.SmiImport{"IP-MIB", "ipAdEntAddr"},
	types.SmiImport{"RFC1213-MIB", "ipAddrEntry"}:             types.SmiImport{"IP-MIB", "ipAddrEntry"},
	types.SmiImport{"RFC1213-MIB", "ipNetToMediaIfIndex"}:     types.SmiImport{"IP-MIB", "ipNetToMediaIfIndex"},
	types.SmiImport{"RFC1213-MIB", "ipNetToMediaNetAddress"}:  types.SmiImport{"IP-MIB", "ipNetToMediaNetAddress"},
	types.SmiImport{"RFC1213-MIB", "ipNetToMediaPhysAddress"}: types.SmiImport{"IP-MIB", "ipNetToMediaPhysAddress"},
	types.SmiImport{"RFC1213-MIB", "ipRouteDest"}:             types.SmiImport{"IP-MIB", "ipRouteDest"},
	types.SmiImport{"RFC1213-MIB", "snmp"}:                    types.SmiImport{"SNMPv2-MIB", "snmp"},
	types.SmiImport{"RFC1213-MIB", "sysContact"}:              types.SmiImport{"SNMPv2-MIB", "sysContact"},
	types.SmiImport{"RFC1213-MIB", "sysDescr"}:                types.SmiImport{"SNMPv2-MIB", "sysDescr"},
	types.SmiImport{"RFC1213-MIB", "sysLocation"}:             types.SmiImport{"SNMPv2-MIB", "sysLocation"},
	types.SmiImport{"RFC1213-MIB", "sysName"}:                 types.SmiImport{"SNMPv2-MIB", "sysName"},
	types.SmiImport{"RFC1213-MIB", "sysObjectID"}:             types.SmiImport{"SNMPv2-MIB", "sysObjectID"},
	types.SmiImport{"RFC1213-MIB", "sysUpTime"}:               types.SmiImport{"SNMPv2-MIB", "sysUpTime"},
	types.SmiImport{"RFC1213-MIB", "system"}:                  types.SmiImport{"SNMPv2-MIB", "system"},
	types.SmiImport{"RFC1213-MIB", "tcpConnLocalAddress"}:     types.SmiImport{"TCP-MIB", "tcpConnLocalAddress"},
	types.SmiImport{"RFC1213-MIB", "tcpConnLocalPort"}:        types.SmiImport{"TCP-MIB", "tcpConnLocalPort"},
	types.SmiImport{"RFC1213-MIB", "tcpConnRemAddress"}:       types.SmiImport{"TCP-MIB", "tcpConnRemAddress"},
	types.SmiImport{"RFC1213-MIB", "tcpConnRemPort"}:          types.SmiImport{"TCP-MIB", "tcpConnRemPort"},
	types.SmiImport{"RFC1213-MIB", "transmission"}:            types.SmiImport{"SNMPv2-SMI", "transmission"},
}
