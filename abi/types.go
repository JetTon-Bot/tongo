package abi

// Code autogenerated. DO NOT EDIT.

import (
	"github.com/tonkeeper/tongo/tlb"
)

type ClosingConfig struct {
	QuarantinDuration        uint32
	MisbehaviorFine          tlb.Grams
	ConditionalCloseDuration uint32
}

type ConditionalPayment struct {
	Amount    tlb.Grams
	Condition tlb.Any
}

type PaymentConfig struct {
	ExcessFee tlb.Grams
	DestA     tlb.MsgAddress
	DestB     tlb.MsgAddress
}

type QuarantinedState struct {
	StateA           SemiChannelBody
	StateB           SemiChannelBody
	QuarantineStarts uint32
	StateCommitedByA bool
}

type SemiChannel struct {
	Magic            tlb.Magic `tlb:"#43685374"`
	ChannelId        tlb.Uint128
	Data             SemiChannelBody
	CounterpartyData tlb.Maybe[tlb.Ref[SemiChannelBody]]
}

type SemiChannelBody struct {
	Seqno        uint64
	Sent         tlb.Grams
	Conditionals tlb.HashmapE[tlb.Uint32, ConditionalPayment]
}

type SignedSemiChannel struct {
	Signature tlb.Bits512
	State     SemiChannel
}

type Storage struct {
	BalanceA       tlb.Grams
	BalanceB       tlb.Grams
	KeyA           tlb.Uint256
	KeyB           tlb.Uint256
	ChannelId      tlb.Uint128
	Config         tlb.Ref[ClosingConfig]
	CommitedSeqnoA uint32
	CommitedSeqnoB uint32
	Quarantin      tlb.Maybe[tlb.Ref[QuarantinedState]]
	Payments       tlb.Ref[PaymentConfig]
}

type TorrentInfo struct {
	PieceSize      uint32
	FileSize       uint64
	RootHash       tlb.Uint256
	HeaderSize     uint64
	HeaderHash     tlb.Uint256
	MicrochunkHash tlb.Maybe[tlb.Uint256]
	Description    tlb.Text
}

type NftRoyaltyParams struct {
	Numerator   uint16
	Denominator uint16
	Destination tlb.MsgAddress
}

type TeleitemAuctionConfig struct {
	BeneficiarAddress tlb.MsgAddress
	InitialMinBid     tlb.Grams
	MaxBid            tlb.Grams
	MinBidStep        uint8
	MinExtendTime     uint32
	Duration          uint32
}

type TelemintData struct {
	Touched           bool
	SubwalletId       uint32
	PublicKey         tlb.Bits256
	CollectionContent tlb.Ref[tlb.Any]
	NftItemCode       tlb.Ref[tlb.Any]
	RoyaltyParams     tlb.Ref[NftRoyaltyParams]
}

type TelemintRestrictions struct {
	ForceSenderAddress   tlb.Maybe[tlb.MsgAddress]
	RewriteSenderAddress tlb.Maybe[tlb.MsgAddress]
}

type TelemintUnsignedDeploy struct {
	SubwalletId   uint32
	ValidSince    uint32
	ValidTill     uint32
	Username      tlb.FixedLengthText
	Content       tlb.Ref[tlb.Any]
	AuctionConfig tlb.Ref[TeleitemAuctionConfig]
	RoyaltyParams tlb.Maybe[tlb.Ref[NftRoyaltyParams]]
}

type TelemintUnsignedDeployV2 struct {
	SubwalletId   uint32
	ValidSince    uint32
	ValidTill     uint32
	TokenName     tlb.FixedLengthText
	Content       tlb.Ref[tlb.Any]
	AuctionConfig tlb.Ref[TeleitemAuctionConfig]
	RoyaltyParams tlb.Maybe[tlb.Ref[NftRoyaltyParams]]
	Restrictions  tlb.Maybe[tlb.Ref[TelemintRestrictions]]
}

type WhalesNominatorsMember struct {
	ProfitPerCoin      tlb.Int128
	Balance            tlb.Grams
	PendingWithdraw    tlb.Grams
	PendingWithdrawAll bool
	PendingDeposit     tlb.Grams
	MemberWithdraw     tlb.Grams
}

