// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type SegReq struct{ capnp.Struct }
type SegReq_flags SegReq

// SegReq_TypeID is the unique identifier for the type SegReq.
const SegReq_TypeID = 0x9d0135027d04861e

func NewSegReq(s *capnp.Segment) (SegReq, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 0})
	return SegReq{st}, err
}

func NewRootSegReq(s *capnp.Segment) (SegReq, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 0})
	return SegReq{st}, err
}

func ReadRootSegReq(msg *capnp.Message) (SegReq, error) {
	root, err := msg.RootPtr()
	return SegReq{root.Struct()}, err
}

func (s SegReq) String() string {
	str, _ := text.Marshal(0x9d0135027d04861e, s.Struct)
	return str
}

func (s SegReq) SrcIA() uint64 {
	return s.Struct.Uint64(0)
}

func (s SegReq) SetSrcIA(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s SegReq) DstIA() uint64 {
	return s.Struct.Uint64(8)
}

func (s SegReq) SetDstIA(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s SegReq) Flags() SegReq_flags { return SegReq_flags(s) }

func (s SegReq_flags) Sibra() bool {
	return s.Struct.Bit(128)
}

func (s SegReq_flags) SetSibra(v bool) {
	s.Struct.SetBit(128, v)
}

func (s SegReq_flags) CacheOnly() bool {
	return s.Struct.Bit(129)
}

func (s SegReq_flags) SetCacheOnly(v bool) {
	s.Struct.SetBit(129, v)
}

// SegReq_List is a list of SegReq.
type SegReq_List struct{ capnp.List }

// NewSegReq creates a new list of SegReq.
func NewSegReq_List(s *capnp.Segment, sz int32) (SegReq_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 0}, sz)
	return SegReq_List{l}, err
}

func (s SegReq_List) At(i int) SegReq { return SegReq{s.List.Struct(i)} }

func (s SegReq_List) Set(i int, v SegReq) error { return s.List.SetStruct(i, v.Struct) }

func (s SegReq_List) String() string {
	str, _ := text.MarshalList(0x9d0135027d04861e, s.List)
	return str
}

// SegReq_Promise is a wrapper for a SegReq promised by a client call.
type SegReq_Promise struct{ *capnp.Pipeline }

func (p SegReq_Promise) Struct() (SegReq, error) {
	s, err := p.Pipeline.Struct()
	return SegReq{s}, err
}

func (p SegReq_Promise) Flags() SegReq_flags_Promise { return SegReq_flags_Promise{p.Pipeline} }

// SegReq_flags_Promise is a wrapper for a SegReq_flags promised by a client call.
type SegReq_flags_Promise struct{ *capnp.Pipeline }

func (p SegReq_flags_Promise) Struct() (SegReq_flags, error) {
	s, err := p.Pipeline.Struct()
	return SegReq_flags{s}, err
}

type SegRecs struct{ capnp.Struct }

// SegRecs_TypeID is the unique identifier for the type SegRecs.
const SegRecs_TypeID = 0x934ba70bfd144ebd

func NewSegRecs(s *capnp.Segment) (SegRecs, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SegRecs{st}, err
}

func NewRootSegRecs(s *capnp.Segment) (SegRecs, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SegRecs{st}, err
}

func ReadRootSegRecs(msg *capnp.Message) (SegRecs, error) {
	root, err := msg.RootPtr()
	return SegRecs{root.Struct()}, err
}

func (s SegRecs) String() string {
	str, _ := text.Marshal(0x934ba70bfd144ebd, s.Struct)
	return str
}

func (s SegRecs) Recs() (PathSegMeta_List, error) {
	p, err := s.Struct.Ptr(0)
	return PathSegMeta_List{List: p.List()}, err
}

