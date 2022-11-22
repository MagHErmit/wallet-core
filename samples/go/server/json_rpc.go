package server

import (
	"encoding/hex"
	jsonparse "encoding/json"
	"errors"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"tw/core"
	"tw/protos/bitcoin"
	"tw/protos/common"
)

type Args struct {
	Gate string                 `json:"gate"`
	Tx   map[string]interface{} `json:"tx"`
}

type JSONServer struct {
	MN string
}

func ParseBtc(args *Args, bw *core.Wallet) (string, error) {

	// extract necessary info
	lockScript := core.BitcoinScriptLockScriptForAddress(bw.Address, bw.CoinType)
	priKeyByte, err := hex.DecodeString(bw.PriKey)

	// create main data
	input := bitcoin.SigningInput{
		HashType:   uint32(core.BitcoinSigHashTypeAll),
		PrivateKey: [][]byte{priKeyByte},
		CoinType:   uint32(core.CoinTypeBitcoin),
	}

	// create marshaled object for unmarshalling in SigningInput
	bt, err := jsonparse.Marshal(args.Tx)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// unmarshalling in SigningInput
	err = jsonparse.Unmarshal(bt, &input)
	if err != nil {
		log.Println(err)
		return "", err
	}

	for _, utxo := range input.Utxo {

		// hash from json parsing is incorrect, make re-unmarshaling
		bHash, err := jsonparse.Marshal(utxo.OutPoint.Hash)
		if err != nil {
			log.Println(err)
			return "", err
		}
		var strHash string
		err = jsonparse.Unmarshal(bHash, &strHash)
		if err != nil {
			log.Println(err)
			return "", err
		}
		hash, err := hex.DecodeString(strHash)
		if err != nil {
			log.Println(err)
			return "", err
		}
		utxo.OutPoint.Hash = hash

		// add lockScript
		utxo.Script = lockScript
	}

	var output bitcoin.SigningOutput
	err = signingTx(&input, bw.CoinType, &output)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if output.GetError() != common.SigningError_OK {
		return "", errors.New(output.GetError().String())
	}
	return hex.EncodeToString(output.GetEncoded()), nil
}

func signingTx(inputData proto.Message, ct core.CoinType, outputData proto.Message) error {
	// signing tx
	err := core.CreateSignedTx(inputData, ct, outputData)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (t *JSONServer) SignTx(r *http.Request, args *Args, reply *string) error {
	coinType, err := core.GetCoinByName(args.Gate)
	if err != nil {
		return err
	}

	wallet, err := core.CreateWalletWithMnemonic(t.MN, coinType)
	if err != nil {
		return err
	}

	// main switch
	switch coinType {
	case core.CoinTypeBitcoin:
		res, err := ParseBtc(args, wallet)
		if err != nil {
			log.Println(err)
			return err
		}
		*reply = res
		return nil
	}

	bz, err := jsonparse.Marshal(args.Tx)
	if err != nil {
		return err
	}

	*reply, err = core.CreateSignedTxFromJson(bz, wallet.PriKey, wallet.CoinType)
	if err != nil {
		return err
	}
	return nil
}
