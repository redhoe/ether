package erc20

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/redhoe/ether"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"testing"
)

var engine *ether.Engine

func init() {
	logger, _ := zap.NewDevelopment()
	engine = ether.NewEngine(logger, "https://mainnet.infura.io/v3/14e5c24b98634138a9127fc8db299970", "")
	engine.SetGasPrice(decimal.New(10, 9).BigInt())
}

func TestErc20Abi_Method(t *testing.T) {

	//b, _ := Erc20Abi.BalanceOf(httpCommon.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d"))
	//t.Log(hexutil.Encode(b))

	_, err := IErc20Abi.abi.EventByID(common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"))

	if err != nil {
		panic(err)
	}

	//fmt.Println(event.Inputs)
}

func TestErc20_Name(t *testing.T) {
	e := NewErc20(engine)
	t.Log(e.Name(common.HexToAddress("0x0000000000095413afC295d19EDeb1Ad7B71c952")))
	t.Log(e.Symbol(common.HexToAddress("0x0000000000095413afC295d19EDeb1Ad7B71c952")))
	t.Log(e.Decimals(common.HexToAddress("0x0000000000095413afC295d19EDeb1Ad7B71c952")))

	//t.Log(e.Allowance(
	//	httpCommon.HexToAddress("0x78867BbEeF44f2326bF8DDd1941a4439382EF2A7"),
	//	httpCommon.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d"),
	//	httpCommon.HexToAddress("0xa1be04c6F760D887Fd83570734c8B06F77B8826e"),
	//))
}

func TestErc20_BalanceOf(t *testing.T) {
	e := NewErc20(engine)
	t.Log(e.BalanceOf(common.HexToAddress("0x762f7957714CfCfd271a9078e8f446692627126f"), common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d")))
}

func TestErc20_Approve(t *testing.T) {
	e := NewErc20(engine)
	t.Log(e.BalanceOf(common.HexToAddress("0x762f7957714CfCfd271a9078e8f446692627126f"), common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d")))

	hash, buildTx, err := e.Approve(
		common.HexToAddress("0x78867BbEeF44f2326bF8DDd1941a4439382EF2A7"),
		common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d"),
		common.HexToAddress("0x762f7957714CfCfd271a9078e8f446692627126f"),
		decimal.New(2, 18).BigInt(),
		"",
	)
	t.Log(hash, buildTx, err)
}

func TestErc20_Transfer(t *testing.T) {
	e := NewErc20(engine)
	hash, buildTx, err := e.Transfer(
		common.HexToAddress("0x762f7957714CfCfd271a9078e8f446692627126f"),
		common.HexToAddress("0xa1be04c6F760D887Fd83570734c8B06F77B8826e"),
		decimal.New(100, 18).BigInt(),
		"",
	)

	if err != nil {
		fmt.Printf("%+v", err)
	}

	t.Log(hash, buildTx)
}

func TestErrors(t *testing.T) {

	err := errors.New("a err")
	new := errors.WithStack(err)

	panic(new)
}
