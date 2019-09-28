// Inspired from https://github.com/arcadeum/multi-token-standard/blob/master/src/tests/utils/contract.ts
// Build ethers.Contract instances from ABI JSON files generated by truffle.
//
// adapted this utility from the handy work by the counterfactual team at:
// https://github.com/counterfactual/monorepo/blob/d9be8524a691c45b6aac1b5e1cf2ff81059203df/packages/contracts/utils/contract.ts

import * as ethers from 'ethers'
import { makeDebug } from './debug'
const debug = makeDebug('abstractContract')

interface NetworkMapping {
  [networkId: number]: { address: string }
}

interface BuildArtifact {
  readonly contractName?: string
  readonly abi: any[]
  readonly bytecode: string
  readonly networks?: NetworkMapping
}

/**
 * Convenience class for an undeployed contract i.e. only the ABI and bytecode.
 */
export class AbstractContract {
  /**
   * Load build artifact by name into an abstract contract
   * @example
   *  const CountingApp = AbstractContract.fromArtifactName("CountingApp", { StaticCall })
   * @param artifactName The name of the artifact to load
   * @param links Optional AbstractContract libraries to link.
   * @returns Truffle artifact wrapped in an AbstractContract.
   */
  public static fromArtifactName(
    artifactName: string,
    links?: Record<string, AbstractContract>,
  ): AbstractContract {
    // these ABI JSON files are generated by truffle
    const contract: BuildArtifact = require(`../../build/contracts/${artifactName}.json`)
    return AbstractContract.fromBuildArtifact(contract, links, artifactName)
  }

  /**
   * Wrap build artifact in abstract contract
   * @param buildArtifact Truffle contract to wrap
   * @param links Optional AbstractContract libraries to link.
   * @returns Truffle artifact wrapped in an AbstractContract.
   */
  public static fromBuildArtifact(
    buildArtifact: BuildArtifact,
    links?: Record<string, AbstractContract>,
    artifactName = 'UntitledContract',
  ): AbstractContract {
    return new AbstractContract(
      buildArtifact.abi,
      buildArtifact.bytecode,
      buildArtifact.networks,
      links,
      artifactName,
    )
  }

  public static async getNetworkID(wallet: ethers.Wallet): Promise<number> {
    return wallet.provider.getNetwork().then(n => n.chainId)
  }

  /**
   * @param abi ABI of the abstract contract
   * @param bytecode Binary of the abstract contract
   * @param networks Network mapping of deployed addresses
   * @param links
   * @param contractName
   */
  constructor(
    readonly abi: any[],
    readonly bytecode: string,
    readonly networks: NetworkMapping = {},
    readonly links?: Record<string, AbstractContract>,
    readonly contractName?: string,
  ) {}

  public toStatic(): AbstractContract {
    debug('abi: %o, %s', this.abi, typeof this.abi)
    this.abi.forEach(m => {
      if (m.constant == null) {
        throw Error(
          'Unknown ABI schema, expected function object with "constant" key',
        )
      }
      m.constant = true
    })
    return this
  }

  /**
   * Get the deployed singleton instance of this abstract contract, if it exists
   * @param Signer (with provider) to use for contract calls
   * @throws Error if AbstractContract has no deployed address
   */
  public async getDeployed(wallet: ethers.Wallet): Promise<ethers.Contract> {
    if (!wallet.provider) {
      throw new Error('Signer requires provider')
    }

    const networkId = await AbstractContract.getNetworkID(wallet)
    const address = this.getDeployedAddress(networkId)
    return new ethers.Contract(address, this.abi, wallet)
  }

  /**
   * Deploy new instance of contract
   * @param wallet Wallet (with provider) to use for contract calls
   * @param args Optional arguments to pass to contract constructor
   * @returns New contract instance
   */
  public async deploy<T>(wallet: ethers.Wallet, args: any[] = []): Promise<T> {
    if (!wallet.provider) {
      throw new Error('Signer requires provider')
    }

    const networkId = await AbstractContract.getNetworkID(wallet)
    const bytecode = this.links
      ? await this.generateLinkedBytecode(networkId)
      : this.bytecode

    const contractFactory = new ethers.ContractFactory(
      this.abi,
      bytecode,
      wallet,
    )

    const contract = await contractFactory.deploy(...args)
    return (contract as any) as T
  }

  /**
   * Connect to a deployed instance of this abstract contract
   * @param signer Signer (with provider) to use for contract calls
   * @param address Address of deployed instance to connect to
   * @returns Contract instance
   */
  public async connect(
    signer: ethers.Signer,
    address: string,
  ): Promise<ethers.Contract> {
    return new ethers.Contract(address, this.abi, signer)
  }

  public getDeployedAddress(networkId: number): string {
    const info = this.networks[networkId]
    if (!info) {
      throw new Error(
        `Abstract contract ${this.contractName} not deployed on network ${networkId}`,
      )
    }

    return info.address
  }

  public async generateLinkedBytecode(networkId: number): Promise<string> {
    if (!this.links) {
      throw new Error('Nothing to link')
    }

    let bytecode = this.bytecode

    for (const name of Object.keys(this.links)) {
      const library = await this.links[name]
      const regex = new RegExp(`__${name}_+`, 'g')
      const address = library.getDeployedAddress(networkId)
      const addressHex = address.replace('0x', '')
      bytecode = bytecode.replace(regex, addressHex)
    }

    return bytecode
  }
}
