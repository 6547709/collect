package main

import (
	"fmt"
	"github.com/soniah/gosnmp"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-snmp-example/config"
	"log"
	"os"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "config/snmp.yaml", "snmp config file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	snmpTarget := viper.GetString("snmp.target")
	snmpPort := viper.GetInt("snmp.port")
	snmpCommunity := viper.GetString("snmp.community")

	// 构建GoSNMP结构体
	// 详细记录数据包
	params := &gosnmp.GoSNMP{
		Target:    snmpTarget,
		Port:      uint16(snmpPort),
		Community: snmpCommunity,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		Logger:    log.New(os.Stdout, "", 0),
	}
	err := params.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer params.Conn.Close()

	// 1.3.6.1.4.1.2021.11.11.0(空闲cpu百分比,ssCpuIdle,GET)
	// 1.3.6.1.4.1.2021.4.5.0(机器总内存单位是字节,memTotalReal,GET)
	oids := []string{"1.3.6.1.4.1.2021.11.11.0", "1.3.6.1.4.1.2021.4.5.0"}
	result, err2 := params.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		switch variable.Type {
		case gosnmp.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
		default:
			fmt.Printf("number: %d\n", gosnmp.ToBigInt(variable.Value))
		}
	}
}
