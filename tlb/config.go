package tlb

// Code autogenerated. DO NOT EDIT.

type ValidatorSet struct {
	SumType
	Validators struct {
		UtimeSince uint32
		UtimeUntil uint32
		Total      uint16
		Main       uint16
		List       Hashmap[Uint16, ValidatorDescr]
	} `tlbSumType:"#11"`
	ValidatorsExt struct {
		UtimeSince  uint32
		UtimeUntil  uint32
		Total       uint16
		Main        uint16
		TotalWeight uint64
		List        HashmapE[Uint16, ValidatorDescr]
	} `tlbSumType:"#12"`
}

type ConfigParam0 struct {
	ConfigAddr Bits256
}

type ConfigParam1 struct {
	ElectorAddr Bits256
}

type ConfigParam2 struct {
	MinterAddr Bits256
}

type ConfigParam3 struct {
	FeeCollectorAddr Bits256
}

type ConfigParam4 struct {
	DnsRootAddr Bits256
}

type BurningConfig struct {
	Magic         Magic    `tlb:"#01"`
	BlackholeAddr *Bits256 `tlb:"maybe"`
	FeeBurnNom    uint32
	FeeBurnDenom  uint32
}

type ConfigParam5 struct {
	BurningConfig BurningConfig
}

type ConfigParam6 struct {
	MintNewPrice Grams
	MintAddPrice Grams
}

type ConfigParam7 struct {
	ToMint ExtraCurrencyCollection
}

type ConfigParam8 struct {
	GlobalVersion GlobalVersion
}

type ConfigParam9 struct {
	MandatoryParams Hashmap[Int32, struct{}]
}

type ConfigParam10 struct {
	CriticalParams Hashmap[Int32, struct{}]
}

type ConfigProposalSetup struct {
	Magic        Magic `tlb:"#36"`
	MinTotRounds uint8
	MaxTotRounds uint8
	MinWins      uint8
	MaxLosses    uint8
	MinStoreSec  uint32
	MaxStoreSec  uint32
	BitPrice     uint32
	CellPrice    uint32
}

type ConfigVotingSetup struct {
	Magic          Magic               `tlb:"#91"`
	NormalParams   ConfigProposalSetup `tlb:"^"`
	CriticalParams ConfigProposalSetup `tlb:"^"`
}

type ConfigParam11 struct {
	ConfigVotingSetup ConfigVotingSetup
}

type ConfigProposal struct {
	Magic       Magic `tlb:"#f3"`
	ParamId     int32
	ParamValue  *Any     `tlb:"maybe^"`
	IfHashEqual *Uint256 `tlb:"maybe"`
}

type ConfigProposalStatus struct {
	Magic           Magic `tlb:"#ce"`
	Expires         uint32
	Proposal        ConfigProposal `tlb:"^"`
	IsCritical      bool
	Voters          HashmapE[Uint16, struct{}]
	RemainingWeight int64
	ValidatorSetId  Uint256
	RoundsRemaining uint8
	Wins            uint8
	Losses          uint8
}

type WorkchainFormat1 struct {
	Magic     Magic `tlb:"#1"`
	VmVersion int32
	VmMode    uint64
}

type WorkchainFormat0 struct {
	Magic           Magic `tlb:"#0"`
	MinAddrLen      Uint12
	MaxAddrLen      Uint12
	AddrLenStep     Uint12
	WorkchainTypeId uint32
}

type WcSplitMergeTimings struct {
	Magic                 Magic `tlb:"#0"`
	SplitMergeDelay       uint32
	SplitMergeInterval    uint32
	MinSplitMergeInterval uint32
	MaxSplitMergeDelay    uint32
}

type WorkchainDescr struct {
	SumType
	Workchain struct {
		EnabledSince      uint32
		ActualMinSplit    uint8
		MinSplit          uint8
		MaxSplit          uint8
		Basic             Uint1
		Active            bool
		AcceptMsgs        bool
		Flags             Uint13
		ZerostateRootHash Bits256
		ZerostateFileHash Bits256
		Version           uint32
	} `tlbSumType:"#a6"`
	WorkchainV2 struct {
		EnabledSince      uint32
		ActualMinSplit    uint8
		MinSplit          uint8
		MaxSplit          uint8
		Basic             Uint1
		Active            bool
		AcceptMsgs        bool
		Flags             Uint13
		ZerostateRootHash Bits256
		ZerostateFileHash Bits256
		Version           uint32
	} `tlbSumType:"#a7"`
}

type ConfigParam12 struct {
	Workchains HashmapE[Uint32, WorkchainDescr]
}

type ComplaintPricing struct {
	Magic     Magic `tlb:"#1a"`
	Deposit   Grams
	BitPrice  Grams
	CellPrice Grams
}

type ConfigParam13 struct {
	ComplaintPricing ComplaintPricing
}

type BlockCreateFees struct {
	Magic               Magic `tlb:"#6b"`
	MasterchainBlockFee Grams
	BasechainBlockFee   Grams
}

