
# 使用说明
> 此程序初版是只做cli版本,压缩体积小、跨平台可用，随着功能越多、参数越多使用人记不住参数，增加了等保模式下的Web平台，渗透测试、信息收集使用cli运行。

# 主要功能
> 主机存活探测、漏洞扫描、子域名扫描、端口扫描、各类服务数据库爆破、poc扫描、xss扫描、webtitle探测、web指纹识别、web敏感信息泄露、web目录浏览、web文件下载、等保安全风险问题风险自查等

## web功能预览
![web](images/web.gif)

## WEB目录扫描功能预览
![dirsearch](images/dirsearch.jpg)

## 端口扫描功能预览
![port](images/port.jpg)

## 等保采集配置模式预览
![Windows](images/windows.png)

## 弱口令/未授权现阶段支持类型
| 序号 |          类型           | 是否支持 |         备注         |
|:--:|:---------------------:|:----:|:------------------:|
| 1  |          SSH          |  √   |                    |
| 2  |          RDP          |  √   |                    |
| 3  |          FTP          |  √   |                    |
| 4  |         MySQL         |  √   |                    |
| 5  |      PostgreSQL       |  √   |                    |
| 6  |         Redis         |  √   |                    |
| 7  |         MSSQL         |  √   |                    |
| 8  |          SMB          |  √   |                    |
| 9  |        Telnet         |  √   |                    |
| 10 |        Tomcat         |  √   |                    |
| 11 |        MangoDB        |  √   |      仅验证未授权访问      |
| 12 |     Elasticsearch     |  √   |      仅验证未授权访问      |
| 13 |        oracle         |  √   |                    |
| 14 |       ZooKeeper       |  √   |      仅验证未授权访问      |
| 15 |         dubbo         |  √   |  仅验证默认账户root/root  |
| 16 |          nps          |  √   |  仅验证默认账户admin/123  |
| 17 |         Druid         |  √   |      仅验证未授权访问      |
| 18 |       activemq        |  √   | 仅验证默认账户admin/admin |
| 20 |        couchdb        |  √   |      仅验证未授权访问      |
| 21 | Hadoop-Administration |  √   |      仅验证未授权访问      |
| 22 |     APACHE-Spark      |  √   |      仅验证未授权访问      |
| 23 |        swagger        |  √   |      仅验证未授权访问      |
| 24 |        Kibana         |  √   |      仅验证未授权访问      |
| 25 |     Kafka-Manager     |  √   |      仅验证未授权访问      |


## 端口扫描现阶段支持功能
| 序号 |     功能     | 是否支持 |                                   备注                                   |
|:--:|:----------:|:----:|:----------------------------------------------------------------------:|
| 1  |    多线程     |  √   |                           默认为50并发,可通过-c指定并发数                           |
| 2  |    指定端口    |  √   |                             格式支持1,2,3,2-20                             |
| 3  |    指定IP    |  √   |    格式支持192.168.1.1,192.168.1.1/24,192.168.1-10,http://www.baidu.com    |
| 4  |    排除端口    |  √   |                               格式支持:1,2,3                               |
| 5  | 扫描前探测主机存活  |  √   |                        基于ping,可通过--noping跳过探测存活                        |
| 6  |   打乱主机顺序   |  √   |                         默认不打乱,可通过--random进行打乱                          |
| 7  |    协议识别    |  √   |            目前支持常见协议:ssh、redis、https、https、MySQL、pgsql、ftp等             |
| 8  |    超时时间    |  √   |                              默认3秒,可通过-t指定                              |
| 9  |   识别web    |  √   |                           目前支持识别server、title                           |
| 10 |    结果保存    |  √   |                          默认保存保存到portscan.xlsx                          |
| 11 |  主机操作系统识别  |  √   |                                 基于ttl                                  |
| 12 |    组件识别    |  √   |                                目前常用200+                                |
| 13 |  自动扫描弱口令   |  √   |             rdp、ssh、redis、mysql、oracle、es、telnet、pgsql等20种             |
| 14 | web自动扫描xss |  √   |                                                                        |
| 15 | web自动扫描漏洞  |  √   |                            扫描poc、未授权访问、目录泄露                            |
| 16 |   快速扫描格式   |  √   | 支持格式：https://192.168.1.1:9090、http://192.168.1.1:9090、192.168.1.1:9090 |
| 17 |  sql注入扫描   |  √   |                                                                        |

