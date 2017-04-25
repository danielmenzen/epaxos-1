package epaxospb

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
)

// ReplicaID is the id of a replica in an EPaxos deployment.
type ReplicaID uint64

// InstanceNum is an instance number ... instance slots ...
type InstanceNum uint64

// SeqNum is a sequence number ... used to break ties ...
type SeqNum uint64

// MaxSeqNum returns the maximum sequence number.
func MaxSeqNum(a, b SeqNum) SeqNum {
	if a > b {
		return a
	}
	return b
}

// WithDestination returns the message with the provided destination.
func (msg Message) WithDestination(dest ReplicaID) Message {
	msg.To = dest
	return msg
}

// WrapMessageInner wraps a union type of Message in a new isMessage_Type.
func WrapMessageInner(msg proto.Message) isMessage_Type {
	switch t := msg.(type) {
	case *PreAccept:
		return &Message_PreAccept{PreAccept: t}
	case *PreAcceptOK:
		return &Message_PreAcceptOk{PreAcceptOk: t}
	case *PreAcceptReply:
		return &Message_PreAcceptReply{PreAcceptReply: t}
	case *Accept:
		return &Message_Accept{Accept: t}
	case *AcceptOK:
		return &Message_AcceptOk{AcceptOk: t}
	case *Commit:
		return &Message_Commit{Commit: t}
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in WrapMessageInner", t))
	}
}

// WrapMessage wraps a union type of Message in a new Message without a
// destination.
func WrapMessage(msg proto.Message) Message {
	return Message{Type: WrapMessageInner(msg)}
}

// IsReply returns whether the message type is a reply or not.
func IsReply(t isMessage_Type) bool {
	switch t.(type) {
	case *Message_PreAcceptOk:
	case *Message_PreAcceptReply:
	case *Message_AcceptOk:
	default:
		return false
	}
	return true
}
