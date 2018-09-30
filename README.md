### go-snmp-example
golang+snmp

#### SNMP定义
```markdown
SNMP协议的全称:Simple Network Management Protocol(简单网络管理协议);

SNMP的主要功能: 通过应答POLLING(轮询)来反馈当前设备状态;

SNMP的工作方式: 管理员需要像设备获取数据,所以SNMP提供了"读"操作;管理员需要向设备执行设置操作,所以SNMP提供了"写"操作;
设备需要在重要状况改变的时候,向管理员通报事件的发生,所以SNMP提供了"Trap"操作;

SNMP被设计为工作在TCP/IP协议族上.SNMP基于TCP/IP协议工作,对网络中支持SNMP协议的设备进行管理.所有支持SNMP协议的设备
都提供SNMP这个统一界面，使得管理员可以使用统一的操作进行管理，而不必理会设备是什么类型、是哪个厂家生产的.
```
#### SNMP技术术语
```markdown
SNMP：Simple Network Management Protocol(简单网络管理协议)，是一个标准的用于管理基于IP网络上设备的协议。

MIB：Management Information Base(管理信息库)，定义代理进程中所有可被查询和修改的参数。

SMI：Structure of Management Information(管理信息结构)，SMI定义了SNMP中使用到的ASN.1类型、语法，并定义了SNMP中使用到的类型、宏、符号等。SMI用于后续协议的描述和MIB的定义。每个版本的SNMP都可能定义自己的SMI。
```
#### 安装snmp
```bash
# 安装
yum install net-snmp net-snmp-utils net-snmp*

# 配置
vim /etc/snmp/snmpd.conf
com2sec notConfigUser  default       ccssoft
view    all           included   .1
access  notConfigGroup ""      any       noauth    exact  all none none
includeAllDisks
rocommunity ccssoft
disk /
disk /home

# 启动snmp服务
systemctl enable snmpd
systemctl start snmpd

# 确保iptables防火墙开放了udp 161端口的访问权限
# 配置防火墙规则运行snmp端口161
iptables -A INPUT -p tcp -m state --state NEW -m tcp --dport 161  -j ACCEPT
iptables -A INPUT -p udp -m state --state NEW -m udp --dport 161  -j ACCEPT
systemctl start snmpd
systemctl enable snmpd
iptables -D INPUT -p tcp -m state --state NEW -m tcp --dport 161  -j ACCEPT
iptables -D INPUT -p udp -m state --state NEW -m udp --dport 161  -j ACCEPT
iptables -I INPUT -p tcp -m state --state NEW -m tcp --dport 161  -j ACCEPT
iptables -I INPUT -p udp -m state --state NEW -m udp --dport 161  -j ACCEPT

```

#### 使用snmp
```bash
# 查看snmp版本
[root@paas ~]# snmpget --version
NET-SNMP version: 5.7.2
[root@paas ~]# 

# 查看一下安装的snmp软件包
rpm -qa | grep net-snmp*

snmpget -c ccssoft -v 2c localhost .1.3.6.1.4.1.2021.11.9.0

# snmpget 模拟snmp的GetRequest操作的工具。用来获取一个或几个管理信息。用来读取管理信息的内容。
# 获取设备的描述信息
[root@paas ~]# snmpget -c ccssoft -v 2c paas-node1.m8.ccs sysDescr.0
SNMPv2-MIB::sysDescr.0 = STRING: Linux paas-node1.m8.ccs 3.10.0-514.26.2.el7.x86_64 #1 SMP Tue Jul 4 15:04:05 UTC 2017 x86_64
[root@paas ~]# uname -a
Linux paas.m8.ccs 3.10.0-693.5.2.el7.x86_64 #1 SMP Fri Oct 20 20:32:50 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux
[root@paas ~]# 

# 获取磁盘信息
[root@paas ~]# snmpdf -v2c -c ccssoft localhost

```

#### snmp常用总结
[SNMP监控一些常用OID的总结](https://www.cnblogs.com/aspx-net/p/3554044.html)

#### snmpget与snmpwalk区别
```markdown
一、snmpwalk和snmpget的区别：

snmpwalk是对OID值的遍历（比如某个OID值下面有N个节点，则依次遍历出这N个节点的值。如果对某个叶子节点的OID值做walk，则取得到数据就不正确了，因为它会认为该节点是某些节点的父节点，
而对其进行遍历，而实际上该节点已经没有子节点了，那么它会取出与该叶子节点平级的下一个叶子节点的值，而不是当前请求的节子节点的值。）

snmpget是取具体的OID的值。（适用于OID值是一个叶子节点的情况）

```