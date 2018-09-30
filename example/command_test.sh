#!/bin/bash

snmpget -c ccssoft -v 2c paas.m8.ccs 1.3.6.1.2.1.25.3.3.1.2

# snmpwalk命令则是测试系统各种信息最有效的方法
# 获取cpu核心数量
snmpwalk -c ccssoft -v 2c paas.m8.ccs 1.3.6.1.2.1.25.3.3.1.2
snmpwalk -c ccssoft -v 2c paas.m8.ccs hrProcessorLoad

# 获取磁盘使用率
snmpwalk -v 2c -c ccssoft localhost .1.3.6.1.4.1.2021.9.1.9