type WhalesNominatorsMembersList struct {
	List tlb.Hashmap[tlb.Bits256, WhalesNominatorsMember]
}

type AccountLists struct {
	List tlb.Hashmap[tlb.Bits256, tlb.Any]
}

type StonfiPayToParams struct {
	Amount0Out    tlb.VarUInteger16
	Token0Address tlb.MsgAddress
	Amount1Out    tlb.VarUInteger16
	Token1Address tlb.MsgAddress
}

type StonfiSwapAddrs struct {
	FromUser tlb.MsgAddress
}

type TextCommentMsgBody struct {
	Text tlb.Text
}

type ProveOwnershipMsgBody struct {
	QueryId        uint64
	Dest           tlb.MsgAddress
	ForwardPayload tlb.Ref[tlb.Any]
	WithContent    bool
}

type NftOwnershipAssignedMsgBody struct {
	QueryId        uint64
	PrevOwner      tlb.MsgAddress
	ForwardPayload tlb.EitherRef[tlb.Any]
}

type OwnershipProofMsgBody struct {
	QueryId   uint64
	ItemId    tlb.Uint256
	Owner     tlb.MsgAddress
	Data      tlb.Ref[tlb.Any]
	RevokedAt uint64
	Content   tlb.Maybe[tlb.Ref[tlb.Any]]
}

type ChallengeQuarantinedChannelStateMsgBody struct {
	ChallengedByA bool
	Signature     tlb.Bits512
	Tag           uint32
	ChannelId     tlb.Uint128
	SchA          tlb.Ref[SignedSemiChannel]
	SchB          tlb.Ref[SignedSemiChannel]
}

type TonstakePoolWithdrawalMsgBody struct {
	QueryId uint64
}

type SbtOwnerInfoMsgBody struct {
	QueryId   uint64
	ItemId    tlb.Uint256
	Initiator tlb.MsgAddress
	Owner     tlb.MsgAddress
	Data      tlb.Ref[tlb.Any]
	RevokedAt uint64
	Content   tlb.Maybe[tlb.Ref[tlb.Any]]
}

type InitPaymentChannelMsgBody struct {
	IsA       bool
	Signature tlb.Bits512
	Tag       uint32
	ChannelId tlb.Uint128
	BalanceA  tlb.Grams
	BalanceB  tlb.Grams
}

type JettonTransferMsgBody struct {
	QueryId             uint64
	Amount              tlb.VarUInteger16
	Destination         tlb.MsgAddress
	ResponseDestination tlb.MsgAddress
	CustomPayload       tlb.Maybe[tlb.Ref[tlb.Any]]
	ForwardTonAmount    tlb.VarUInteger16
	ForwardPayload      tlb.EitherRef[tlb.Any]
}

type OfferStorageContractMsgBody struct {
	QueryId uint64
}

type TonstakeNftInitMsgBody struct {
	QueryId uint64
	Owner   tlb.MsgAddress
	Amount  tlb.Grams
	Prev    tlb.MsgAddress
	Next    tlb.MsgAddress
}

type TonstakeControllerPoolHaltMsgBody struct {
	QueryId uint64
}

type WhalesNominatorsForceKickMsgBody struct {
	QueryId int64
}

type TonstakeControllerCreditMsgBody struct {
	QueryId uint64
	Amount  tlb.Grams
}

type JettonInternalTransferMsgBody struct {
	QueryId          uint64
	Amount           tlb.VarUInteger16
	From             tlb.MsgAddress
	ResponseAddress  tlb.MsgAddress
	ForwardTonAmount tlb.VarUInteger16
}

type WhalesNominatorsWithdrawUnownedResponseMsgBody struct {
	QueryId uint64
}

type SbtDestroyMsgBody struct {
	QueryId uint64
}

type StartUncooperativeChannelCloseMsgBody struct {
	SignedByA bool
	Signature tlb.Bits512
	Tag       uint32
	ChannelId tlb.Uint128
	SchA      tlb.Ref[SignedSemiChannel]
	SchB      tlb.Ref[SignedSemiChannel]
}

type EncryptedTextCommentMsgBody struct {
	CipherText tlb.Bytes
}

