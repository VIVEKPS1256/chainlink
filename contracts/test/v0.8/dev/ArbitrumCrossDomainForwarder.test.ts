import { ethers } from 'hardhat'
import { assert, expect } from 'chai'
import { Contract, ContractFactory } from 'ethers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import {
  impersonateAs,
  publicAbi,
  toArbitrumL2AliasAddress,
} from '../../test-helpers/helpers'

let owner: SignerWithAddress
let stranger: SignerWithAddress
let l1OwnerAddress: string
let crossdomainMessenger: SignerWithAddress
let forwarderFactory: ContractFactory
let greeterFactory: ContractFactory
// let multisendFactory: ContractFactory
let forwarder: Contract
let greeter: Contract
// let multisend: Contract

before(async () => {
  const accounts = await ethers.getSigners()
  owner = accounts[0]
  stranger = accounts[1]

  // Forwarder config
  l1OwnerAddress = owner.address
  crossdomainMessenger = await impersonateAs(
    toArbitrumL2AliasAddress(l1OwnerAddress),
  )

  // Contract factories
  forwarderFactory = await ethers.getContractFactory(
    'src/v0.8/dev/ArbitrumCrossDomainForwarder.sol:ArbitrumCrossDomainForwarder',
    owner,
  )
  greeterFactory = await ethers.getContractFactory(
    'src/v0.8/tests/Greeter.sol:Greeter',
    owner,
  )
  // multisendFactory = await ethers.getContractFactory(
  //   'src/v0.8/tests/Multisend.sol:Multisend',
  //   owner,
  // )
})

describe('ArbitrumCrossDomainForwarder', () => {
  beforeEach(async () => {
    forwarder = await forwarderFactory.deploy(l1OwnerAddress)
    greeter = await greeterFactory.deploy()
    // multisend = await multisendFactory.deploy()
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(forwarder, [
      'typeAndVersion',
      'crossDomainMessenger',
      'forward',
      'forwardDelegate',
      'l1Owner',
      'transferL1Ownership',
      // ConfirmedOwner methods:
      'acceptOwnership',
      'owner',
      'transferOwnership',
    ])
  })

  describe('#constructor', () => {
    it('should set the owner correctly', async () => {
      const response = await forwarder.owner()
      assert.equal(response, owner.address)
    })

    it('should set the l1Owner correctly', async () => {
      const response = await forwarder.l1Owner()
      assert.equal(response, l1OwnerAddress)
    })

    it('should set the crossdomain messenger correctly', async () => {
      const response = await forwarder.crossDomainMessenger()
      assert.equal(response, crossdomainMessenger.address)
    })
  })

  describe('#forward', () => {
    it('should not be callable by unknown crossdomain messenger address', async () => {
      await expect(
        forwarder.connect(stranger).forward(greeter.address, '0x'),
      ).to.be.revertedWith('Sender is not the L2 messenger')
    })
  })

  //   TODO: test forward()
  //   TODO: test forwardDelegate()

  //   describe('#raiseFlag', () => {
  //     describe('when called by the owner', () => {
  //       it('updates the warning flag', async () => {
  //         assert.equal(false, await flags.getFlag(consumer.address))

  //         await flags.connect(personas.Nelly).raiseFlag(consumer.address)

  //         assert.equal(true, await flags.getFlag(consumer.address))
  //       })

  //       it('emits an event log', async () => {
  //         await expect(flags.connect(personas.Nelly).raiseFlag(consumer.address))
  //           .to.emit(flags, 'FlagRaised')
  //           .withArgs(consumer.address)
  //       })

  //       describe('if a flag has already been raised', () => {
  //         beforeEach(async () => {
  //           await flags.connect(personas.Nelly).raiseFlag(consumer.address)
  //         })

  //         it('emits an event log', async () => {
  //           const tx = await flags
  //             .connect(personas.Nelly)
  //             .raiseFlag(consumer.address)
  //           const receipt = await tx.wait()
  //           assert.equal(0, receipt.events?.length)
  //         })
  //       })
  //     })

  //     describe('when called by a non-enabled setter', () => {
  //       it('reverts', async () => {
  //         await expect(
  //           flags.connect(personas.Neil).raiseFlag(consumer.address),
  //         ).to.be.revertedWith('Not allowed to raise flags')
  //       })
  //     })
})
