package main

import (
	"os"

	"gitlab.10101111.com/oped/uparse.git/define"
	"gitlab.10101111.com/oped/uparse.git/lib/rpcclient"
	"gitlab.10101111.com/oped/uparse.git/pipeline"
	"gitlab.10101111.com/oped/uparse.git/rpcproto"

	log "github.com/Sirupsen/logrus"
)

func main() {

	var rconn rpcclient.RpcClient
	if len(os.Args) != 0 {
		rconn = rpcclient.RpcClient{Seraddr: os.Args[1]}
	} else {
		rconn = rpcclient.RpcClient{Seraddr: "10.212.1.147:2734"}
	}
	if err := rconn.Init(); err != nil {
		log.Errorf("open rpc err:%v", err)
		return
	}

	req := &rpcproto.PerfData{Uparseid: os.Args[0]}
	reply := &rpcproto.PerfDataReply{PerfDatas: make(map[string]pipeline.PluginPerfData)}
	err := rconn.CallRpc(define.RPC_CALL_PERFDATA, req, reply)
	if err != nil {
		log.Errorf("call rpc err:%v", err)
		return
	}

	log.Infof("%#v", reply.PerfDatas)

}
