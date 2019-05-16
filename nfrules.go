package nftableslib

import (
	"sync"

	"github.com/google/nftables"
)

// RulesInterface defines third level interface operating with nf Rules
type RulesInterface interface {
	Rules() RuleFuncs
}

// RuleFuncs defines funcations to operate with Rules
type RuleFuncs interface {
	Create()
}

type nfRules struct {
	conn  *nftables.Conn
	table *nftables.Table
	sync.Mutex
}

type nfRule struct {
}

func (nfr *nfRules) Rules() RuleFuncs {
	return nfr
}

func (nfr *nfRules) Create() {

}

func newRules(conn *nftables.Conn, t *nftables.Table) RulesInterface {
	return &nfRules{
		conn:  conn,
		table: t,
	}
}
