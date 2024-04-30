package casper

import (
	"github.com/gin-gonic/gin"
)

func (cn *Casper) CheckChainMainnet() gin.HandlerFunc {
	return func(c *gin.Context) {
		//check if set get param rpcurl
		rpcurl := c.Query("rpc_node")
		if IsNullRpcNode(rpcurl) {
			c.Next()
			return
		}
		rpcurl, err := IpToUrl(rpcurl)
		if err != nil {
			c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
			c.Abort()
			return
		}
		ok, err := cn.client.CheckChain(rpcurl, "casper")
		if err != nil {
			c.JSON(500, ErrorResponse{"error querying casper network - check chain"})
			c.Abort()
			return
		}
		if !ok {
			c.JSON(400, gin.H{"error": "Wrong chain id for this node"})
			c.Abort()
			return
		}
		c.Set("chain", "casper")
		c.Next()
	}
}
func (cn *Casper) CheckChainTestnet() gin.HandlerFunc {
	return func(c *gin.Context) {
		//check if set get param rpcurl
		rpcurl := c.Query("rpc_node")
		if IsNullRpcNode(rpcurl) {
			c.Next()
			return
		}
		rpcurl, err := IpToUrl(rpcurl)
		if err != nil {
			c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
			c.Abort()
			return
		}
		ok, err := cn.client.CheckChain(rpcurl, "casper-test")
		if err != nil {
			c.JSON(500, ErrorResponse{"error querying casper network - check chain"})
			c.Abort()
			return
		}
		if !ok {
			c.JSON(400, gin.H{"error": "Wrong chain id for this node"})
			c.Abort()
			return
		}
		c.Set("chain", "casper-test")
		c.Next()
	}
}

func IsNullRpcNode(rpcurl string) bool {
	return (rpcurl == "0.0.0.0" || rpcurl == "")
}

func (cn *Casper) CheckAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Query("address")
		if address == "" {
			c.Next()
			return
		}
		ok := cn.client.IsAddress(address)
		if !ok {
			c.JSON(400, gin.H{"error": "Invalid address"})
			c.Abort()
			return
		}
		c.Next()
	}
}
