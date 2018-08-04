const Web3 = require('web3');
const w1 = new Web3('http://localhost:8080/rpc');
const w2 = new Web3('http://localhost:8545');

test('net_version', async () => {
  const id = await w1.eth.net.getId();
  expect(id).toEqual(1);
});

test('eth_accounts', async () => {
  const accounts = await w1.eth.getAccounts();
  const balance = await w1.eth.getBalance(accounts[0]);

  expect(accounts[0]).not.toBeNull();
  expect(balance).not.toBeNull();
});

test('personal_newAccount', async () => {
  const account = await w1.eth.personal.newAccount('');
  const balance = await w1.eth.getBalance(account);

  expect(account).not.toBeNull();
  expect(balance).toEqual('0');
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
