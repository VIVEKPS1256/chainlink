import {
  assertActionThrows,
  bigNum,
  calculateSAID,
  checkPublicABI,
  deploy,
  newAddress,
  newHash,
  oracleNode,
  personalSign,
  recoverPersonalSignature,
  stranger,
  toHex
} from './support/helpers'
import { assertBigNum } from './support/matchers'

contract('Coordinator', () => {
  const sourcePath = 'Coordinator.sol'
  let coordinator

  beforeEach(async () => {
    coordinator = await deploy(sourcePath)
  })

  it('has a limited public interface', () => {
    checkPublicABI(artifacts.require(sourcePath), [
      'getId',
      'initiateServiceAgreement',
      'serviceAgreements'
    ])
  })

  describe('#getId', () => {
    it('matches the ID generated by the oracle off-chain', async () => {
      let result = await coordinator.getId.call(
        1,
        2,
        ['0x70AEc4B9CFFA7b55C0711b82DD719049d615E21d', '0xd26114cd6EE289AccF82350c8d8487fedB8A0C07'],
        '0x85820c5ec619a1f517ee6cfeff545ec0ca1a90206e1a38c47f016d4137e801dd'
      )
      assert.equal(result, '0x2249a9e0a0463724551b2980299a16406da6f4e80d911628ee77445be4e04e7c')
    })
  })

  describe('#initiateServiceAgreement', () => {
    let payment, expiration, oracle, oracles, requestDigest,
      serviceAgreementID, oracleSignature

    beforeEach(async () => {
      payment = newHash('1000000000000000000')
      expiration = newHash('300')
      oracle = newAddress(oracleNode)
      oracles = [oracle]
      requestDigest = newHash('0x9ebed6ae16d275059bf4de0e01482b0eca7ffc0ffcc1918db61e17ac0f7dedc8')

      serviceAgreementID = calculateSAID(payment, expiration, oracles, requestDigest)
    })

    context("with valid oracle signatures", () => {
      beforeEach(async () => {
        oracleSignature = personalSign(oracle, serviceAgreementID)
        const requestDigestAddr = recoverPersonalSignature(serviceAgreementID, oracleSignature)
        assert.equal(toHex(oracle), toHex(requestDigestAddr))
      })

      it('saves a service agreement struct from the parameters', async () => {
        await coordinator.initiateServiceAgreement(
          toHex(payment),
          toHex(expiration),
          oracles.map(toHex),
          [oracleSignature.v],
          [oracleSignature.r].map(toHex),
          [oracleSignature.s].map(toHex),
          toHex(requestDigest)
        )

        const sa = await coordinator.serviceAgreements.call(toHex(serviceAgreementID))

        assertBigNum(sa[0], bigNum(toHex(payment)))
        assertBigNum(sa[1], bigNum(toHex(expiration)))
        assert.equal(sa[2], toHex(requestDigest))

        /// / TODO:
        /// / Web3.js doesn't support generating an artifact for arrays within a struct.
        /// / This means that we aren't returned the list of oracles and
        /// / can't assert on their values.
        /// /
        /// / However, we can pass them into the function to generate the ID
        /// / & solidity won't compile unless we pass the correct number and
        /// / type of params when initializing the ServiceAgreement struct,
        /// / so we have some indirect test coverage.
        /// /
        /// / https://github.com/ethereum/web3.js/issues/1241
        /// / assert.equal(
        /// /   sa[2],
        /// /   ['0x70AEc4B9CFFA7b55C0711b82DD719049d615E21d', '0xd26114cd6EE289AccF82350c8d8487fedB8A0C07']
        /// / )
      })
    })

    context("with an invalid oracle signatures", () => {
      beforeEach(async () => {
        oracleSignature = personalSign(newAddress(stranger), serviceAgreementID)
      })

      it('saves a service agreement struct from the parameters', async () => {
        assertActionThrows(async () => {
          await coordinator.initiateServiceAgreement(
            toHex(payment),
            toHex(expiration),
            oracles.map(toHex),
            [oracleSignature.v],
            [oracleSignature.r].map(toHex),
            [oracleSignature.s].map(toHex),
            toHex(requestDigest)
          )
        })

        const sa = await coordinator.serviceAgreements.call(toHex(serviceAgreementID))
        assertBigNum(sa[0], bigNum(0))
        assertBigNum(sa[1], bigNum(0))
        assert.equal(sa[2], '0x0000000000000000000000000000000000000000000000000000000000000000')
      })
    })
  })
})
