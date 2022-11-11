package core

// #cgo CFLAGS: -I../../../include
// #cgo LDFLAGS: -L../../../build -L../../../build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lstdc++ -lm
// #include <TrustWalletCore/TWCoinType.h>
// #include <TrustWalletCore/TWCoinTypeConfiguration.h>
import "C"

import "tw/types"

type CoinType uint32

const (
	CoinTypeBitcoin              CoinType = C.TWCoinTypeBitcoin
	CoinTypeBinance              CoinType = C.TWCoinTypeBinance
	CoinTypeEthereum             CoinType = C.TWCoinTypeEthereum
	CoinTypeTron                 CoinType = C.TWCoinTypeTron
	CoinTypeAeternity            CoinType = C.TWCoinTypeAeternity
	CoinTypeAion                 CoinType = C.TWCoinTypeAion
	CoinTypeBitcoinCash          CoinType = C.TWCoinTypeBitcoinCash
	CoinTypeBitcoinGold          CoinType = C.TWCoinTypeBitcoinGold
	CoinTypeCallisto             CoinType = C.TWCoinTypeCallisto
	CoinTypeCardano              CoinType = C.TWCoinTypeCardano
	CoinTypeCosmos               CoinType = C.TWCoinTypeCosmos
	CoinTypeDash                 CoinType = C.TWCoinTypeDash
	CoinTypeDecred               CoinType = C.TWCoinTypeDecred
	CoinTypeDigiByte             CoinType = C.TWCoinTypeDigiByte
	CoinTypeDogecoin             CoinType = C.TWCoinTypeDogecoin
	CoinTypeEOS                  CoinType = C.TWCoinTypeEOS
	CoinTypeEthereumClassic      CoinType = C.TWCoinTypeEthereumClassic
	CoinTypeFIO                  CoinType = C.TWCoinTypeFIO
	CoinTypeGoChain              CoinType = C.TWCoinTypeGoChain
	CoinTypeGroestlcoin          CoinType = C.TWCoinTypeGroestlcoin
	CoinTypeICON                 CoinType = C.TWCoinTypeICON
	CoinTypeIoTeX                CoinType = C.TWCoinTypeIoTeX
	CoinTypeKava                 CoinType = C.TWCoinTypeKava
	CoinTypeKin                  CoinType = C.TWCoinTypeKin
	CoinTypeLitecoin             CoinType = C.TWCoinTypeLitecoin
	CoinTypeMonacoin             CoinType = C.TWCoinTypeMonacoin
	CoinTypeNebulas              CoinType = C.TWCoinTypeNebulas
	CoinTypeNULS                 CoinType = C.TWCoinTypeNULS
	CoinTypeNano                 CoinType = C.TWCoinTypeNano
	CoinTypeNEAR                 CoinType = C.TWCoinTypeNEAR
	CoinTypeNimiq                CoinType = C.TWCoinTypeNimiq
	CoinTypeOntology             CoinType = C.TWCoinTypeOntology
	CoinTypePOANetwork           CoinType = C.TWCoinTypePOANetwork
	CoinTypeQtum                 CoinType = C.TWCoinTypeQtum
	CoinTypeXRP                  CoinType = C.TWCoinTypeXRP
	CoinTypeSolana               CoinType = C.TWCoinTypeSolana
	CoinTypeStellar              CoinType = C.TWCoinTypeStellar
	CoinTypeTezos                CoinType = C.TWCoinTypeTezos
	CoinTypeTheta                CoinType = C.TWCoinTypeTheta
	CoinTypeThunderToken         CoinType = C.TWCoinTypeThunderToken
	CoinTypeNEO                  CoinType = C.TWCoinTypeNEO
	CoinTypeTomoChain            CoinType = C.TWCoinTypeTomoChain
	CoinTypeVeChain              CoinType = C.TWCoinTypeVeChain
	CoinTypeViacoin              CoinType = C.TWCoinTypeViacoin
	CoinTypeWanchain             CoinType = C.TWCoinTypeWanchain
	CoinTypeZcash                CoinType = C.TWCoinTypeZcash
	CoinTypeFiro                 CoinType = C.TWCoinTypeFiro
	CoinTypeZilliqa              CoinType = C.TWCoinTypeZilliqa
	CoinTypeZelcash              CoinType = C.TWCoinTypeZelcash
	CoinTypeRavencoin            CoinType = C.TWCoinTypeRavencoin
	CoinTypeWaves                CoinType = C.TWCoinTypeWaves
	CoinTypeTerra                CoinType = C.TWCoinTypeTerra
	CoinTypeTerraV2              CoinType = C.TWCoinTypeTerraV2
	CoinTypeHarmony              CoinType = C.TWCoinTypeHarmony
	CoinTypeAlgorand             CoinType = C.TWCoinTypeAlgorand
	CoinTypeKusama               CoinType = C.TWCoinTypeKusama
	CoinTypePolkadot             CoinType = C.TWCoinTypePolkadot
	CoinTypeFilecoin             CoinType = C.TWCoinTypeFilecoin
	CoinTypeElrond               CoinType = C.TWCoinTypeElrond
	CoinTypeBandChain            CoinType = C.TWCoinTypeBandChain
	CoinTypeSmartChainLegacy     CoinType = C.TWCoinTypeSmartChainLegacy
	CoinTypeSmartChain           CoinType = C.TWCoinTypeSmartChain
	CoinTypeOasis                CoinType = C.TWCoinTypeOasis
	CoinTypePolygon              CoinType = C.TWCoinTypePolygon
	CoinTypeTHORChain            CoinType = C.TWCoinTypeTHORChain
	CoinTypeBluzelle             CoinType = C.TWCoinTypeBluzelle
	CoinTypeOptimism             CoinType = C.TWCoinTypeOptimism
	CoinTypeZksync               CoinType = C.TWCoinTypeZksync
	CoinTypeArbitrum             CoinType = C.TWCoinTypeArbitrum
	CoinTypeECOChain             CoinType = C.TWCoinTypeECOChain
	CoinTypeAvalancheCChain      CoinType = C.TWCoinTypeAvalancheCChain
	CoinTypeXDai                 CoinType = C.TWCoinTypeXDai
	CoinTypeFantom               CoinType = C.TWCoinTypeFantom
	CoinTypeCryptoOrg            CoinType = C.TWCoinTypeCryptoOrg
	CoinTypeCelo                 CoinType = C.TWCoinTypeCelo
	CoinTypeRonin                CoinType = C.TWCoinTypeRonin
	CoinTypeOsmosis              CoinType = C.TWCoinTypeOsmosis
	CoinTypeECash                CoinType = C.TWCoinTypeECash
	CoinTypeCronosChain          CoinType = C.TWCoinTypeCronosChain
	CoinTypeSmartBitcoinCash     CoinType = C.TWCoinTypeSmartBitcoinCash
	CoinTypeKuCoinCommunityChain CoinType = C.TWCoinTypeKuCoinCommunityChain
	CoinTypeBoba                 CoinType = C.TWCoinTypeBoba
	CoinTypeMetis                CoinType = C.TWCoinTypeMetis
	CoinTypeAurora               CoinType = C.TWCoinTypeAurora
	CoinTypeEvmos                CoinType = C.TWCoinTypeEvmos
	CoinTypeNativeEvmos          CoinType = C.TWCoinTypeNativeEvmos
	CoinTypeMoonriver            CoinType = C.TWCoinTypeMoonriver
	CoinTypeMoonbeam             CoinType = C.TWCoinTypeMoonbeam
	CoinTypeKavaEvm              CoinType = C.TWCoinTypeKavaEvm
	CoinTypeKlaytn               CoinType = C.TWCoinTypeKlaytn
	CoinTypeMeter                CoinType = C.TWCoinTypeMeter
	CoinTypeOKXChain             CoinType = C.TWCoinTypeOKXChain
	CoinTypeNervos               CoinType = C.TWCoinTypeNervos
	CoinTypeEverscale            CoinType = C.TWCoinTypeEverscale
	CoinTypeAptos                CoinType = C.TWCoinTypeAptos
	CoinTypeHedera               CoinType = C.TWCoinTypeHedera
)

func (c CoinType) GetName() string {
	name := C.TWCoinTypeConfigurationGetName(C.enum_TWCoinType(c))
	defer C.TWStringDelete(name)
	return types.TWStringGoString(name)
}

func (c CoinType) Decimals() int {
	return int(C.TWCoinTypeConfigurationGetDecimals(C.enum_TWCoinType(c)))
}
