package data

type Peer struct {
	ID string
	IP string
}

type PeerList struct {
	Peers  []Peer
	Length uint32
}
