package gjkr

import (
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/keep-network/keep-core/pkg/beacon/relay/group"
	"github.com/keep-network/keep-core/pkg/net/ephemeral"
)

func (epkm *EphemeralPublicKeyMessage) RemovePublicKey(
	memberIndex group.MemberIndex,
) {
	delete(epkm.ephemeralPublicKeys, memberIndex)
}

func (mcm *MemberCommitmentsMessage) SetCommitment(
	index int,
	commitment *bn256.G1,
) {
	mcm.commitments[index] = commitment
}

func (mcm *MemberCommitmentsMessage) RemoveCommitment(
	index int,
) {
	mcm.commitments = append(
		mcm.commitments[:index],
		mcm.commitments[index+1:]...,
	)
}

func (psm *PeerSharesMessage) SetShares(
	memberIndex group.MemberIndex,
	encryptedShareS, encryptedShareT []byte,
) {
	psm.shares[memberIndex] = &peerShares{
		encryptedShareS: encryptedShareS,
		encryptedShareT: encryptedShareT,
	}
}

func (psm *PeerSharesMessage) RemoveShares(memberIndex group.MemberIndex) {
	delete(psm.shares, memberIndex)
}

func (ssam *SecretSharesAccusationsMessage) SetAccusedMemberKey(
	memberIndex group.MemberIndex,
	privateKey *ephemeral.PrivateKey,
) {
	ssam.accusedMembersKeys[memberIndex] = privateKey
}

func (mpkspm *MemberPublicKeySharePointsMessage) SetPublicKeyShare(
	index int,
	publicKeyShare *bn256.G2,
) {
	mpkspm.publicKeySharePoints[index] = publicKeyShare
}

func (pam *PointsAccusationsMessage) SetAccusedMemberKey(
	memberIndex group.MemberIndex,
	privateKey *ephemeral.PrivateKey,
) {
	pam.accusedMembersKeys[memberIndex] = privateKey
}

func (dekm *DisqualifiedEphemeralKeysMessage) SetPrivateKey(
	memberIndex group.MemberIndex,
	privateKey *ephemeral.PrivateKey,
) {
	dekm.privateKeys[memberIndex] = privateKey
}