type ConfigParam14 struct {
	BlockCreateFees BlockCreateFees
}

type ConfigParam15 struct {
	ValidatorsElectedFor uint32
	ElectionsStartBefore uint32
	ElectionsEndBefore   uint32
	StakeHeldFor         uint32
}

type ConfigParam16 struct {
	MaxValidators     uint16
	MaxMainValidators uint16
	MinValidators     uint16
}

type ConfigParam17 struct {
	MinStake       Grams
	MaxStake       Grams
	MinTotalStake  Grams
	MaxStakeFactor uint32
}

type StoragePrices struct {
	Magic         Magic `tlb:"#cc"`
	UtimeSince    uint32
	BitPricePs    uint64
	CellPricePs   uint64
	McBitPricePs  uint64
	McCellPricePs uint64
}

type ConfigParam18 struct {
	Value Hashmap[Uint32, StoragePrices]
}

type GasLimitsPrices struct {
	SumType
	GasPrices struct {
		GasPrice       uint64
		GasLimit       uint64
		GasCredit      uint64
		BlockGasLimit  uint64
		FreezeDueLimit uint64
		DeleteDueLimit uint64
	} `tlbSumType:"#dd"`
	GasPricesExt struct {
		GasPrice        uint64
		GasLimit        uint64
		SpecialGasLimit uint64
		GasCredit       uint64
		BlockGasLimit   uint64
		FreezeDueLimit  uint64
		DeleteDueLimit  uint64
	} `tlbSumType:"#de"`
	GasFlatPfx struct {
		FlatGasLimit uint64
		FlatGasPrice uint64
		Other        *GasLimitsPrices
	} `tlbSumType:"#d1"`
}

type ConfigParam20 struct {
	GasLimitsPrices GasLimitsPrices
}

type ConfigParam21 struct {
	GasLimitsPrices GasLimitsPrices
}

type ParamLimits struct {
	Magic     Magic `tlb:"#c3"`
	Underload uint32
	SoftLimit uint32
	HardLimit uint32
}

type BlockLimits struct {
	Magic   Magic `tlb:"#5d"`
	Bytes   ParamLimits
	Gas     ParamLimits
	LtDelta ParamLimits
}

type ConfigParam22 struct {
	BlockLimits BlockLimits
}

type ConfigParam23 struct {
	BlockLimits BlockLimits
}

type MsgForwardPrices struct {
	Magic          Magic `tlb:"#ea"`
	LumpPrice      uint64
	BitPrice       uint64
	CellPrice      uint64
	IhrPriceFactor uint32
	FirstFrac      uint16
	NextFrac       uint16
}

type ConfigParam24 struct {
	MsgForwardPrices MsgForwardPrices
}

type ConfigParam25 struct {
	MsgForwardPrices MsgForwardPrices
}

type CatchainConfig struct {
	SumType
	CatchainConfig struct {
		McCatchainLifetime      uint32
		ShardCatchainLifetime   uint32
		ShardValidatorsLifetime uint32
		ShardValidatorsNum      uint32
	} `tlbSumType:"#c1"`
	CatchainConfigNew struct {
		Flags                   Uint7
		ShuffleMcValidators     bool
		McCatchainLifetime      uint32
		ShardCatchainLifetime   uint32
		ShardValidatorsLifetime uint32
		ShardValidatorsNum      uint32
	} `tlbSumType:"#c2"`
}

type ConsensusConfig struct {
	SumType
	ConsensusConfig struct {
		RoundCandidates      uint32
		NextCandidateDelayMs uint32
		ConsensusTimeoutMs   uint32
		FastAttempts         uint32
		AttemptDuration      uint32
		CatchainMaxDeps      uint32
		MaxBlockBytes        uint32
		MaxCollatedBytes     uint32
	} `tlbSumType:"#d6"`
	ConsensusConfigNew struct {
		Flags                Uint7
		NewCatchainIds       bool
		RoundCandidates      uint8
		NextCandidateDelayMs uint32
		ConsensusTimeoutMs   uint32
		FastAttempts         uint32
		AttemptDuration      uint32
		CatchainMaxDeps      uint32
		MaxBlockBytes        uint32
		MaxCollatedBytes     uint32
	} `tlbSumType:"#d7"`
	ConsensusConfigV3 struct {
		Flags                Uint7
		NewCatchainIds       bool
		RoundCandidates      uint8
		NextCandidateDelayMs uint32
		ConsensusTimeoutMs   uint32
		FastAttempts         uint32
		AttemptDuration      uint32
		CatchainMaxDeps      uint32
		MaxBlockBytes        uint32
		MaxCollatedBytes     uint32
		ProtoVersion         uint16
	} `tlbSumType:"#d8"`
	ConsensusConfigV4 struct {
		Flags                  Uint7
		NewCatchainIds         bool
		RoundCandidates        uint8
		NextCandidateDelayMs   uint32
		ConsensusTimeoutMs     uint32
		FastAttempts           uint32
		AttemptDuration        uint32
		CatchainMaxDeps        uint32
		MaxBlockBytes          uint32
		MaxCollatedBytes       uint32
		ProtoVersion           uint16
		CatchainMaxBlocksCoeff uint32
	} `tlbSumType:"#d9"`
}