## web目录扫描现阶段支持功能
| 序号 |      功能       | 是否支持 |       备注        |
|:--:|:-------------:|:----:|:---------------:|
| 1  |      多线程      |  √   |     默认为30并发     |
| 2  |    自定义状态码     |  √   |     默认为200      |
| 3  |     代理模式      |  √   |  http/s、socks   |
| 4  |    返回title    |  √   |                 |
| 5  |    超时等待时常     |  √   |      默认为3秒      |
| 6  |     循环等待      |  √   |     默认为无限制      |
| 7  |     内置url     |  √   |       3W+       |
| 8  | 自定义User-Agent |  √   |                 |
| 9  |      重传       |      |                 |
| 10 |      爬虫       |      |                 |
| 11 |     结果保存      |  √   | 保存到dirScan.json |
| 12 |     内置字典      |  √   |     3W条目录路径     |
| 13 |    识别目录浏览     |  √   |                 |
| 14 |   识别敏感信息泄露    |  √   |                 |
| 15 |    识别文件下载     |  √   |                 |
| 16 |     xss扫描     |  √   |                 |
| 17 |     组件识别      |  √   |    目前常用200+     |

## 自动化测评现阶段支持类型
| 序号 |      类型       | 是否支持 |     备注     |
|:--:|:-------------:|:----:|:----------:|
| 1  |    Centos     |  √   | SSH远程或本地运行 |
| 2  |    Windows    |  √   |    本地运行    |
| 3  |     Redis     |  √   |  远程或本地运行   |
| 4  |  PostgreSQL   |  √   |  远程或本地运行   |
| 5  |    Oracle     |  √   |  远程或本地运行   |
| 6  |     MSSQL     |  √   |  远程或本地运行   |
| 7  |      H3C      |  √   |   SSH远程    |
| 8  |      华为       |  √   |   SSH远程    |
| 9  |      AIX      |  √   | 可自定义命令未内置  |
| 10 |    Ubuntu     |  √   | 可自定义命令未内置  |
| 11 |    MongoDB    |      |            |
| 12 | Elasticsearch |      |            |


## 子域名扫描现阶段支持功能
| 序号 |     功能     | 是否支持 |              备注              |
|:--:|:----------:|:----:|:----------------------------:|
| 1  |    多线程     |  √   |           默认为30并发            |
| 2  |   内置常用域名   |  √   |             100+             |
| 3  |   指定字典文件   |  √   |        可基于字典文件实现多层扫描         |
| 4  |   fofa调用   |  √   | 需在环境变量设置 FOFA_EMAIL、FOFA_KEY |
| 5  | RapidDNS调用 |  √   |                              |


# 常用启动参数
```
golin web (通过web方式启动,仅支持等保功能)
golin port -i 192.168.1.1/24 (扫描c段端口并扫描弱口令、xss、poc漏洞)
golin port -i 192.168.1.1/24 --ipfile ip.txt (扫描指定IP段的同时扫描ip.txt文件中的主机,默认读取ip.txt,目录下如果存在不使用--ipfile也会读取)
golin port -i 192.168.1.1:8080 (快速扫描某一主机的特定端口)
golin port -i 192.168.1.1/24 -c 1000 -t 10(仅扫描c段端口并设置并发数为1000,端口连接超时为10秒)
golin port -i 192.168.1.1/24 --noping --nocrack --random(扫描c段端口但不探测存活不扫描弱口令,并且打乱主机顺序扫描)
golin port -i 192.168.1.1/24 --nopoc(扫描c段端口但禁用扫描漏洞)
golin dirsearch -u https://test.com -f 字典.txt --code 200,404 (扫描状态码为200以及404的web目录)
golin domain -u baidu.com --api (扫描子域名,并且调用fofa、RapidDNS的API)
golin [linux、mysql、oracle、sqlserver、redis、windows...] (按照3级等保要求核查各项安全配置生成html形式报告)
golin update (检查是否可更新)
```

# 免责声明
本工具仅面向合法授权的企业安全建设行为，如您需要测试本工具的可用性，请自行搭建靶机环境。

在使用本工具进行检测时，您应确保该行为符合当地的法律法规，并且已经取得了足够的授权。请勿对非授权目标进行扫描。

如您在使用本工具的过程中存在任何非法行为，您需自行承担相应一切后果。