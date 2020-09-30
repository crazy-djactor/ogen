// Code generated by fastssz. DO NOT EDIT.
package p2p

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the MsgVersion object
func (m *MsgVersion) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the MsgVersion object to a target array
func (m *MsgVersion) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Tip'
	dst = ssz.MarshalUint64(dst, m.Tip)

	// Field (1) 'TipSlot'
	dst = ssz.MarshalUint64(dst, m.TipSlot)

	// Field (2) 'TipHash'
	dst = append(dst, m.TipHash[:]...)

	// Field (3) 'Nonce'
	dst = ssz.MarshalUint64(dst, m.Nonce)

	// Field (4) 'Timestamp'
	dst = ssz.MarshalUint64(dst, m.Timestamp)

	// Field (5) 'JustifiedSlot'
	dst = ssz.MarshalUint64(dst, m.JustifiedSlot)

	// Field (6) 'JustifiedHeight'
	dst = ssz.MarshalUint64(dst, m.JustifiedHeight)

	// Field (7) 'JustifiedHash'
	dst = append(dst, m.JustifiedHash[:]...)

	// Field (8) 'FinalizedSlot'
	dst = ssz.MarshalUint64(dst, m.FinalizedSlot)

	// Field (9) 'FinalizedHeight'
	dst = ssz.MarshalUint64(dst, m.FinalizedHeight)

	// Field (10) 'FinalizedHash'
	dst = append(dst, m.FinalizedHash[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the MsgVersion object
func (m *MsgVersion) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 160 {
		return ssz.ErrSize
	}

	// Field (0) 'Tip'
	m.Tip = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'TipSlot'
	m.TipSlot = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'TipHash'
	copy(m.TipHash[:], buf[16:48])

	// Field (3) 'Nonce'
	m.Nonce = ssz.UnmarshallUint64(buf[48:56])

	// Field (4) 'Timestamp'
	m.Timestamp = ssz.UnmarshallUint64(buf[56:64])

	// Field (5) 'JustifiedSlot'
	m.JustifiedSlot = ssz.UnmarshallUint64(buf[64:72])

	// Field (6) 'JustifiedHeight'
	m.JustifiedHeight = ssz.UnmarshallUint64(buf[72:80])

	// Field (7) 'JustifiedHash'
	copy(m.JustifiedHash[:], buf[80:112])

	// Field (8) 'FinalizedSlot'
	m.FinalizedSlot = ssz.UnmarshallUint64(buf[112:120])

	// Field (9) 'FinalizedHeight'
	m.FinalizedHeight = ssz.UnmarshallUint64(buf[120:128])

	// Field (10) 'FinalizedHash'
	copy(m.FinalizedHash[:], buf[128:160])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MsgVersion object
func (m *MsgVersion) SizeSSZ() (size int) {
	size = 160
	return
}

// HashTreeRoot ssz hashes the MsgVersion object
func (m *MsgVersion) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the MsgVersion object with a hasher
func (m *MsgVersion) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Tip'
	hh.PutUint64(m.Tip)

	// Field (1) 'TipSlot'
	hh.PutUint64(m.TipSlot)

	// Field (2) 'TipHash'
	hh.PutBytes(m.TipHash[:])

	// Field (3) 'Nonce'
	hh.PutUint64(m.Nonce)

	// Field (4) 'Timestamp'
	hh.PutUint64(m.Timestamp)

	// Field (5) 'JustifiedSlot'
	hh.PutUint64(m.JustifiedSlot)

	// Field (6) 'JustifiedHeight'
	hh.PutUint64(m.JustifiedHeight)

	// Field (7) 'JustifiedHash'
	hh.PutBytes(m.JustifiedHash[:])

	// Field (8) 'FinalizedSlot'
	hh.PutUint64(m.FinalizedSlot)

	// Field (9) 'FinalizedHeight'
	hh.PutUint64(m.FinalizedHeight)

	// Field (10) 'FinalizedHash'
	hh.PutBytes(m.FinalizedHash[:])

	hh.Merkleize(indx)
	return
}
