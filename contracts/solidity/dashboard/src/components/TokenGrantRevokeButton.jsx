import React, { Component } from 'react'
import { Button } from 'react-bootstrap'
import WithWeb3Context from './WithWeb3Context'

class TokenGrantRevokeButton extends Component {

  revoke = async () => {
    const { web3, item } = this.props
    await web3.grantContract.methods.revoke(item.id).send({from: web3.yourAddress})
  }

  render() {
    const { item } = this.props
    let button = 'Non revocable'

    if (item.revoked) {
      button = 'Revoked'
    }

    if (item.revocable && !item.revoked) {
      button = <Button bsSize="small" bsStyle="primary" onClick={this.revoke}>Revoke</Button>
    }

    return button
  }
}

export default WithWeb3Context(TokenGrantRevokeButton)