func (s SegRecs) HasRecs() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SegRecs) SetRecs(v PathSegMeta_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewRecs sets the recs field to a newly
// allocated PathSegMeta_List, preferring placement in s's segment.
func (s SegRecs) NewRecs(n int32) (PathSegMeta_List, error) {
	l, err := NewPathSegMeta_List(s.Struct.Segment(), n)
	if err != nil {
		return PathSegMeta_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s SegRecs) SRevInfos() (SignedBlob_List, error) {
	p, err := s.Struct.Ptr(1)
	return SignedBlob_List{List: p.List()}, err
}

func (s SegRecs) HasSRevInfos() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SegRecs) SetSRevInfos(v SignedBlob_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewSRevInfos sets the sRevInfos field to a newly
// allocated SignedBlob_List, preferring placement in s's segment.
func (s SegRecs) NewSRevInfos(n int32) (SignedBlob_List, error) {
	l, err := NewSignedBlob_List(s.Struct.Segment(), n)
	if err != nil {
		return SignedBlob_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// SegRecs_List is a list of SegRecs.
type SegRecs_List struct{ capnp.List }

// NewSegRecs creates a new list of SegRecs.
func NewSegRecs_List(s *capnp.Segment, sz int32) (SegRecs_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SegRecs_List{l}, err
}

func (s SegRecs_List) At(i int) SegRecs { return SegRecs{s.List.Struct(i)} }

func (s SegRecs_List) Set(i int, v SegRecs) error { return s.List.SetStruct(i, v.Struct) }

func (s SegRecs_List) String() string {
	str, _ := text.MarshalList(0x934ba70bfd144ebd, s.List)
	return str
}

// SegRecs_Promise is a wrapper for a SegRecs promised by a client call.
type SegRecs_Promise struct{ *capnp.Pipeline }

func (p SegRecs_Promise) Struct() (SegRecs, error) {
	s, err := p.Pipeline.Struct()
	return SegRecs{s}, err
}

type SegReply struct{ capnp.Struct }

// SegReply_TypeID is the unique identifier for the type SegReply.
const SegReply_TypeID = 0x9359e1b2db37dbbb

func NewSegReply(s *capnp.Segment) (SegReply, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SegReply{st}, err
}

func NewRootSegReply(s *capnp.Segment) (SegReply, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SegReply{st}, err
}

func ReadRootSegReply(msg *capnp.Message) (SegReply, error) {
	root, err := msg.RootPtr()
	return SegReply{root.Struct()}, err
}

func (s SegReply) String() string {
	str, _ := text.Marshal(0x9359e1b2db37dbbb, s.Struct)
	return str
}

func (s SegReply) Req() (SegReq, error) {
	p, err := s.Struct.Ptr(0)
	return SegReq{Struct: p.Struct()}, err
}

func (s SegReply) HasReq() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SegReply) SetReq(v SegReq) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewReq sets the req field to a newly
// allocated SegReq struct, preferring placement in s's segment.
func (s SegReply) NewReq() (SegReq, error) {
	ss, err := NewSegReq(s.Struct.Segment())
	if err != nil {
		return SegReq{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s SegReply) Recs() (SegRecs, error) {
	p, err := s.Struct.Ptr(1)
	return SegRecs{Struct: p.Struct()}, err
}

func (s SegReply) HasRecs() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SegReply) SetRecs(v SegRecs) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewRecs sets the recs field to a newly
// allocated SegRecs struct, preferring placement in s's segment.
func (s SegReply) NewRecs() (SegRecs, error) {
	ss, err := NewSegRecs(s.Struct.Segment())
	if err != nil {
		return SegRecs{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// SegReply_List is a list of SegReply.
type SegReply_List struct{ capnp.List }

// NewSegReply creates a new list of SegReply.
func NewSegReply_List(s *capnp.Segment, sz int32) (SegReply_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SegReply_List{l}, err
}

func (s SegReply_List) At(i int) SegReply { return SegReply{s.List.Struct(i)} }

func (s SegReply_List) Set(i int, v SegReply) error { return s.List.SetStruct(i, v.Struct) }

func (s SegReply_List) String() string {
	str, _ := text.MarshalList(0x9359e1b2db37dbbb, s.List)
	return str
}

// SegReply_Promise is a wrapper for a SegReply promised by a client call.
type SegReply_Promise struct{ *capnp.Pipeline }

func (p SegReply_Promise) Struct() (SegReply, error) {
	s, err := p.Pipeline.Struct()
	return SegReply{s}, err
}

func (p SegReply_Promise) Req() SegReq_Promise {
	return SegReq_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p SegReply_Promise) Recs() SegRecs_Promise {
	return SegRecs_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type SegChangesIdReq struct{ capnp.Struct }

// SegChangesIdReq_TypeID is the unique identifier for the type SegChangesIdReq.
const SegChangesIdReq_TypeID = 0xc88dfa6be7a1d091

func NewSegChangesIdReq(s *capnp.Segment) (SegChangesIdReq, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return SegChangesIdReq{st}, err
}

func NewRootSegChangesIdReq(s *capnp.Segment) (SegChangesIdReq, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return SegChangesIdReq{st}, err
}

func ReadRootSegChangesIdReq(msg *capnp.Message) (SegChangesIdReq, error) {
	root, err := msg.RootPtr()
	return SegChangesIdReq{root.Struct()}, err
}

func (s SegChangesIdReq) String() string {
	str, _ := text.Marshal(0xc88dfa6be7a1d091, s.Struct)
	return str
}

func (s SegChangesIdReq) LastCheck() uint32 {
	return s.Struct.Uint32(0)
}

func (s SegChangesIdReq) SetLastCheck(v uint32) {
	s.Struct.SetUint32(0, v)
}

// SegChangesIdReq_List is a list of SegChangesIdReq.
type SegChangesIdReq_List struct{ capnp.List }

// NewSegChangesIdReq creates a new list of SegChangesIdReq.
func NewSegChangesIdReq_List(s *capnp.Segment, sz int32) (SegChangesIdReq_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return SegChangesIdReq_List{l}, err
}

func (s SegChangesIdReq_List) At(i int) SegChangesIdReq { return SegChangesIdReq{s.List.Struct(i)} }

func (s SegChangesIdReq_List) Set(i int, v SegChangesIdReq) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s SegChangesIdReq_List) String() string {
	str, _ := text.MarshalList(0xc88dfa6be7a1d091, s.List)
	return str
}

// SegChangesIdReq_Promise is a wrapper for a SegChangesIdReq promised by a client call.
type SegChangesIdReq_Promise struct{ *capnp.Pipeline }

func (p SegChangesIdReq_Promise) Struct() (SegChangesIdReq, error) {
	s, err := p.Pipeline.Struct()
	return SegChangesIdReq{s}, err
}

type SegIds struct{ capnp.Struct }

// SegIds_TypeID is the unique identifier for the type SegIds.
const SegIds_TypeID = 0xabf979c3f68dae4b

func NewSegIds(s *capnp.Segment) (SegIds, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SegIds{st}, err
}

func NewRootSegIds(s *capnp.Segment) (SegIds, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SegIds{st}, err
}

func ReadRootSegIds(msg *capnp.Message) (SegIds, error) {
	root, err := msg.RootPtr()
	return SegIds{root.Struct()}, err
}

func (s SegIds) String() string {
	str, _ := text.Marshal(0xabf979c3f68dae4b, s.Struct)
	return str
}

func (s SegIds) SegId() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s SegIds) HasSegId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SegIds) SetSegId(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s SegIds) FullId() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s SegIds) HasFullId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SegIds) SetFullId(v []byte) error {
	return s.Struct.SetData(1, v)
}

// SegIds_List is a list of SegIds.
type SegIds_List struct{ capnp.List }

// NewSegIds creates a new list of SegIds.
func NewSegIds_List(s *capnp.Segment, sz int32) (SegIds_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SegIds_List{l}, err
}

func (s SegIds_List) At(i int) SegIds { return SegIds{s.List.Struct(i)} }

func (s SegIds_List) Set(i int, v SegIds) error { return s.List.SetStruct(i, v.Struct) }

func (s SegIds_List) String() string {
	str, _ := text.MarshalList(0xabf979c3f68dae4b, s.List)
	return str
}

// SegIds_Promise is a wrapper for a SegIds promised by a client call.
type SegIds_Promise struct{ *capnp.Pipeline }

func (p SegIds_Promise) Struct() (SegIds, error) {
	s, err := p.Pipeline.Struct()
	return SegIds{s}, err
}

type SegChangesIdReply struct{ capnp.Struct }

// SegChangesIdReply_TypeID is the unique identifier for the type SegChangesIdReply.
const SegChangesIdReply_TypeID = 0xbd56ceeaf8c65140

func NewSegChangesIdReply(s *capnp.Segment) (SegChangesIdReply, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SegChangesIdReply{st}, err
}

func NewRootSegChangesIdReply(s *capnp.Segment) (SegChangesIdReply, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SegChangesIdReply{st}, err
}

func ReadRootSegChangesIdReply(msg *capnp.Message) (SegChangesIdReply, error) {
	root, err := msg.RootPtr()
	return SegChangesIdReply{root.Struct()}, err
}

func (s SegChangesIdReply) String() string {
	str, _ := text.Marshal(0xbd56ceeaf8c65140, s.Struct)
	return str
}

func (s SegChangesIdReply) Ids() (SegIds_List, error) {
	p, err := s.Struct.Ptr(0)
	return SegIds_List{List: p.List()}, err
}

func (s SegChangesIdReply) HasIds() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SegChangesIdReply) SetIds(v SegIds_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewIds sets the ids field to a newly
// allocated SegIds_List, preferring placement in s's segment.
func (s SegChangesIdReply) NewIds(n int32) (SegIds_List, error) {
	l, err := NewSegIds_List(s.Struct.Segment(), n)
	if err != nil {
		return SegIds_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// SegChangesIdReply_List is a list of SegChangesIdReply.
type SegChangesIdReply_List struct{ capnp.List }

// NewSegChangesIdReply creates a new list of SegChangesIdReply.
func NewSegChangesIdReply_List(s *capnp.Segment, sz int32) (SegChangesIdReply_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SegChangesIdReply_List{l}, err
}

func (s SegChangesIdReply_List) At(i int) SegChangesIdReply {
	return SegChangesIdReply{s.List.Struct(i)}
}

func (s SegChangesIdReply_List) Set(i int, v SegChangesIdReply) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s SegChangesIdReply_List) String() string {
	str, _ := text.MarshalList(0xbd56ceeaf8c65140, s.List)
	return str
}

// SegChangesIdReply_Promise is a wrapper for a SegChangesIdReply promised by a client call.
type SegChangesIdReply_Promise struct{ *capnp.Pipeline }

func (p SegChangesIdReply_Promise) Struct() (SegChangesIdReply, error) {
	s, err := p.Pipeline.Struct()
	return SegChangesIdReply{s}, err
}

type SegChangesReq struct{ capnp.Struct }

// SegChangesReq_TypeID is the unique identifier for the type SegChangesReq.
const SegChangesReq_TypeID = 0xa7ad0c62a234c68b

func NewSegChangesReq(s *capnp.Segment) (SegChangesReq, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SegChangesReq{st}, err
}

func NewRootSegChangesReq(s *capnp.Segment) (SegChangesReq, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SegChangesReq{st}, err
}

func ReadRootSegChangesReq(msg *capnp.Message) (SegChangesReq, error) {
	root, err := msg.RootPtr()
	return SegChangesReq{root.Struct()}, err
}

func (s SegChangesReq) String() string {
	str, _ := text.Marshal(0xa7ad0c62a234c68b, s.Struct)
	return str
}

func (s SegChangesReq) SegIds() (capnp.DataList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.DataList{List: p.List()}, err
}

func (s SegChangesReq) HasSegIds() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SegChangesReq) SetSegIds(v capnp.DataList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewSegIds sets the segIds field to a newly
// allocated capnp.DataList, preferring placement in s's segment.
func (s SegChangesReq) NewSegIds(n int32) (capnp.DataList, error) {
	l, err := capnp.NewDataList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.DataList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// SegChangesReq_List is a list of SegChangesReq.
type SegChangesReq_List struct{ capnp.List }

// NewSegChangesReq creates a new list of SegChangesReq.
func NewSegChangesReq_List(s *capnp.Segment, sz int32) (SegChangesReq_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SegChangesReq_List{l}, err
}

func (s SegChangesReq_List) At(i int) SegChangesReq { return SegChangesReq{s.List.Struct(i)} }

func (s SegChangesReq_List) Set(i int, v SegChangesReq) error { return s.List.SetStruct(i, v.Struct) }

func (s SegChangesReq_List) String() string {
	str, _ := text.MarshalList(0xa7ad0c62a234c68b, s.List)
	return str
}

// SegChangesReq_Promise is a wrapper for a SegChangesReq promised by a client call.
type SegChangesReq_Promise struct{ *capnp.Pipeline }

func (p SegChangesReq_Promise) Struct() (SegChangesReq, error) {
	s, err := p.Pipeline.Struct()
	return SegChangesReq{s}, err
}

type PathMgmt struct{ capnp.Struct }
type PathMgmt_Which uint16

const (
	PathMgmt_Which_unset             PathMgmt_Which = 0
	PathMgmt_Which_segReq            PathMgmt_Which = 1
	PathMgmt_Which_segReply          PathMgmt_Which = 2
	PathMgmt_Which_segReg            PathMgmt_Which = 3
	PathMgmt_Which_segSync           PathMgmt_Which = 4
	PathMgmt_Which_sRevInfo          PathMgmt_Which = 5
	PathMgmt_Which_ifStateReq        PathMgmt_Which = 6
	PathMgmt_Which_ifStateInfos      PathMgmt_Which = 7
	PathMgmt_Which_segChangesIdReq   PathMgmt_Which = 8
	PathMgmt_Which_segChangesIdReply PathMgmt_Which = 9
	PathMgmt_Which_segChangesReq     PathMgmt_Which = 10
	PathMgmt_Which_segChangesReply   PathMgmt_Which = 11
)

func (w PathMgmt_Which) String() string {
	const s = "unsetsegReqsegReplysegRegsegSyncsRevInfoifStateReqifStateInfossegChangesIdReqsegChangesIdReplysegChangesReqsegChangesReply"
	switch w {
	case PathMgmt_Which_unset:
		return s[0:5]
	case PathMgmt_Which_segReq:
		return s[5:11]
	case PathMgmt_Which_segReply:
		return s[11:19]
	case PathMgmt_Which_segReg:
		return s[19:25]
	case PathMgmt_Which_segSync:
		return s[25:32]
	case PathMgmt_Which_sRevInfo:
		return s[32:40]
	case PathMgmt_Which_ifStateReq:
		return s[40:50]
	case PathMgmt_Which_ifStateInfos:
		return s[50:62]
	case PathMgmt_Which_segChangesIdReq:
		return s[62:77]
	case PathMgmt_Which_segChangesIdReply:
		return s[77:94]
	case PathMgmt_Which_segChangesReq:
		return s[94:107]
	case PathMgmt_Which_segChangesReply:
		return s[107:122]

	}
	return "PathMgmt_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// PathMgmt_TypeID is the unique identifier for the type PathMgmt.
const PathMgmt_TypeID = 0xa12cfa420c9ad0ca

func NewPathMgmt(s *capnp.Segment) (PathMgmt, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PathMgmt{st}, err
}

func NewRootPathMgmt(s *capnp.Segment) (PathMgmt, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PathMgmt{st}, err
}

func ReadRootPathMgmt(msg *capnp.Message) (PathMgmt, error) {
	root, err := msg.RootPtr()
	return PathMgmt{root.Struct()}, err
}

func (s PathMgmt) String() string {
	str, _ := text.Marshal(0xa12cfa420c9ad0ca, s.Struct)
	return str
}

func (s PathMgmt) Which() PathMgmt_Which {
	return PathMgmt_Which(s.Struct.Uint16(0))
}
func (s PathMgmt) SetUnset() {
	s.Struct.SetUint16(0, 0)

}

func (s PathMgmt) SegReq() (SegReq, error) {
	p, err := s.Struct.Ptr(0)
	return SegReq{Struct: p.Struct()}, err
}

func (s PathMgmt) HasSegReq() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathMgmt) SetSegReq(v SegReq) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSegReq sets the segReq field to a newly
// allocated SegReq struct, preferring placement in s's segment.
func (s PathMgmt) NewSegReq() (SegReq, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewSegReq(s.Struct.Segment())
	if err != nil {
		return SegReq{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PathMgmt) SegReply() (SegReply, error) {
	p, err := s.Struct.Ptr(0)
	return SegReply{Struct: p.Struct()}, err
}

func (s PathMgmt) HasSegReply() bool {
	if s.Struct.Uint16(0) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathMgmt) SetSegReply(v SegReply) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSegReply sets the segReply field to a newly
// allocated SegReply struct, preferring placement in s's segment.
func (s PathMgmt) NewSegReply() (SegReply, error) {
	s.Struct.SetUint16(0, 2)
	ss, err := NewSegReply(s.Struct.Segment())
	if err != nil {
		return SegReply{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PathMgmt) SegReg() (SegRecs, error) {
	p, err := s.Struct.Ptr(0)
	return SegRecs{Struct: p.Struct()}, err
}

func (s PathMgmt) HasSegReg() bool {
	if s.Struct.Uint16(0) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathMgmt) SetSegReg(v SegRecs) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSegReg sets the segReg field to a newly
// allocated SegRecs struct, preferring placement in s's segment.
func (s PathMgmt) NewSegReg() (SegRecs, error) {
	s.Struct.SetUint16(0, 3)
	ss, err := NewSegRecs(s.Struct.Segment())
	if err != nil {
		return SegRecs{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PathMgmt) SegSync() (SegRecs, error) {
	p, err := s.Struct.Ptr(0)
	return SegRecs{Struct: p.Struct()}, err
}

func (s PathMgmt) HasSegSync() bool {
	if s.Struct.Uint16(0) != 4 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathMgmt) SetSegSync(v SegRecs) error {
	s.Struct.SetUint16(0, 4)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSegSync sets the segSync field to a newly
// allocated SegRecs struct, preferring placement in s's segment.
func (s PathMgmt) NewSegSync() (SegRecs, error) {
	s.Struct.SetUint16(0, 4)
	ss, err := NewSegRecs(s.Struct.Segment())
	if err != nil {
		return SegRecs{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PathMgmt) SRevInfo() (SignedBlob, error) {
	p, err := s.Struct.Ptr(0)
	return SignedBlob{Struct: p.Struct()}, err
}

func (s PathMgmt) HasSRevInfo() bool {
	if s.Struct.Uint16(0) != 5 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathMgmt) SetSRevInfo(v SignedBlob) error {
	s.Struct.SetUint16(0, 5)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSRevInfo sets the sRevInfo field to a newly
// allocated SignedBlob struct, preferring placement in s's segment.
func (s PathMgmt) NewSRevInfo() (SignedBlob, error) {
	s.Struct.SetUint16(0, 5)
	ss, err := NewSignedBlob(s.Struct.Segment())
	if err != nil {
		return SignedBlob{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PathMgmt) IfStateReq() (IFStateReq, error) {
	p, err := s.Struct.Ptr(0)
	return IFStateReq{Struct: p.Struct()}, err
}

func (s PathMgmt) HasIfStateReq() bool {
	if s.Struct.Uint16(0) != 6 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathMgmt) SetIfStateReq(v IFStateReq) error {
	s.Struct.SetUint16(0, 6)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewIfStateReq sets the ifStateReq field to a newly
// allocated IFStateReq struct, preferring placement in s's segment.
func (s PathMgmt) NewIfStateReq() (IFStateReq, error) {
	s.Struct.SetUint16(0, 6)
	ss, err := NewIFStateReq(s.Struct.Segment())
	if err != nil {
		return IFStateReq{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PathMgmt) IfStateInfos() (IFStateInfos, error) {
	p, err := s.Struct.Ptr(0)
	return IFStateInfos{Struct: p.Struct()}, err
}

func (s PathMgmt) HasIfStateInfos() bool {
	if s.Struct.Uint16(0) != 7 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathMgmt) SetIfStateInfos(v IFStateInfos) error {
	s.Struct.SetUint16(0, 7)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewIfStateInfos sets the ifStateInfos field to a newly
// allocated IFStateInfos struct, preferring placement in s's segment.
func (s PathMgmt) NewIfStateInfos() (IFStateInfos, error) {
	s.Struct.SetUint16(0, 7)
	ss, err := NewIFStateInfos(s.Struct.Segment())
	if err != nil {
		return IFStateInfos{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PathMgmt) SegChangesIdReq() (SegChangesIdReq, error) {
	p, err := s.Struct.Ptr(0)
	return SegChangesIdReq{Struct: p.Struct()}, err
}

func (s PathMgmt) HasSegChangesIdReq() bool {
	if s.Struct.Uint16(0) != 8 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathMgmt) SetSegChangesIdReq(v SegChangesIdReq) error {
	s.Struct.SetUint16(0, 8)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSegChangesIdReq sets the segChangesIdReq field to a newly
// allocated SegChangesIdReq struct, preferring placement in s's segment.
func (s PathMgmt) NewSegChangesIdReq() (SegChangesIdReq, error) {
	s.Struct.SetUint16(0, 8)
	ss, err := NewSegChangesIdReq(s.Struct.Segment())
	if err != nil {
		return SegChangesIdReq{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PathMgmt) SegChangesIdReply() (SegChangesIdReply, error) {
	p, err := s.Struct.Ptr(0)
	return SegChangesIdReply{Struct: p.Struct()}, err
}

func (s PathMgmt) HasSegChangesIdReply() bool {
	if s.Struct.Uint16(0) != 9 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathMgmt) SetSegChangesIdReply(v SegChangesIdReply) error {
	s.Struct.SetUint16(0, 9)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSegChangesIdReply sets the segChangesIdReply field to a newly
// allocated SegChangesIdReply struct, preferring placement in s's segment.
func (s PathMgmt) NewSegChangesIdReply() (SegChangesIdReply, error) {
	s.Struct.SetUint16(0, 9)
	ss, err := NewSegChangesIdReply(s.Struct.Segment())
	if err != nil {
		return SegChangesIdReply{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PathMgmt) SegChangesReq() (SegChangesReq, error) {
	p, err := s.Struct.Ptr(0)
	return SegChangesReq{Struct: p.Struct()}, err
}

func (s PathMgmt) HasSegChangesReq() bool {
	if s.Struct.Uint16(0) != 10 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathMgmt) SetSegChangesReq(v SegChangesReq) error {
	s.Struct.SetUint16(0, 10)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSegChangesReq sets the segChangesReq field to a newly
// allocated SegChangesReq struct, preferring placement in s's segment.
func (s PathMgmt) NewSegChangesReq() (SegChangesReq, error) {
	s.Struct.SetUint16(0, 10)
	ss, err := NewSegChangesReq(s.Struct.Segment())
	if err != nil {
		return SegChangesReq{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PathMgmt) SegChangesReply() (SegRecs, error) {
	p, err := s.Struct.Ptr(0)
	return SegRecs{Struct: p.Struct()}, err
}

func (s PathMgmt) HasSegChangesReply() bool {
	if s.Struct.Uint16(0) != 11 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathMgmt) SetSegChangesReply(v SegRecs) error {
	s.Struct.SetUint16(0, 11)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSegChangesReply sets the segChangesReply field to a newly
// allocated SegRecs struct, preferring placement in s's segment.
func (s PathMgmt) NewSegChangesReply() (SegRecs, error) {
	s.Struct.SetUint16(0, 11)
	ss, err := NewSegRecs(s.Struct.Segment())
	if err != nil {
		return SegRecs{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// PathMgmt_List is a list of PathMgmt.
type PathMgmt_List struct{ capnp.List }

// NewPathMgmt creates a new list of PathMgmt.
func NewPathMgmt_List(s *capnp.Segment, sz int32) (PathMgmt_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return PathMgmt_List{l}, err
}

func (s PathMgmt_List) At(i int) PathMgmt { return PathMgmt{s.List.Struct(i)} }

func (s PathMgmt_List) Set(i int, v PathMgmt) error { return s.List.SetStruct(i, v.Struct) }

func (s PathMgmt_List) String() string {
	str, _ := text.MarshalList(0xa12cfa420c9ad0ca, s.List)
	return str
}

// PathMgmt_Promise is a wrapper for a PathMgmt promised by a client call.
type PathMgmt_Promise struct{ *capnp.Pipeline }

func (p PathMgmt_Promise) Struct() (PathMgmt, error) {
	s, err := p.Pipeline.Struct()
	return PathMgmt{s}, err
}

func (p PathMgmt_Promise) SegReq() SegReq_Promise {
	return SegReq_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PathMgmt_Promise) SegReply() SegReply_Promise {
	return SegReply_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PathMgmt_Promise) SegReg() SegRecs_Promise {
	return SegRecs_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PathMgmt_Promise) SegSync() SegRecs_Promise {
	return SegRecs_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PathMgmt_Promise) SRevInfo() SignedBlob_Promise {
	return SignedBlob_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PathMgmt_Promise) IfStateReq() IFStateReq_Promise {
	return IFStateReq_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PathMgmt_Promise) IfStateInfos() IFStateInfos_Promise {
	return IFStateInfos_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PathMgmt_Promise) SegChangesIdReq() SegChangesIdReq_Promise {
	return SegChangesIdReq_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PathMgmt_Promise) SegChangesIdReply() SegChangesIdReply_Promise {
	return SegChangesIdReply_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PathMgmt_Promise) SegChangesReq() SegChangesReq_Promise {
	return SegChangesReq_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PathMgmt_Promise) SegChangesReply() SegRecs_Promise {
	return SegRecs_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

const schema_8fcd13516850d142 = "x\xda|T\x7fh\x1be\x18~\x9f\xef.\xcdeM" +
	"m\xea\x9d(\x88V\xc7\x06N\xb6Q\xd3\x8di\xe9h" +
	"\xd7\xad\xb8\xb8Vs9\x1d\x9d(\x9a%\x97\xa46\xcd" +
	"\x92\\\xb6\x12\xdc\xd0\x89\x13E\x8b?\xff\x18\xe2\x18\x13" +
	"\xa4\x0cQQ\xf0\x8fU*LqH\xc11\xf1\xcf\x81" +
	" \"\xca\xc0\x7f\xc4\xcd\xad\xdb\xfc\xe4\xfd\xb2]\xb2k" +
	"\xe6\x7f\x97\xefy\xbe\xe7}\xbe7\xef\xf3\xf6\xfd&\x86" +
	"\xc5\x03\xa1\xa3\x06\x91\x9d\x0du\xc8\x85G\xad\xab\x9ds" +
	";\xde\xa5\x9e\x18\xe4\xc8\x8f\xc9\x82m\xfe\xf0&\x85D" +
	"\x98\xc8\xdc\xa8\x7fcn\xd6\xf9\xeb!}\x86 \xbf:" +
	"\xbb\xe9\xec\x17\xbf\xecj\xcb=\xa1/\x9a\xdf)\xeeI" +
	"\xc5\xbd\xfb\x15\xfd\x80\xd8\x88#d\xc7\xa05\xc9\x8aq" +
	"Wh\xde\\\x1d\xe2\xaf{C\x9f\x11\xe4\xe2\x99\xf7\xa3" +
	"#Kk\x8f1\xb7Ex\x14\xe1(\x91y$\xb4h" +
	"\x1egv\xffG\xa19A\x90\xaf\x9f\xda\xf0\xe1\xee\xe8" +
	"'s\x01\x1b`\xc1\x03\xc6_\xe6k\x06\x7f\x1d2\x86" +
	"\x08r\xc7\xa7\xb3\x17\xbe\xad_\xfa\xb8\x9d\xe5/\x8dy" +
	"sAqO\x18ly\xd8>u\xf1\xdc\xe9\x9d\x0bm" +
	"t\xfbo\x8b\xdc\x0asuDy\x8e\xb0\xf0\xdbg\x8e" +
	"\xfd>\xb54\xfb}\xc03\xbf\xaf\xdf\x8e\x08\x98O+" +
	"\xf2.E\xdew\xe7\xb6\xcb\xe3k\xde\xfb\x99\xec\xdb\xa1" +
	"5[c\x85\xc1\x96#\xe7\x08\xe6\xc1\xc8\x0c}.\xcb" +
	"\xe9Z\xe1\x99\xe9\xfc\xb4\xa8\xad\xcf\xa4\xcb\xa5\xf2\x80\xe3" +
	"\xe6Sn\xc6\xa3$`\x1b\x9aN\xa4\x83\xa8g\xcd\xfd" +
	"D\xf6*\x0d\xf6\xb0@\x0f`\x81\x0f7\xa7\x88\xecA" +
	"\x0d\xf6\x84@w\xd5\xcdx\xb8\x85\x90\xd4\x80\x98\xec\x9e" +
	"_1\xf8\xfc\xce\xc3G\x89\xc0\x87\xd2K\xb9\xfb\x12\xa5" +
	"\xdc\x1eB\x0b)\xf5\xeb\xd2\xa6C\x0f\xc7}R{/" +
	"\xe5b\x9d\x82nV^s\xd3\xd7\xe2f\x1d[\xbcO" +
	"\x83\xbdA \\u+\x885\xdfM@\x8c\xaey\x8c" +
	"5'\xb0q\xec\xd7\xc5\xf5\xbaC\xaap\x85\x8bF\xfd" +
	"\xa2\xa3q\"{X\x83=\xd6R4\xc1\x87\xdb4\xd8" +
	"I\x01\x08\xb44\xbeg<N\xa2\xd7\xabf\x12[\x10" +
	"!\x81\x08\xa17\xeb\xd5Z~\xe5\x8a\xe9\xbc\xb7\xfc\xd1" +
	"\xc9t\xad0\x9e\x9f\xae5\x1e\xdd\xa7\xe9Q)\xd9\x80" +
	"\xb9\x05q\"g\x10\x1a\x9c\xed\x10\xe8\xc2\xbfR\x990" +
	"G1@\xe4\x0c30\xc6\x80\xb8*-\x08\"3\x81" +
	"G\x88\x9c\xed\x0c<\xce\x80vEZ\xd0\x88L[\xdd" +
	"\x18c`\x82\x01\xfd\xb2\xb4\xa0\x13\x99O`\x84\xc8I" +
	"2\xf0\x14\x03\xa1%i!\xc4c\xa5\xa4&\x18\xc82" +
	"\xd0qIZ\xe8 2\xd3x\x92\xc8y\x96\x81\"\x03" +
	"\xe1\x8b\xd2R\xc1\x98\xc4sDN\x81\x81\x1a\x03\xc6?" +
	"\xd2\x82AdV\xf0\x12\x91Sf`?\x03\x91\x0b\xd2" +
	"B\x84\xc8\xac\xe3\x1d\"g?\x03\xaf2\xb0\xe2\xbc\xb4" +
	"\xb0\x82\x93\x85*\x91\xf32\x03o1\xd0\xf9\xb7\xb4\xd0" +
	"Id\xce*\xa97\x188\x0c\x81\xde\xbd%\xcf\xadQ" +
	"\xc7\x90\xa7\xfe\xbf\xe5\x13 =\x7f\xa2\x08\xb1\xe6zi" +
	"\xa0\x8dk\xf9\xe5\x13\xf2\x82\xe7\xe6\x9dz)\xd3fv" +
	"\xfc\xc1V\x82\xad\x13\xcd\xe8d\xce\xa9\xa5kn\x8a4" +
	"\xe5\xe6x\xee\x83\x93S3\x7f\xfe\x11\x80\x13\xd4]\xca" +
	"\xed\xe1\xc9\xfc\xba\xf3\xca\xf9\xf8\xe9{~j\xb1\xbb\xb5" +
	"\x90.\xe5\xe1z\x89l\xca\xadp\x0d\x7f\x0f\xb4\xe7\x94" +
	"\x8b\xa8#\xd6\\-\x01\x16\xf5\xba^\xa35\xfeR[" +
	"\xae\xd3h\xd1\xff$\xa55\xa1\xeaRC\x95\x07V\xf7" +
	"\x03\xd35@d\x1b\x1a\xecUB\xb56\x91\xf5\xb3\xdf" +
	"E\xe2\x86\xc4\xdf\x90\xbcD\xd6\x0b\xc4=\xde.\xee\x03" +
	"\xcd\xb8\xf7*u\xa5\xdaE\x18\xca\xed-\x16\x9b?\xfd" +
	"\"Z\xd0t\xb3a\x01\xe3+\x9b\xc6\xc3\x93\xd9\x96\x8d" +
	"\xe5/\xf7\xc0\xc6\xba\x89t\x85\x02\xc2\xbc0\xa3\x1a\xec" +
	";\x04d1\xed\xd5\xb6\x16\xdc\x0ca\x0a\x06\x09\x187" +
	"\xdd\x80\x95\xf5\xb9n\xde\x15\xdc\x92\xd8\x8b*\xda\x81\x9e" +
	"\x1cT\xb1\xeeY\xc7\xfak5\xd8\x0frO&wW" +
	"\xd3\x00\x09\x80 3\xe9L\xc1}\xacT$\xd4\xaf\x9f" +
	"\xfd\x17\x00\x00\xff\xff\xa9\xd5\xe3K"

func init() {
	schemas.Register(schema_8fcd13516850d142,
		0x934ba70bfd144ebd,
		0x9359e1b2db37dbbb,
		0x9d0135027d04861e,
		0xa12cfa420c9ad0ca,
		0xa7ad0c62a234c68b,
		0xabf979c3f68dae4b,
		0xbd56ceeaf8c65140,
		0xc88dfa6be7a1d091,
		0xde94294dfb441b76)
}