const KeepToken = artifacts.require('./KeepToken.sol');

contract('TestToken', function(accounts) {

  let token,
    account_one = accounts[0],
    account_two = accounts[1];

  before(async () => {
    token = await KeepToken.new();
  });

  it("should send tokens correctly", async function() {
    let amount = web3.utils.toBN(1000000000);

    // Starting balances
    let account_one_starting_balance = await token.balanceOf.call(account_one);
    let account_two_starting_balance = await token.balanceOf.call(account_two);

    // Send tokens
    await token.transfer(account_two, amount, {from: account_one});

    // Ending balances
    let account_one_ending_balance = await token.balanceOf.call(account_one);
    let account_two_ending_balance = await token.balanceOf.call(account_two);

    assert.equal(account_one_ending_balance.eq(account_one_starting_balance.sub(amount)), true, "Amount wasn't correctly taken from the sender");
    assert.equal(account_two_ending_balance.eq(account_two_starting_balance.add(amount)), true, "Amount wasn't correctly sent to the receiver");

  });
});
