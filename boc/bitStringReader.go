package boc

import (
	"errors"
	"fmt"
	"github.com/startfellows/tongo"
	"math"
	"math/big"
)

var ErrNotEnoughBits = errors.New("not enough bits")

type BitStringReader struct {
	buf    []byte
	len    int
	cursor int
}

func NewBitStringReader(bitString *BitString) BitStringReader {
	var reader = BitStringReader{
		buf:    bitString.Buffer(),
		len:    bitString.len,
		cursor: 0,
	}
	return reader
}

func (s *BitStringReader) available() int {
	return s.len - s.cursor
}

func (s *BitStringReader) getBit(n int) bool {
	return (s.buf[n/8] & (1 << (7 - (n % 8)))) > 0
}

func (s *BitStringReader) readBit() bool {
	var bit = s.getBit(s.cursor)
	s.cursor++
	return bit
}

func (s *BitStringReader) Skip(n int) error {
	if s.available() < n {
		return ErrNotEnoughBits
	}
	s.cursor += n
	return nil
}

func (s *BitStringReader) ReadBit() (bool, error) {
	if s.available() < 1 {
		return false, ErrNotEnoughBits
	}
	var bit = s.getBit(s.cursor)
	s.cursor++
	return bit, nil
}

func (s *BitStringReader) ReadBigUint(bitLen int) (*big.Int, error) {
	if s.available() < bitLen {
		return nil, ErrNotEnoughBits
	}
	if bitLen == 0 {
		return big.NewInt(0), nil
	}
	var res = ""
	for i := 0; i < bitLen; i++ {
		if s.readBit() {
			res += "1"
		} else {
			res += "0"
		}
	}
	var num = big.NewInt(0)
	num.SetString(res, 2)
	return num, nil
}

func (s *BitStringReader) ReadBigInt(bitLen int) (*big.Int, error) {
	if s.available() < bitLen {
		return nil, ErrNotEnoughBits
	}
	if bitLen == 0 {
		return big.NewInt(0), nil
	}
	if bitLen == 1 {
		if s.readBit() {
			return big.NewInt(-1), nil
		}
		return big.NewInt(0), nil
	}
	if s.readBit() {
		var base, _ = s.ReadBigUint(bitLen - 1)
		var b = big.NewInt(2)
		var nb = b.Exp(b, big.NewInt(int64(bitLen-1)), nil)
		return base.Sub(base, nb), nil
	}
	return s.ReadBigUint(bitLen - 1)
}

func (s *BitStringReader) ReadUint(bitLen int) (uint, error) {
	if s.available() < bitLen {
		return 0, ErrNotEnoughBits
	}
	if bitLen == 0 {
		return 0, nil
	}
	var res uint = 0
	for i := bitLen - 1; i >= 0; i-- {
		if s.readBit() {
			res |= 1 << i
		}
	}
	return res, nil
}

func (s *BitStringReader) ReadInt(bitLen int) (int, error) {
	if s.available() < bitLen {
		return 0, ErrNotEnoughBits
	}
	if bitLen == 0 {
		return 0, nil
	}
	if bitLen == 1 {
		if s.readBit() {
			return -1, nil
		}
		return 0, nil
	}
	if s.readBit() {
		base, err := s.ReadUint(bitLen - 1)
		if err != nil {
			return 0, err
		}
		return int(base - uint(math.Pow(2, float64(bitLen-1)))), nil
	}
	res, err := s.ReadUint(bitLen - 1)
	if err != nil {
		return 0, err
	}
	return int(res), nil
}

// ReadGrams
// TL-B: nanograms$_ amount:(VarUInteger 16) = Grams;
func (s *BitStringReader) ReadGrams() (uint64, error) {
	grams, err := s.ReadVarUint(16)
	if err != nil {
		return 0, err
	}
	if !grams.IsUint64() {
		return 0, fmt.Errorf("grams uint64 overflow")
	}
	return grams.Uint64(), nil
}

