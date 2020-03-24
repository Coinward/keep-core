const OperatorParamsStub = artifacts.require('./stubs/OperatorParamsStub.sol');

contract('OperatorParamsStub', (accounts) => {
  let opUtils;
  const eighteen = web3.utils.toBN(18)
  const ten = web3.utils.toBN(10)
  const keepDecimals = ten.pow(eighteen);
  const billion = web3.utils.toBN(1000000000)
  const allKeepEver = billion.mul(keepDecimals);

  const blocksPerYear = web3.utils.toBN(3153600);
  const recently = blocksPerYear.muln(5);
  const billionYearsFromNow = blocksPerYear.mul(billion);

  before(async () => {
      opUtils = await OperatorParamsStub.new();
  });

  describe("pack", async () => {
    it("should roundtrip values", async () => {
      const params = await opUtils.publicPack(
        allKeepEver,
        recently,
        billionYearsFromNow);
      const amount = await opUtils.publicGetAmount(params);
      const createdAt = await opUtils.publicGetCreationBlock(params);
      const undelegatedAt = await opUtils.publicGetUndelegationBlock(params);

      assert.equal(
        amount.toJSON(),
        allKeepEver.toJSON(),
        "The amount should be the same");
      assert.equal(
        createdAt.toJSON(),
        recently.toJSON(),
        "The creation block should be the same");
      assert.equal(
        undelegatedAt.toJSON(),
        billionYearsFromNow.toJSON(),
        "The undelegation block should be the same");
    })
  })

  describe("setAmount", async () => {
    it("should set the amount", async () => {
      const params = await opUtils.publicPack(allKeepEver, recently, 0);
      const newParams = await opUtils.publicSetAmount(params, billion);
      const amount = await opUtils.publicGetAmount(newParams);
      assert.equal(
        amount.toJSON(),
        billion.toJSON(),
        "The amount should be the same");
    })
  })

  describe("setCreationBlock", async () => {
    it("should set the creation block", async () => {
      const params = await opUtils.publicPack(allKeepEver, recently, 0);
      const newParams = await opUtils.publicSetCreationBlock(
        params,
        billionYearsFromNow);
      const creationBlock = await opUtils.publicGetCreationBlock(newParams);
      assert.equal(
        creationBlock.toJSON(),
        billionYearsFromNow.toJSON(),
        "The creation block should be the same");
    })
  })

  describe("setUndelegationBlock", async () => {
    it("should set the undelegation block", async () => {
      const params = await opUtils.publicPack(allKeepEver, recently, 0);
      const newParams = await opUtils.publicSetUndelegationBlock(
        params,
        recently);
      const undelegationBlock = await opUtils.publicGetUndelegationBlock(newParams);
      assert.equal(
        undelegationBlock.toJSON(),
        recently.toJSON(),
        "The undelegationBlock should be the same");
    })
  })
})
