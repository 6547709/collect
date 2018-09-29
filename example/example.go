package main

import (
	"fmt"
	"log"

	g "github.com/soniah/gosnmp"
)

/**
* @author wangpengcheng@ccssoft.com.cn
* @date 18-9-29 下午3:58
 */

func main() {
	// Default is a pointer to a GoSNMP struct that contains sensible defaults
	// eg port 161, community public, etc
	g.Default.Target = "172.18.1.133"
	err := g.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	// 1.3.6.1.2.1.1.4.0(系统联系人,sysContact,GET)
	// 1.3.6.1.2.1.1.7.0(机器提供的服务,SysService,GET)
	// 1.3.6.1.4.1.2021.11.11.0(空闲cpu百分比,ssCpuIdle,GET)
	// 1.3.6.1.2.1.25.2.2.0(内存大小,hrMemorySize,GET)
	oids := []string{"1.3.6.1.4.1.2021.11.11.0", "1.3.6.1.2.1.25.2.2.0"}
	result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
		case g.OctetString:
			bytes := variable.Value.([]byte)
			fmt.Printf("string: %s\n", string(bytes))
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.
			fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
		}
	}
}
