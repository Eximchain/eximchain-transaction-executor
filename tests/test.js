const Web3 = require('web3');
const w1 = new Web3('http://localhost:8080/rpc');
const w2 = new Web3('http://localhost:8545');

test('net_version', async () => {
  const id = await w1.eth.net.getId();
  expect(id).toEqual(1);
});

test('eth_accounts', async () => {
  const accounts = await w1.eth.getAccounts();
  const b1 = await w1.eth.getBalance(accounts[0]);
  const b2 = await w2.eth.getBalance(accounts[0]);

  expect(accounts[0]).not.toBeNull();
  expect(b1).not.toBeNull();
  expect(b1).toEqual(b2);
});

test('personal_newAccount', async () => {
  const a = await w1.eth.personal.newAccount('');
  const b = await w1.eth.getBalance(a);

  expect(a).not.toBeNull();
  expect(b).toEqual('0');
});

test('eth_sign', async () => {
  const accounts = await w1.eth.getAccounts();
  const account = accounts[0];

  const s1 = await w1.eth.sign('Hello World', account);
  const s2 = await w2.eth.sign('Hello World', account);
  const s3 = await w2.eth.personal.sign('Hello World', account, '');

  expect(s1).toEqual(s2);
  expect(s1).toEqual(s3);
});

test('eth_sendTransaction', async () => {
  const accounts = await w1.eth.getAccounts();
  const a1 = accounts[0];
  const a2 = await w1.eth.personal.newAccount('');

  const b2 = await w1.eth.getBalance(a2);

  const receipt = await w1.eth.sendTransaction({
    from: a1,
    to: a2,
    gas: 1000000,
    value: '1000'
  });

  const bb2 = await w1.eth.getBalance(a2);

  expect(bb2 - b2).toEqual(1000);
  expect(receipt).not.toBeNull();
});