type WhalesNominatorsStakeWithdrawCompletedMsgBody struct {
	QueryId int64
}

type WhalesNominatorsWithdrawUnownedMsgBody struct {
	QueryId  uint64
	GasLimit tlb.Grams
}

type FinishUncooperativeChannelCloseMsgBody struct{}

type StonfiSwapMsgBody struct {
	QueryId       uint64
	ToAddress     tlb.MsgAddress
	SenderAddress tlb.MsgAddress
	JettonAmount  tlb.VarUInteger16
	MinOut        tlb.VarUInteger16
	HasRefAddress bool
	Addrs         tlb.Ref[StonfiSwapAddrs]
}

type TonstakeControllerPoolSendMessageMsgBody struct {
	QueryId uint64
	Mode    uint8
	Msg     tlb.Ref[tlb.Any]
}

type TeleitemDeployMsgBody struct {
	SenderAddress tlb.MsgAddress
	Bid           tlb.Grams
	Username      tlb.FixedLengthText
	Content       tlb.Ref[tlb.Any]
	AuctionConfig tlb.Ref[TeleitemAuctionConfig]
	RoyaltyParams tlb.Ref[NftRoyaltyParams]
}

type TonstakePoolSetGovernanceFeeMsgBody struct {
	QueryId       uint64
	GovernanceFee uint16
}

type GetStaticDataMsgBody struct {
	QueryId uint64
}

type TonstakeControllerValidatorWithdrawalMsgBody struct {
	QueryId uint64
	Amount  tlb.Grams
}

type TonstakePoolWithdrawMsgBody struct {
	QueryId         uint64
	JettonAmount    tlb.Grams
	FromAddress     tlb.MsgAddress
	ResponseAddress tlb.MsgAddress
}

type AuctionFillUpMsgBody struct {
	QueryId uint64
}

type TeleitemCancelAuctionMsgBody struct {
	QueryId int64
}

type ProofStorageMsgBody struct {
	QueryId       uint64
	FileDictProof tlb.Ref[tlb.Any]
}

type TelemintDeployMsgBody struct {
	Sig tlb.Bits512
	Msg TelemintUnsignedDeploy
}

type TelemintDeployV2MsgBody struct {
	Sig tlb.Bits512
	Msg TelemintUnsignedDeployV2
}

type StorageWithdrawMsgBody struct {
	QueryId uint64
}

type ElectorRecoverStakeRequestMsgBody struct {
	QueryId uint64
}

type TonstakePoolDepositMsgBody struct {
	QueryId uint64
}

type TeleitemStartAuctionMsgBody struct {
	QueryId       int64
	AuctionConfig tlb.Ref[TeleitemAuctionConfig]
}

type TonstakePoolTouchMsgBody struct {
	QueryId uint64
}

type ElectorNewStakeMsgBody struct {
	QueryId         uint64
	ValidatorPubkey tlb.Bits256
	StakeAt         uint32
	MaxFactor       uint32
	AdnlAddr        tlb.Bits256
	Signature       tlb.Ref[tlb.Bits512]
}

type UpdatePubkeyMsgBody struct {
	QueryId   uint64
	NewPubkey tlb.Bits256
}

type UpdateStorageParamsMsgBody struct {
	QueryId            uint64
	AcceptNewContracts bool
	RatePerMbDay       tlb.Grams
	MaxSpan            uint32
	MinimalFileSize    uint64
	MaximalFileSize    uint64
}

type TonstakeImanagerOperationFeeMsgBody struct {
	QueryId uint64
}

type ChannelCooperativeCloseMsgBody struct {
	SigA      tlb.Ref[tlb.Bits512]
	SigB      tlb.Ref[tlb.Bits512]
	Tag       uint32
	ChannelId tlb.Uint128
	BalanceA  tlb.Grams
	BalanceB  tlb.Grams
	SeqnoA    uint64
	SeqnoB    uint64
}

type TonstakeControllerReturnAvailableFundsMsgBody struct {
	QueryId uint64
}

type JettonBurnMsgBody struct {
	QueryId             uint64
	Amount              tlb.VarUInteger16
	ResponseDestination tlb.MsgAddress
	CustomPayload       tlb.Maybe[tlb.Ref[tlb.Any]]
}