type ConfigParam28 struct {
	CatchainConfig CatchainConfig
}

type ConfigParam29 struct {
	ConsensusConfig ConsensusConfig
}

type ConfigParam31 struct {
	FundamentalSmcAddr HashmapE[Bits256, struct{}]
}

type ConfigParam32 struct {
	PrevValidators ValidatorSet
}

type ConfigParam33 struct {
	PrevTempValidators ValidatorSet
}

type ConfigParam34 struct {
	CurValidators ValidatorSet
}

type ConfigParam35 struct {
	CurTempValidators ValidatorSet
}

type ConfigParam36 struct {
	NextValidators ValidatorSet
}

type ConfigParam37 struct {
	NextTempValidators ValidatorSet
}

type ValidatorTempKey struct {
	Magic         Magic `tlb:"#3"`
	AdnlAddr      Bits256
	TempPublicKey SigPubKey
	Seqno         uint32
	ValidUntil    uint32
}

type ValidatorSignedTempKey struct {
	Magic     Magic            `tlb:"#4"`
	Key       ValidatorTempKey `tlb:"^"`
	Signature CryptoSignature
}

type ConfigParam39 struct {
	Value HashmapE[Bits256, ValidatorSignedTempKey]
}

type MisbehaviourPunishmentConfig struct {
	Magic                    Magic `tlb:"#01"`
	DefaultFlatFine          Grams
	DefaultProportionalFine  uint32
	SeverityFlatMult         uint16
	SeverityProportionalMult uint16
	UnpunishableInterval     uint16
	LongInterval             uint16
	LongFlatMult             uint16
	LongProportionalMult     uint16
	MediumInterval           uint16
	MediumFlatMult           uint16
	MediumProportionalMult   uint16
}

type ConfigParam40 struct {
	MisbehaviourPunishmentConfig MisbehaviourPunishmentConfig
}

type SizeLimitsConfig struct {
	SumType
	SizeLimitsConfig struct {
		MaxMsgBits      uint32
		MaxMsgCells     uint32
		MaxLibraryCells uint32
		MaxVmDataDepth  uint16
		MaxExtMsgSize   uint32
		MaxExtMsgDepth  uint16
	} `tlbSumType:"#01"`
	SizeLimitsConfigV2 struct {
		MaxMsgBits       uint32
		MaxMsgCells      uint32
		MaxLibraryCells  uint32
		MaxVmDataDepth   uint16
		MaxExtMsgSize    uint32
		MaxExtMsgDepth   uint16
		MaxAccStateCells uint32
		MaxAccStateBits  uint32
	} `tlbSumType:"#02"`
}

type ConfigParam43 struct {
	SizeLimitsConfig SizeLimitsConfig
}

type SuspendedAddressList struct {
	Magic          Magic `tlb:"#00"`
	Addresses      HashmapE[AddressWithWorkchain, struct{}]
	SuspendedUntil uint32
}

type ConfigParam44 struct {
	SuspendedAddressList SuspendedAddressList
}

type OracleBridgeParams struct {
	BridgeAddress         Bits256
	OracleMutlisigAddress Bits256
	Oracles               HashmapE[Bits256, Bits256]
	ExternalChainAddress  Bits256
}

type ConfigParam71 struct {
	OracleBridgeParams OracleBridgeParams
}

type ConfigParam72 struct {
	OracleBridgeParams OracleBridgeParams
}

type ConfigParam73 struct {
	OracleBridgeParams OracleBridgeParams
}

type JettonBridgePrices struct {
	BridgeBurnFee           Grams
	BridgeMintFee           Grams
	WalletMinTonsForStorage Grams
	WalletGasConsumption    Grams
	MinterMinTonsForStorage Grams
	DiscoverGasConsumption  Grams
}

type JettonBridgeParams struct {
	SumType
	JettonBridgeParamsV0 struct {
		BridgeAddress  Bits256
		OraclesAddress Bits256
		Oracles        HashmapE[Bits256, Bits256]
		StateFlags     uint8
		BurnBridgeFee  Grams
	} `tlbSumType:"#00"`
	JettonBridgeParamsV1 struct {
		BridgeAddress        Bits256
		OraclesAddress       Bits256
		Oracles              HashmapE[Bits256, Bits256]
		StateFlags           uint8
		Prices               JettonBridgePrices `tlb:"^"`
		ExternalChainAddress Bits256
	} `tlbSumType:"#01"`
}

type ConfigParam79 struct {
	JettonBridgeParams JettonBridgeParams
}

type ConfigParam81 struct {
	JettonBridgeParams JettonBridgeParams
}

type ConfigParam82 struct {
	JettonBridgeParams JettonBridgeParams
}