func (s *BitStringReader) ReadByte() (byte, error) {
	res, err := s.ReadUint(8)
	if err != nil {
		return 0, err
	}
	return byte(res), nil
}

func (s *BitStringReader) ReadBytes(size int) ([]byte, error) {
	if s.available() < size*8 {
		return nil, ErrNotEnoughBits
	}
	res := make([]byte, size)
	for i := 0; i < size; i++ {
		b, err := s.ReadUint(8)
		if err != nil {
			return nil, err
		}
		res[i] = byte(b)
	}
	return res, nil
}

func (s *BitStringReader) ReadAddress() (*tongo.AccountID, error) {
	prefix, err := s.ReadUint(2)
	if err != nil {
		return nil, err
	}
	if prefix == 0 { // adr_none prefix
		return nil, nil
	}
	if prefix != 2 { // not adr_std prefix
		return nil, fmt.Errorf("not std address")
	}
	maybe, err := s.ReadBit()
	if err != nil {
		return nil, err
	}
	if maybe == true {
		return nil, fmt.Errorf("anycast not being processed") //TODO: add anycast processing
	}
	workchain, err := s.ReadInt(8)
	if err != nil {
		return nil, err
	}
	addr, err := s.ReadBytes(32)
	if err != nil {
		return nil, err
	}
	var address tongo.AccountID
	address.Workchain = int32(workchain)
	address.Address = addr
	return &address, nil
}