type TonstakePoolSetRolesMsgBody struct {
	QueryId         uint64
	Governor        tlb.Maybe[tlb.MsgAddress]
	InterestManager tlb.Maybe[tlb.MsgAddress]
	Halter          tlb.Maybe[tlb.MsgAddress]
}

type NftTransferMsgBody struct {
	QueryId             uint64
	NewOwner            tlb.MsgAddress
	ResponseDestination tlb.MsgAddress
	CustomPayload       tlb.Maybe[tlb.Ref[tlb.Any]]
	ForwardAmount       tlb.VarUInteger16
	ForwardPayload      tlb.EitherRef[tlb.Any]
}

type TonstakeControllerSendRequestLoanMsgBody struct {
	QueryId    uint64
	MinLoan    tlb.Grams
	MaxLoan    tlb.Grams
	MaxInterst uint16
}

type WalletPluginDestructMsgBody struct{}

type SettleChannelConditionalsMsgBody struct {
	FromA                bool
	Signature            tlb.Bits512
	Tag                  uint32
	ChannelId            tlb.Uint128
	ConditionalsToSettle tlb.HashmapE[tlb.Uint32, tlb.Any]
}

type TopUpChannelBalanceMsgBody struct {
	AddA tlb.Grams
	AddB tlb.Grams
}

type GetRoyaltyParamsMsgBody struct {
	QueryId uint64
}

type SbtRevokeMsgBody struct {
	QueryId uint64
}

type PaymentRequestMsgBody struct {
	QueryId uint64
	Amount  tlb.CurrencyCollection
}

type TonstakeControllerPoolUnhaltMsgBody struct {
	QueryId uint64
}

type JettonNotifyMsgBody struct {
	QueryId        uint64
	Amount         tlb.VarUInteger16
	Sender         tlb.MsgAddress
	ForwardPayload tlb.EitherRef[tlb.Any]
}

type SubscriptionPaymentMsgBody struct{}

type WhalesNominatorsStakeWithdrawDelayedMsgBody struct {
	QueryId int64
}

type ChannelCooperativeCommitMsgBody struct {
	SigA      tlb.Ref[tlb.Bits512]
	SigB      tlb.Ref[tlb.Bits512]
	Tag       uint32
	ChannelId tlb.Uint128
	SeqnoA    uint64
	SeqnoB    uint64
}

type TonstakeControllerPoolSetSudoerMsgBody struct {
	QueryId uint64
	Sudoer  tlb.MsgAddress
}

type CloseStorageContractMsgBody struct {
	QueryId uint64
}

type AcceptStorageContractMsgBody struct {
	QueryId uint64
}

type TonstakeControllerApproveMsgBody struct {
	QueryId uint64
}

type WhalesNominatorsDepositMsgBody struct {
	QueryId int64
	Gas     tlb.Grams
}

type JettonBurnNotificationMsgBody struct {
	QueryId             uint64
	Amount              tlb.VarUInteger16
	Sender              tlb.MsgAddress
	ResponseDestination tlb.MsgAddress
}

type ReportStaticDataMsgBody struct {
	QueryId    uint64
	Index      tlb.Uint256
	Collection tlb.MsgAddress
}

type TonstakeControllerWithdrawValidatorMsgBody struct {
	QueryId uint64
	Value   tlb.Grams
}

type TonstakeControllerPoolUpgradeMsgBody struct {
	QueryId      uint64
	Data         tlb.Maybe[tlb.Ref[tlb.Any]]
	Code         tlb.Maybe[tlb.Ref[tlb.Any]]
	AfterUpgrade tlb.Maybe[tlb.Ref[tlb.Any]]
}

type TonstakePoolPrepareGovernanceMigrationMsgBody struct {
	QueryId             uint64
	GovernorUpdateAfter tlb.Uint48
}

type WhalesNominatorsAcceptStakeMsgBody struct {
	QueryId uint64
	Members tlb.Any
}

type TonstakePoolSetDepositSettingsMsgBody struct {
	QueryId                      uint64
	OptimisticDepositWithdrawals bool
	DepositsOpen                 bool
}

type WhalesNominatorsAcceptWithdrawsMsgBody struct {
	QueryId uint64
	Members tlb.Any
}

