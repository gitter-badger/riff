package start

import (
	"flag"
	"fmt"
	"strings"
)

const synopsis = "Start Riff service"
const help = `Usage: start [options]

  Start riff service

Options:

  -name       Node name
  -dc         DataCenter name
  -http       Http address of riff (-http 127.0.0.1:8610)
  -dns        Dns address of riff (-dns 127.0.0.1:8620)
  -rpc        RPC address of riff (-rpc [::]:8630)
  -join       Join RPC address (-join 192.168.1.1:8630,192.168.1.2:8630,192.168.1.3:8630)
`

const infoServerPrefix = "[INFO] riff.server: "

type cmd struct {
	flags *flag.FlagSet
	help  string
	// flags
	name string
	dc   string
	http string
	dns  string
	rpc  string
	join string
}

func New() *cmd {
	c := &cmd{}
	c.init()
	return c
}

func (c *cmd) init() {
	c.flags = flag.NewFlagSet("start", flag.ContinueOnError)
	c.flags.StringVar(&c.http, "http", "", "usage")
	c.flags.StringVar(&c.dns, "dns", "", "usage")
	c.flags.StringVar(&c.rpc, "rpc", "", "usage")
	c.flags.StringVar(&c.name, "name", "", "usage")
	c.flags.StringVar(&c.join, "join", "", "usage")
	c.flags.StringVar(&c.dc, "dc", "", "usage")

	c.flags.Usage = func() {
		fmt.Println(c.Help())
	}
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return strings.TrimSpace(help)

}

//func (c *cmd) Ping() {
//	conn, err := net.DialTimeout("tcp", ":8530", time.Second*10)
//	if err != nil {
//		fmt.Println("error", err)
//		return
//	}
//	//encBuf := bufio.NewWriter(conn)
//	var exit = make(chan bool)
//	codec := common.NewGobClientCodec(conn)
//	//codec := jsonrpc.NewClientCodec(conn)
//	cmd := rpc.NewClientWithCodec(codec)
//	var reply string
//	for {
//		err = cmd.Call("Status.Ping", struct{}{}, &reply)
//		if err != nil {
//			fmt.Println("error", err)
//			close(exit)
//			break
//		}
//		fmt.Println(reply)
//		time.Sleep(5 * time.Second)
//	}
//	<-exit
//	//cmd.Close()
//	//if err != nil && errc != nil {
//	//	return fmt.Errorf("%s %s", err, errc)
//	//}
//	//if err != nil {
//	//	return err
//	//} else {
//	//	return errc
//	//}
//}
