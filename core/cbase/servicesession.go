package cbase

import (
	"github.com/zmloong/lotus/core"
	"github.com/zmloong/lotus/sys/registry"
	"github.com/zmloong/lotus/sys/rpc"
)

func NewServiceSession(node *registry.ServiceNode) (ss core.IServiceSession, err error) {
	session := new(ServiceSession)
	session.node = node
	session.rpc, err = rpc.NewRpcClient(node.Id, node.RpcId)
	ss = session
	return
}

type ServiceSession struct {
	node *registry.ServiceNode
	rpc  rpc.IRpcClient
}

func (this *ServiceSession) GetId() string {
	return this.node.Id
}
func (this *ServiceSession) GetIp() string {
	return this.node.IP
}
func (this *ServiceSession) GetRpcId() string {
	return this.node.RpcId
}

func (this *ServiceSession) GetType() string {
	return this.node.Type
}
func (this *ServiceSession) GetVersion() string {
	return this.node.Version
}
func (this *ServiceSession) SetVersion(v string) {
	this.node.Version = v
}

func (this *ServiceSession) GetPreWeight() float64 {
	return this.node.PreWeight
}
func (this *ServiceSession) SetPreWeight(p float64) {
	this.node.PreWeight = p
}
func (this *ServiceSession) Done() {
	this.rpc.Stop()
}
func (this *ServiceSession) Call(f core.Rpc_Key, params ...interface{}) (interface{}, error) {
	return this.rpc.Call(string(f), params...)
}
func (this *ServiceSession) CallNR(f core.Rpc_Key, params ...interface{}) (err error) {
	return this.rpc.CallNR(string(f), params...)
}