// ReadVarUint
// TL-B: var_uint$_ {n:#} len:(#< n) value:(uint (len * 8)) = VarUInteger n;
func (s *BitStringReader) ReadVarUint(byteLen int) (*big.Int, error) {
	if byteLen < 2 {
		return nil, fmt.Errorf("invalid varuint length")
	}
	lenBits := int(math.Ceil(math.Log2(float64(byteLen))))
	uintLen, err := s.ReadUint(lenBits)
	if err != nil {
		return nil, err
	}
	value, err := s.ReadBigUint(int(uintLen) * 8)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// ReadStorageUsed
// TL-B: storage_used$_ cells:(VarUInteger 7) bits:(VarUInteger 7) public_cells:(VarUInteger 7) = StorageUsed;
func (s *BitStringReader) ReadStorageUsed() error {
	_, err := s.ReadVarUint(7) // cells
	if err != nil {
		return err
	}
	//fmt.Printf("Cells: %v\n", cells)
	_, err = s.ReadVarUint(7) // bits
	if err != nil {
		return err
	}
	//fmt.Printf("Bits: %v\n", bits)
	_, err = s.ReadVarUint(7) // publicCells
	if err != nil {
		return err
	}
	//fmt.Printf("Public cells: %v\n", publicCells)
	return nil
}

// ReadStorageInfo
// TL-B: storage_info$_ used:StorageUsed last_paid:uint32 due_payment:(Maybe Grams) = StorageInfo;
func (s *BitStringReader) ReadStorageInfo() error {
	err := s.ReadStorageUsed()
	if err != nil {
		return err
	}
	_, err = s.ReadUint(32) //lastPaid
	//fmt.Printf("Last paid: %v\n", lastPaid)
	maybe, err := s.ReadMaybe()
	if err != nil {
		return err
	}
	if maybe {
		_, err = s.ReadGrams() // duePayment
		if err != nil {
			return err
		}
		//fmt.Printf("Due payment: %v\n", duePayment.String())
	}
	return nil
}

// ReadMaybe
// TL-B:
// nothing$0 {X:Type} = Maybe X;
// just$1 {X:Type} value:X = Maybe X;
func (s *BitStringReader) ReadMaybe() (bool, error) {
	return s.ReadBit()
}

// ReadAccountStorage
// TL-B: account_storage$_ last_trans_lt:uint64 balance:CurrencyCollection state:AccountState = AccountStorage;
func (s *BitStringReader) ReadAccountStorage() (RawAccount, error) {
	lastTransLt, err := s.ReadUint(64)
	if err != nil {
		return RawAccount{}, err
	}
	balance, err := s.ReadCurrencyCollection()
	if err != nil {
		return RawAccount{}, err
	}
	state, err := s.ReadAccountState()
	if err != nil {
		return RawAccount{}, err
	}
	return RawAccount{RawAccountState: state, Balance: balance, LastTransactionLt: uint64(lastTransLt)}, nil
}

// ReadCurrencyCollection
// TL-B: currencies$_ grams:Grams other:ExtraCurrencyCollection = CurrencyCollection;
func (s *BitStringReader) ReadCurrencyCollection() (grams uint64, err error) {
	grams, err = s.ReadGrams()
	if err != nil {
		return 0, err
	}
	//fmt.Printf("Grams: %v\n", grams.String())
	err = s.ReadExtraCurrencyCollection()
	if err != nil {
		return 0, err
	}
	return grams, nil
}

// ReadExtraCurrencyCollection
// TL-B: extra_currencies$_ dict:(HashmapE 32 (VarUInteger 32)) = ExtraCurrencyCollection;
func (s *BitStringReader) ReadExtraCurrencyCollection() error {
	// TODO: implement
	err := s.ReadHashmapE(0)
	return err
}

type RawAccount struct {
	RawAccountState
	Balance           uint64
	LastTransactionLt uint64
}

// ReadAccount
// TL-B:
// account_none$0 = Account;
// account$1 addr:MsgAddressInt storage_stat:StorageInfo storage:AccountStorage = Account;
func (s *BitStringReader) ReadAccount() (RawAccount, error) {
	// TODO: implement
	tag, err := s.ReadBit()
	if err != nil {
		return RawAccount{}, err
	}
	if tag == false {
		var account RawAccount
		account.Status = tongo.AccountNone
		return account, nil
	}
	_, err = s.ReadAddress() //addr
	if err != nil {
		return RawAccount{}, err
	}
	err = s.ReadStorageInfo()
	if err != nil {
		return RawAccount{}, err
	}
	return s.ReadAccountStorage()
}

// ReadHashmapE
// TODO: replace with CellReader
// TL-B:
// hme_empty$0 {n:#} {X:Type} = HashmapE n X;
// hme_root$1 {n:#} {X:Type} root:^(Hashmap n X) = HashmapE n X;
func (s *BitStringReader) ReadHashmapE(len int) error {
	// TODO: implement
	tag, err := s.ReadBit()
	if err != nil {
		return err
	}
	if tag == false {
		//fmt.Printf("HashmapE: empty\n")
		return nil
	}
	return fmt.Errorf("hashmapE not empty. not implemented")
}

type RawAccountState struct {
	Status     tongo.AccountStatus
	CodeFlag   bool
	DataFlag   bool
	FrozenHash [32]byte
}

// ReadAccountState
// TL-B:
// account_uninit$00 = AccountState;
// account_active$1 _:StateInit = AccountState;
// account_frozen$01 state_hash:bits256 = AccountState;
func (s *BitStringReader) ReadAccountState() (RawAccountState, error) {
	tag, err := s.ReadBit()
	if err != nil {
		return RawAccountState{}, err
	}
	var state RawAccountState
	if tag == true {
		codeFlag, dataFlag, err := s.ReadStateInit()
		if err != nil {
			return RawAccountState{}, err
		}
		state.Status = tongo.AccountActive
		state.CodeFlag = codeFlag
		state.DataFlag = dataFlag
		return state, nil
	}
	tag, err = s.ReadBit()
	if err != nil {
		return RawAccountState{}, err
	}
	if tag == false {
		state.Status = tongo.AccountUninit
		return state, err
	}
	stateHash, err := s.ReadBytes(32)
	if err != nil {
		return RawAccountState{}, err
	}
	state.Status = tongo.AccountFrozen
	copy(state.FrozenHash[:], stateHash[:])
	return state, nil
}

// ReadStateInit
// TL-B: _ split_depth:(Maybe (## 5)) special:(Maybe TickTock) code:(Maybe ^Cell) data:(Maybe ^Cell) library:(HashmapE 256 SimpleLib) = StateInit;
func (s *BitStringReader) ReadStateInit() (codeFlag bool, dataFlag bool, err error) {
	// TODO: implement
	maybe, err := s.ReadMaybe()
	if err != nil {
		return false, false, err
	}
	if maybe {
		return false, false, fmt.Errorf("splitDepth reading not implemented")
	}
	maybe, err = s.ReadMaybe()
	if err != nil {
		return false, false, err
	}
	if maybe {
		err = s.ReadTickTock()
		if err != nil {
			return false, false, err
		}
	}
	maybe, err = s.ReadMaybe()
	if err != nil {
		return false, false, err
	}
	if maybe {
		codeFlag = true
	}
	maybe, err = s.ReadMaybe()
	if err != nil {
		return false, false, err
	}
	if maybe {
		dataFlag = true
	}
	err = s.ReadHashmapE(0)
	if err != nil {
		return false, false, err
	}
	return codeFlag, dataFlag, nil
}

// ReadTickTock
// TL-B: tick_tock$_ tick:Bool tock:Bool = TickTock;
func (s *BitStringReader) ReadTickTock() error {
	_, err := s.ReadBit() // tick
	if err != nil {
		return err
	}
	_, err = s.ReadBit() // tock
	if err != nil {
		return err
	}
	// TODO: implement
	return nil
}

// Unary
// x == nil unary_zero
// x != nil unary_succ
type Unary struct {
	ConstructorName string
	X               *Unary
	n               uint32
}

// ReadUnary
// TL-B:
// unary_zero$0 = Unary ~0;
// unary_succ$1 {n:#} x:(Unary ~n) = Unary ~(n + 1);
func (s *BitStringReader) ReadUnary() (Unary, error) {
	unarySucc, err := s.ReadBit()
	if err != nil {
		return Unary{ConstructorName: "unary_zero", n: 0}, err
	}
	if !unarySucc {
		return Unary{}, nil
	}
	unary, err := s.ReadUnary()
	if err != nil {
		return Unary{}, err
	}
	return Unary{ConstructorName: "unary_succ", X: &unary, n: unary.n + 1}, nil
}

// HmLabel
// hml_short (len, s)
// hml_long (n, s)
// hml_same (v, n)
type HmLabel struct {
	ConstructorName string
	Len             *Unary
	S               *BitString
	N, M            uint32
	V               *bool
}

// ReadHmLabel
// TL-B:
// hml_short$0 {m:#} {n:#} len:(Unary ~n) {n <= m} s:(n * Bit) = HmLabel ~n m;
// hml_long$10 {m:#} n:(#<= m) s:(n * Bit) = HmLabel ~n m;
// hml_same$11 {m:#} v:Bit n:(#<= m) = HmLabel ~n m;
func (s *BitStringReader) ReadHmLabel(m uint32) (HmLabel, error) {
	notShort, err := s.ReadBit()
	if err != nil {
		return HmLabel{}, err
	}
	if !notShort {
		same, err := s.ReadBit()
		if err != nil {
			return HmLabel{}, err
		}
		if same {
			// decode hml_same
			v, err := s.ReadBit()
			if err != nil {
				return HmLabel{}, err
			}
			nLen := int(math.Ceil(math.Log2(float64(m + 1))))
			n, err := s.ReadUint(nLen)
			if err != nil {
				return HmLabel{}, err
			}
			return HmLabel{ConstructorName: "hml_same", V: &v, N: uint32(n), M: m}, nil
		}
		// decode hml_long
		nLen := int(math.Ceil(math.Log2(float64(m + 1))))
		n, err := s.ReadUint(nLen)
		if err != nil {
			return HmLabel{}, err
		}
		bits, err := s.ReadBits(n)
		if err != nil {
			return HmLabel{}, err
		}
		return HmLabel{ConstructorName: "hml_long", S: &bits, N: uint32(n), M: m}, nil
	}
	// decode hml_short
	ln, err := s.ReadUnary()
	if err != nil {
		return HmLabel{}, err
	}
	bits, err := s.ReadBits(uint(ln.n))
	if err != nil {
		return HmLabel{}, err
	}
	return HmLabel{ConstructorName: "hml_short", Len: &ln, S: &bits, N: ln.n, M: m}, nil
}

// ReadBits
// {n:#} (n * Bit)
func (s *BitStringReader) ReadBits(n uint) (BitString, error) {
	var bitString BitString
	for i := uint(0); i < n; i++ {
		bit, err := s.ReadBit()
		if err != nil {
			return BitString{}, err
		}
		err = bitString.WriteBit(bit)
		if err != nil {
			return BitString{}, err
		}
	}
	return bitString, nil
}

type HashmapNode struct {
	ConstructorName string
	NPlus           uint32
	Value           *any
	Left            *Hashmap
	Right           *Hashmap
}

// ReadHashmapNode
// TL-B:
// hmn_leaf#_ {X:Type} value:X = HashmapNode 0 X;
// hmn_fork#_ {n:#} {X:Type} left:^(Hashmap n X) right:^(Hashmap n X) = HashmapNode (n + 1) X;
func (c *CellReader) ReadHashmapNode(nPlus uint32, x any) (HashmapNode, error) {
	if nPlus == 0 {
		value, err := c.ReadAnyType(x)
		if err != nil {
			return HashmapNode{}, err
		}
		return HashmapNode{ConstructorName: "hmn_leaf", Value: value, NPlus: nPlus}, nil
	}
	n := nPlus - 1
	leftCell, err := c.getRef()
	if err != nil {
		return HashmapNode{}, err
	}
	left, err := leftCell.ReadHashmap(n, x)
	if err != nil {
		return HashmapNode{}, err
	}
	rightCell, err := c.getRef()
	if err != nil {
		return HashmapNode{}, err
	}
	right, err := rightCell.ReadHashmap(n, x)
	if err != nil {
		return HashmapNode{}, err
	}
	return HashmapNode{ConstructorName: "hmn_fork", NPlus: nPlus, Left: left, Right: right}, nil
}

type CellReader struct {
	Cell            Cell
	BitStringReader BitStringReader
}

type Hashmap struct {
	ConstructorName string
	Label           HmLabel
	Node            HashmapNode
	N               uint32
	X               any
}

// ReadHashmap
// TL-B:
// hm_edge#_ {n:#} {X:Type} {l:#} {m:#} label:(HmLabel ~l n) {n = (~m) + l} node:(HashmapNode m X) = Hashmap n X;
func (c *CellReader) ReadHashmap(n uint32, x any) (Hashmap, error) {
	label, err := c.ReadHmLabel(n)
	if err != nil {
		return Hashmap{}, err
	}
	m := n - label.n
	node, err := c.ReadHashmapNode(m, x)
	if err != nil {
		return Hashmap{}, err
	}
	return Hashmap{ConstructorName: "hm_edge", Label: label, Node: node, N: n, X: x}, nil
}

type HashmapE struct {
	ConstructorName string
	Root            *Hashmap
	N               uint32
	X               any
}

// ReadHashmapE
// TODO: replace with CellReader
// TL-B:
// hme_empty$0 {n:#} {X:Type} = HashmapE n X;
// hme_root$1 {n:#} {X:Type} root:^(Hashmap n X) = HashmapE n X;
func (c *CellReader) ReadHashmapE(n uint32, x any) (HashmapE, error) {
	root, err := c.BitStringReader.ReadBit()
	if err != nil {
		return HashmapE{}, err
	}
	if root {
		rootCell, err := c.getRef()
		if err != nil {
			return HashmapE{}, err
		}
		root, err := rootCell.ReadHashmap(n, x)
		if err != nil {
			return HashmapE{}, err
		}
		return HashmapE{ConstructorName: "hme_root", Root: root, N: n, X: x}, nil
	}
	return HashmapE{ConstructorName: "hme_empty", N: n, X: x}, nil
}