type WhalesNominatorsSendStakeMsgBody struct {
	QueryId         uint64
	GasLimit        tlb.Grams
	Stake           tlb.Grams
	ValidatorPubkey tlb.Bits256
	StakeAt         uint32
	MaxFactor       uint32
	AdnlAddr        tlb.Bits256
	Signature       tlb.Ref[tlb.Bits512]
}

type TeleitemOkMsgBody struct {
	QueryId int64
}

type TeleitemReturnBidMsgBody struct {
	CurLt int64
}

type ReportRoyaltyParamsMsgBody struct {
	QueryId     uint64
	Numerator   uint16
	Denominator uint16
	Destination tlb.MsgAddress
}

type StorageRewardWithdrawalMsgBody struct {
	QueryId uint64
}

type TonstakeImanagerRequestNotificationMsgBody struct {
	QueryId     uint64
	MinLoan     tlb.Grams
	MaxLoan     tlb.Grams
	MaxInterest uint16
}

type TonstakePoolDeployControllerMsgBody struct {
	ControllerId uint32
	QueryId      uint64
}

type StorageContractTerminatedMsgBody struct {
	CurLt       uint64
	TorrentHash tlb.Bits256
}

type TonstakeImanagerStatsMsgBody struct {
	QueryId      uint64
	Borrowed     tlb.Grams
	Expected     tlb.Grams
	Returned     tlb.Grams
	ProfitSign   tlb.Int1
	Profit       tlb.Grams
	TotalBalance tlb.Grams
}

type TonstakeImanagerSetInterestMsgBody struct {
	QueryId      uint64
	InterestRate uint16
}

type SbtRequestOwnerMsgBody struct {
	QueryId        uint64
	Dest           tlb.MsgAddress
	ForwardPayload tlb.Ref[tlb.Any]
	WithContent    bool
}

type TonstakeControllerTopUpMsgBody struct {
	QueryId uint64
}

type StorageContractConfirmedMsgBody struct {
	CurLt       uint64
	TorrentHash tlb.Bits256
}

type ExcessMsgBody struct {
	QueryId uint64
}

type WhalesNominatorsWithdrawMsgBody struct {
	QueryId int64
	Gas     tlb.Grams
	Amount  tlb.Grams
}

type ChannelClosedMsgBody struct {
	ChannelId tlb.Uint128
}

type TonstakePoolLoanRepaymentMsgBody struct {
	QueryId uint64
}

type WalletPluginDestructResponseMsgBody struct{}

type DeployStorageContractMsgBody struct {
	QueryId         uint64
	Info            tlb.Ref[TorrentInfo]
	MerkleHash      tlb.Bits256
	ExpectedRate    tlb.Grams
	ExpectedMaxSpan uint32
}

type TonstakePoolRequestLoanMsgBody struct {
	QueryId      uint64
	MinLoan      tlb.Grams
	MaxLoan      tlb.Grams
	MaxInterest  uint16
	ControllerId uint32
	Validator    tlb.MsgAddress
}

type TonstakeControllerDisapproveMsgBody struct {
	QueryId uint64
}

type TonstakeControllerRecoverStakeMsgBody struct {
	QueryId uint64
}

type TonstakeNftBurnNotificationMsgBody struct {
	QueryId uint64
	Amount  tlb.Grams
	Owner   tlb.MsgAddress
	Index   uint64
}

type TonstakeControllerReturnUnusedLoanMsgBody struct {
	QueryId uint64
}

type PaymentRequestResponseMsgBody struct{}

type TonstakeControllerUpdateValidatorHashMsgBody struct {
	QueryId uint64
}

type TonstakeNftBurnMsgBody struct {
	QueryId uint64
}

type ElectorNewStakeConfirmationMsgBody struct {
	QueryId uint64
}

type StonfiPaymentRequestMsgBody struct {
	QueryId  uint64
	Owner    tlb.MsgAddress
	ExitCode uint32
	Params   tlb.EitherRef[StonfiPayToParams]
}

type ElectorRecoverStakeResponseMsgBody struct {
	QueryId uint64
}

type BounceMsgBody struct {
	Payload tlb.Any
}
