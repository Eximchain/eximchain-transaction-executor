
const Chai = require('chai')
assert = Chai.assert
const Web3 = require('web3')
const BigNumber = require('bignumber.js')
const Request = require('request-promise')

//
// CONFIG
//
const web3Url = 'http://localhost:8080/'


describe('Web3 RPC Tests', () => {

   var web3 = null

   var receipt  = null
   var accounts = null


   before(async () => {
      web3 = new Web3(web3Url)

      accounts = await web3.eth.getAccounts()
   })


   it('eth_sendTransaction', async () => {
      receipt = await web3.eth.sendTransaction({ from: accounts[0], to: accounts[1], gas: 100000, value: 1 })
      assert.typeOf(receipt, 'object')
      assert.propertyVal(receipt, 'status', true)
   })

   it('eth_sendTransaction (fail)', async () => {
      try {
         await web3.eth.sendTransaction({ from: accounts[0], gas: 100000, value: 1 })
         assert.fail('sendTransaction should fail if no "to" defined.')
      } catch (e) {}
   })

   it('personal_newAccount', async () => {
      const a = await web3.eth.personal.newAccount('')
      assert.typeOf(a, 'string')
      assert.equal(a.length, 42)
   })

   it('web3_clientVersion', async () => {
      const o = await web3.eth.getNodeInfo()
      assert.typeOf(o, 'string')
      assert.isAtLeast(o.length, 1)
   })

   it('eth_blockNumber', async () => {
      const o = await web3.eth.getBlockNumber()
      assert.typeOf(o, 'number')
      assert.isAtLeast(o, 0)
   })

   it('web3_sha3', async () => {
      // Note: Web3 doesn't forward sha3 as RPC calls so we'll use Request instead.
      //
      const options = {
         method : 'POST',
         uri : web3Url,
         body : {
            "jsonrpc" : "2.0",
            "method" : "web3_sha3",
            "params" : [ web3.utils.utf8ToHex("EXIMCHAIN") ],
            "id" : 1
         },
         json : true
      }

      const o = await Request(options)
      assert.typeOf(o, 'object')
      assert.propertyVal(o, 'result', '0xf8747e43def95db9f64ce2760690464e98a4e41fa6a0d8405d818d796d19f2d9')
   })

   it('net_version', async () => {
      const o = await web3.eth.net.getId()
      assert.typeOf(o, 'number')
      assert.isAtLeast(o, 1)
   })

   it('net_peerCount', async () => {
      const o = await web3.eth.net.getPeerCount()
      assert.typeOf(o, 'number')
      assert.isAtLeast(o, 0)
   })

   it('net_listening', async () => {
      const o = await web3.eth.net.isListening()
      assert.equal(o, true)
   })

   it('eth_protocolVersion', async () => {
      const o = await web3.eth.getProtocolVersion()
      assert.typeOf(o, 'string')
      assert.isTrue(o.startsWith('0x'))
   })

   it('eth_syncing', async () => {
      const o = await web3.eth.isSyncing()
      assert.equal(o, false)
   })

   it('eth_coinbase', async () => {
      const o = await web3.eth.getCoinbase()
      assert.typeOf(o, 'string')
      assert.equal(o.length, 42)
   })

   it('eth_mining', async () => {
      const o = await web3.eth.isMining()
      assert.typeOf(o, 'boolean')
   })

   it('eth_hashrate', async () => {
      const o = await web3.eth.getHashrate()
      assert.typeOf(o, 'number')
      assert.isAtLeast(o, 0)
   })

   it('eth_gasPrice', async () => {
      const o = await web3.eth.getGasPrice()
      assert.typeOf(o, 'string')
      assert.isAtLeast(parseInt(o), 0)
   })

   it('eth_blockNumber', async () => {
      const o = await web3.eth.getBlockNumber()
      assert.typeOf(o, 'number')
      assert.isAtLeast(o, 0)
   })

   it('eth_accounts', async () => {
      const o = await web3.eth.getAccounts()
      assert.isArray(o)
      assert.isAtLeast(o.length, 1)
   })

   it('eth_getBalance', async () => {
      const o = new BigNumber(await web3.eth.getBalance(accounts[0]))
      assert.isTrue(o.gt(0))
   })

   it('eth_getBalance (fail)', async () => {
      try {
         await web3.eth.getBalance('0x')
         assert.fail('getBalance should fail with invalid address')
      } catch (e) {}
   })

   it('eth_getStorageAt', async () => {
      const o = await web3.eth.getStorageAt(accounts[0])
      assert.typeOf(o, 'string')
      assert.isTrue(o.startsWith('0x'))
      assert.isAtLeast(o.length, 4)
   })

   it('eth_getTransactionCount', async () => {
      const o = await web3.eth.getTransactionCount(accounts[0])
      assert.typeOf(o, 'number')
      assert.isAtLeast(o, 0)
   })

   it('eth_getBlockTransactionCountByHash', async () => {
      const block = await web3.eth.getBlock('latest')
      const o = await web3.eth.getBlockTransactionCount(block.hash)
      assert.typeOf(o, 'number')
      assert.isAtLeast(o, 0)
   })

   it('eth_getBlockTransactionCountByNumber', async () => {
      const block = await web3.eth.getBlock('latest')
      const o = await web3.eth.getBlockTransactionCount(block.number)
      assert.typeOf(o, 'number')
      assert.isAtLeast(o, 0)
   })

   it('eth_getUncleCountByBlockHash', async () => {
      const o = await web3.eth.getBlockUncleCount(receipt.blockHash)
      assert.typeOf(o, 'number')
      assert.isAtLeast(o, 0)
   })

   it('eth_getUncleCountByBlockNumber', async () => {
      const o = await web3.eth.getBlockUncleCount(receipt.blockNumber)
      assert.typeOf(o, 'number')
      assert.isAtLeast(o, 0)
   })

   it('eth_getCode', async () => {
      const o = await web3.eth.getCode(accounts[0])
      assert.typeOf(o, 'string')
      assert.isTrue(o.startsWith('0x'))
   })

   it('eth_sign', async () => {
      const o = await web3.eth.sign('EXIMCHAIN', accounts[0]);
      assert.typeOf(o, 'string')
      assert.isAtLeast(o.length, 20)
   })

   it('eth_sendRawTransaction', async () => {
      var rawTx = {
         from   : accounts[0],
         to     : accounts[1],
         gas    : 1000000,
         value  : 1,
         data   : ""
      }

      const signed = await web3.eth.signTransaction(rawTx, accounts[3])

      const o = await web3.eth.sendSignedTransaction(signed)
      assert.typeOf(o, 'object')
      assert.propertyVal(o, 'status', true)
   })

   it('eth_call', async () => {
      const o = await web3.eth.call({})
      assert.typeOf(o, 'string')
      assert.equal(o, '0x')
   })

   it('eth_estimageGas', async () => {
      const o = await web3.eth.estimateGas({})
      assert.typeOf(o, 'number')
      assert.isAtLeast(o, 21000)
   })

   it('eth_getBlockByNumber', async () => {
      const o = await web3.eth.getBlock(0)
      assert.typeOf(o, 'object')
      assert.propertyVal(o, 'number', 0)
   })

   it('eth_getBlockByHash', async () => {
      var  o = await web3.eth.getBlock('latest')
      o = await web3.eth.getBlock(o.hash)
      assert.typeOf(o, 'object')
      assert.isAtLeast(o.number, 0)
   })

   it('eth_getTransactionByHash', async () => {
      const o = await web3.eth.getTransaction(receipt.transactionHash)
      assert.typeOf(o, 'object')
      assert.propertyVal(o, 'blockNumber', receipt.blockNumber)
   })

   it('eth_getTransactionByBlockHashAndIndex', async () => {
      const o = await web3.eth.getTransactionFromBlock(receipt.blockHash, receipt.transactionIndex)
      assert.typeOf(o, 'object')
      assert.propertyVal(o, 'blockNumber', receipt.blockNumber)
   })

   it('eth_getTransactionByBlockNumberAndIndex', async () => {
      const o = await web3.eth.getTransactionFromBlock(receipt.blockNumber, receipt.transactionIndex)
      assert.typeOf(o, 'object')
      assert.propertyVal(o, 'blockNumber', receipt.blockNumber)
   })

   it('eth_getTransactionReceipt', async () => {
      const o = await web3.eth.getTransactionReceipt(receipt.transactionHash)
      assert.typeOf(o, 'object')
      assert.propertyVal(o, 'transactionHash', receipt.transactionHash)
   })

   it('eth_getUncleByBlockHashAndIndex', async () => {
      const o = await web3.eth.getUncle(receipt.blockHash, 0)
      assert.isNull(o)
   })

   it('eth_getUncleByBlockNumberAndIndex', async () => {
      const o = await web3.eth.getUncle(receipt.blockNumber, 0)
      assert.isNull(o)
   })

   it('eth_getLogs', async () => {
      const o = await web3.eth.getPastLogs({})
      assert.isArray(o)
      assert.isAtLeast(o.length, 0)
   })

   it('eth_getWork', async () => {
      const o = await web3.eth.getWork()
      assert.isArray(o)
      assert.equal(o.length, 3)
   })


   /*
   m["eth_submitWork"]
   m["eth_newFilter"]
   m["eth_newBlockFilter"]
   m["eth_newPendingTransactionFilter"]
   m["eth_uninstallFilter"]
   m["eth_getFilterChanges"]
   m["eth_getFilterLogs"]
   m["eth_submitHashrate"]
   */
})

