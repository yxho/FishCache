package FishCache

import pb "fishcachepb"

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	rpcGet(in *pb.Request, out *pb.Response) error
}
