package signer

import (
	"errors"

	client "github.com/tendermint/tendermint/rpc/lib/client"
)

// RemoteCosigner uses tendermint rpc to request signing from a remote cosigner
type RemoteCosigner struct {
	id      int
	address string
}

// NewRemoteCosigner returns a newly initialized RemoteCosigner
func NewRemoteCosigner(id int, address string) *RemoteCosigner {
	cosigner := &RemoteCosigner{
		id:      id,
		address: address,
	}
	return cosigner
}

// GetID returns the ID of the remote cosigner
// Implements the cosigner interface
func (cosigner *RemoteCosigner) GetID() int {
	return cosigner.id
}

// Sign the sign request using the cosigner's share
// Return the signed bytes or an error
func (cosigner *RemoteCosigner) Sign(signReq CosignerSignRequest) (CosignerSignResponse, error) {
	params := map[string]interface{}{
		"arg": RpcSignRequest{
			SignBytes: signReq.SignBytes,
		},
	}

	remoteClient := client.NewJSONRPCClient(cosigner.address)
	result := &CosignerSignResponse{}
	_, err := remoteClient.Call("Sign", params, result)
	if err != nil {
		return CosignerSignResponse{}, err
	}

	return CosignerSignResponse{
		Timestamp: result.Timestamp,
		Signature: result.Signature,
	}, nil
}

func (cosigner *RemoteCosigner) GetEphemeralSecretPart(req CosignerGetEphemeralSecretPartRequest) (CosignerGetEphemeralSecretPartResponse, error) {
	resp := CosignerGetEphemeralSecretPartResponse{}

	params := map[string]interface{}{
		"arg": RpcGetEphemeralSecretPartRequest{
			ID:     req.ID,
			Height: req.Height,
			Round:  req.Round,
			Step:   req.Step,
		},
	}

	remoteClient := client.NewJSONRPCClient(cosigner.address)
	result := &RpcGetEphemeralSecretPartResponse{}
	_, err := remoteClient.Call("GetEphemeralSecretPart", params, result)
	if err != nil {
		return CosignerGetEphemeralSecretPartResponse{}, err
	}

	resp.SourceID = result.SourceID
	resp.SourceEphemeralSecretPublicKey = result.SourceEphemeralSecretPublicKey
	resp.EncryptedSharePart = result.EncryptedSharePart
	resp.SourceSig = result.SourceSig

	return resp, nil
}

func (cosigner *RemoteCosigner) HasEphemeralSecretPart(req CosignerHasEphemeralSecretPartRequest) (CosignerHasEphemeralSecretPartResponse, error) {
	res := CosignerHasEphemeralSecretPartResponse{}
	return res, errors.New("Not Implemented")
}

func (cosigner *RemoteCosigner) SetEphemeralSecretPart(req CosignerSetEphemeralSecretPartRequest) error {
	return errors.New("Not Implemented")
}
