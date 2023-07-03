package cmd

import (
	"github.com/spf13/cobra"
	"golin/crack"
)

var expCmd = &cobra.Command{
	Use:   "crack",
	Short: "弱口令检测功能,目前支持：ssh、mysql、redis、postgresql、sqlserver、ftp",
}

// 破解 ssh
var ssh = &cobra.Command{
	Use:   "ssh",
	Short: "ssh弱口令检测",
	Run:   crack.Run,
}

// 破解 mysql
var mysql = &cobra.Command{
	Use:   "mysql",
	Short: "mysql弱口令检测",
	Run:   crack.Run,
}

// 破解 redis
var redis = &cobra.Command{
	Use:   "redis",
	Short: "redis弱口令检测",
	Run:   crack.Run,
}

// 破解 pgsql
var pgsql = &cobra.Command{
	Use:   "pgsql",
	Short: "pgsql弱口令检测",
	Run:   crack.Run,
}

// 破解 sqlserver
var sqlserver = &cobra.Command{
	Use:   "sqlserver",
	Short: "sqlserver弱口令检测",
	Run:   crack.Run,
}

// 破解 ftp
var ftp = &cobra.Command{
	Use:   "ftp",
	Short: "ftp弱口令检测",
	Run:   crack.Run,
}

// 破解 smb
var smb = &cobra.Command{
	Use:   "smb",
	Short: "smb弱口令检测",
	Run:   crack.Run,
}

// 破解 telnet
var telnet = &cobra.Command{
	Use:   "telnet",
	Short: "telnet弱口令检测",
	Run:   crack.Run,
}

// 破解 tomcat
var tomcat = &cobra.Command{
	Use:   "tomcat",
	Short: "tomcat弱口令检测",
	Run:   crack.Run,
}

// 破解 xlsx
var xlsx = &cobra.Command{
	Use:   "xlsx",
	Short: "xlsx口令检测",
	Run:   crack.XlsxCheck,
}

func init() {
	rootCmd.AddCommand(expCmd)
	commands := []*cobra.Command{ssh, mysql, redis, pgsql, sqlserver, ftp, smb, telnet, tomcat}
	for _, cmd := range commands {
		expCmd.AddCommand(cmd)
		cmd.Flags().StringP("ip", "i", "", "此参数是指定验证的IP 支持格式如下：192.168.0.1，或192.168.0.1/24，或192.168.0.1:22")
		cmd.Flags().IntP("port", "P", 0, "此参数是指定验证的端口")
		cmd.Flags().StringP("user", "u", "", "此参数是指定用户文件")
		cmd.Flags().StringP("passwd", "p", "", "此参数是指定密码文件")
		cmd.Flags().StringP("fire", "f", "", "此参数是指定主机列表，格式IP:Port 一行一个")
		cmd.Flags().IntP("chan", "c", 30, "并发数量")
		cmd.Flags().IntP("time", "t", 3, "超时等待时常/s")
		cmd.Flags().Bool("noping", false, "此参数是指定不运行网络监测")
	}
	// xlsx
	expCmd.AddCommand(xlsx)
	xlsx.Flags().StringP("file", "f", "", "此参数是指定待破解的文件路径 注：xlsx 格式")
	xlsx.Flags().StringP("passwdfile", "p", "", "此参数是指定待尝试的密码文件")
	xlsx.Flags().IntP("chan", "c", 30, "并发数量")

}
