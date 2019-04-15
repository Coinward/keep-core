package result

import (
	"github.com/keep-network/keep-core/pkg/beacon/relay/chain"
	"github.com/keep-network/keep-core/pkg/beacon/relay/dkg/result/gen/pb"
	"github.com/keep-network/keep-core/pkg/beacon/relay/group"
	"github.com/keep-network/keep-core/pkg/operator"
)

// Type returns a string describing a DKGResultHashSignatureMessage type for
// marshalling purposes.
func (d *DKGResultHashSignatureMessage) Type() string {
	return "result/dkg_result_hash_signature_message"
}

// Marshal converts this DKGResultHashSignatureMessage to a byte array suitable
// for network communication.
func (d *DKGResultHashSignatureMessage) Marshal() ([]byte, error) {
	return (&pb.DKGResultHashSignature{
		SenderIndex: uint32(d.senderIndex),
		ResultHash:  d.resultHash[:],
		Signature:   d.signature,
		PublicKey:   operator.Marshal(d.publicKey),
	}).Marshal()
}

// Unmarshal converts a byte array produced by Marshal to a
// DKGResultHashSignatureMessage.
func (d *DKGResultHashSignatureMessage) Unmarshal(bytes []byte) error {
	pbMsg := pb.DKGResultHashSignature{}
	if err := pbMsg.Unmarshal(bytes); err != nil {
		return err
	}
	d.senderIndex = group.MemberIndex(pbMsg.SenderIndex)

	resultHash, err := chain.DKGResultHashFromBytes(pbMsg.ResultHash)
	if err != nil {
		return err
	}
	d.resultHash = resultHash

	d.signature = pbMsg.Signature

	publicKey, err := operator.Unmarshal(pbMsg.PublicKey)
	if err != nil {
		return err
	}
	d.publicKey = publicKey

	return nil
}
