const Web3 = require('web3');
const web3 = new Web3('http://localhost:8080/rpc');

test('netVersion', async () => {
  const id = await web3.eth.net.getId();
  expect(id).toEqual(1);
});
