require('babel-register');
require('babel-polyfill');
const HDWalletProvider = require("@truffle/hdwallet-provider");

module.exports = {
  networks: {
    local: {
      host: "localhost",
      port: 8545,
      network_id: "*"
    },
    keep_dev: {
      provider: function() {
        return new HDWalletProvider(process.env.CONTRACT_OWNER_ETH_ACCOUNT_PRIVATE_KEY, "http://localhost:8545")
      },
      gas: 6721975,
      network_id: 1101
    },

    keep_dev_vpn: {
      provider: function() {
        return new HDWalletProvider(process.env.CONTRACT_OWNER_ETH_ACCOUNT_PRIVATE_KEY, "http://eth-tx-node.default.svc.cluster.local:8545")
      },
      gas: 6721975,
      network_id: 1101
    },

    ropsten: {
      provider: function() {
        return new HDWalletProvider(process.env.CONTRACT_OWNER_ETH_ACCOUNT_PRIVATE_KEY, "https://ropsten.infura.io/v3/59fb36a36fa4474b890c13dd30038be5")
      },
      gas: 6721975,
      network_id: 3
    },

    // TODO: update Infura url
    mainnet: {
      provider: function() {
        return new HDWalletProvider(process.env.CONTRACT_OWNER_ETH_ACCOUNT_PRIVATE_KEY, "https://mainnet.infura.io/v3/")
      },
      network_id: 1
    }
  },

  mocha: {
    useColors: true,
    reporter: 'eth-gas-reporter',
    reporterOptions: {
      currency: 'USD',
      gasPrice: 21,
      showTimeSpent: true
    }
  },

  compilers: {
    solc: {
      version: "0.5.17",
      optimizer: {
        enabled: true,
        runs: 200
      }
    }
  }
};
