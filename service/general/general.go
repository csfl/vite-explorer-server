package general

import "math/big"

func GetFfmCap () (string, error) {
	var ffmCap = "3,071,003,208"
	return ffmCap, nil
}

func GetSysTps () (string, error) {
	var sysTps = "9.1"
	return sysTps, nil
}

func GetCirPrice () (int, error) {
	var cirPrice = 30
	return cirPrice, nil
}

func GetTxTotalTAmount () (*big.Int, error) {
	var txTotalTAmount = &big.Int{}
	return txTotalTAmount, nil
}

func GetTxMonAmount () (*big.Int, error) {
	var txMonAmount = &big.Int{}
	return txMonAmount, nil
}