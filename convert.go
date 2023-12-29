package jito_client

import (
	"jito_client/proto"
	"github.com/desperatee/solana-go"
)

func ConvertTransactionToProtobufPacket(transaction *solana.Transaction) (proto.Packet, error) {
	rawTx, err := transaction.MarshalBinary()
	if err != nil {
		return proto.Packet{}, err
	}
	return proto.Packet{
		Data: rawTx,
		Meta: &proto.Meta{
			Size:        uint64(len(rawTx)),
			Addr:        "0.0.0.0",
			Port:        0,
			Flags:       nil,
			SenderStake: 0,
		},
	}, nil
}
