package cmd

import (
	"fmt"
	"github.com/go-ping/ping"
	"github.com/spf13/cobra"
	"net"
	"os"
	"time"
)

// testCmd represents the domain
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test all ip after domain by ping ",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Fprintf(os.Stderr, "Your target domian is: %s \nThe ips bebind are : \n", domainName)

		/*		//使用dig命令进行解析
				command := exec.Command("dig", " +nocmd ", domainName, "+noall", "+answer")
				out, err := command.CombinedOutput()
				if err != nil {
					log.Fatalf("cmd.Run() failed with %s\n", err)
				}
				fmt.Printf("combined out:\n%s\n", string(out))*/

		iprecords, _ := net.LookupIP(domainName)
		for _, ip := range iprecords {
			fmt.Println(ip)
			ip.String()
		}
		fmt.Println("The Delay for ips by ping : ")
		for _, ip := range iprecords {
			checkServer(ip.String())
		}
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	// 解析参数
	testCmd.PersistentFlags().StringVarP(&domainName, "domain", "d", "", "域名")

}

func checkServer(ip string) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		panic(err)
	}
	pinger.Timeout = time.Duration(time.Millisecond * 1000)
	pinger.Count = 3
	pinger.TTL = 64
	pinger.Interval = time.Duration(time.Millisecond * 100)
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	fmt.Println(stats)
}
