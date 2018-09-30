package main

import (
	"fmt"
	"github.com/soniah/gosnmp"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-snmp-collect-network-devices-information/config"
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
	cpuCores:=viper.GetString("snmp.oid.cpu.cpuCores")
	cpuIdle:=viper.GetString("snmp.oid.cpu.cpuIdle")
	memTotal:=viper.GetString("snmp.oid.memory.memTotal")
	memAvail:=viper.GetString("snmp.oid.memory.memAvail")
	diskTotal:=viper.GetString("snmp.oid.disk.diskTotal")
	diskPercent:=viper.GetString("snmp.oid.disk.diskPercent")

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

	oids := []string{cpuIdle,memTotal,memAvail,diskTotal,diskPercent}
	resultGet, err2 := params.Get(oids) // Get() accepts up to g.MAX_OIDS
	resultWalk,err2:= params.Walk(cpuCores)
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range resultGet.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		switch variable.Type {
		// if value is zero,
		case gosnmp.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
		default:
			fmt.Printf("number: %d\n", gosnmp.ToBigInt(variable.Value))
		}
	}
}
