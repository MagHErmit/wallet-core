package core

// #cgo CFLAGS: -I../../../include
// #cgo LDFLAGS: -L../../../build -L../../../build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lstdc++ -lm
// #include <TrustWalletCore/TWCoinType.h>
// #include <TrustWalletCore/TWAnySigner.h>
import "C"

import (
	"errors"
	"google.golang.org/protobuf/proto"
	"tw/types"
)

func CreateSignedTx(inputData proto.Message, ct CoinType, outputData proto.Message) error {
	ibytes, _ := proto.Marshal(inputData)
	idata := types.TWDataCreateWithGoBytes(ibytes)
	defer C.TWDataDelete(idata)

	odata := C.TWAnySignerSign(idata, C.enum_TWCoinType(ct))
	defer C.TWDataDelete(odata)

	err := proto.Unmarshal(types.TWDataGoBytes(odata), outputData)
	if err != nil {
		return err
	}
	return nil
}

func CreateSignedTxFromJson(json []byte, key string, ct CoinType) (string, error) {
	supported := C.TWAnySignerSupportsJSON(C.enum_TWCoinType(ct))
	if !supported {
		return "", errors.New("json signing not supported for given coin")
	}

	ijson := types.TWStringCreateWithGoString(string(json))
	defer C.TWStringDelete(ijson)

	ikey := types.TWDataCreateWithGoHex(key)
	defer C.TWDataDelete(ikey)

	odata := C.TWAnySignerSignJSON(ijson, ikey, C.enum_TWCoinType(ct))
	defer C.TWStringDelete(odata)

	res := types.TWStringGoString(odata)

	return res, nil

}
