package wallet_test

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/olympus-protocol/ogen/internal/chain"
	"github.com/olympus-protocol/ogen/internal/hostnode"
	"github.com/olympus-protocol/ogen/internal/mempool"
	"github.com/olympus-protocol/ogen/pkg/p2p"
	"os"
	"reflect"
	"testing"

	"github.com/olympus-protocol/ogen/internal/wallet"
	"github.com/olympus-protocol/ogen/pkg/bls"
	"github.com/olympus-protocol/ogen/pkg/params"
)

var testPass = "test_password"

var serTestPriv = []byte{111, 88, 18, 177, 76, 20, 221, 218, 2, 86, 121, 165, 156, 129, 88, 88, 57, 132, 159, 180, 206, 45, 20, 200, 46, 20, 186, 56, 22, 234, 137, 24}

func Test_NewWallet(t *testing.T) {
	wm, err := createWallet(t, false)
	if err != nil {
		t.Fatal(err)
	}
	secret, err := wm.GetSecret()
	if err != nil {
		t.Fatal(err)
	}
	if equal := reflect.DeepEqual(secret.Marshal(), serTestPriv); !equal {
		t.Fatal("secret keys don't match")
	}
	account, err := wm.GetAccount()
	if err != nil {
		t.Fatal(err)
	}
	testAccount := secret.PublicKey().ToAccount()

	if equal := reflect.DeepEqual(account, testAccount); !equal {
		t.Fatal("account keys don't match")
	}
	public, err := wm.GetPublic()
	if err != nil {
		t.Fatal(err)
	}
	if equal := reflect.DeepEqual(secret.PublicKey().Marshal(), public.Marshal()); !equal {
		t.Fatal("account keys don't match")
	}
	accountRaw, err := wm.GetAccountRaw()
	if err != nil {
		t.Fatal(err)
	}
	testAccRaw, err := secret.PublicKey().Hash()
	if err != nil {
		t.Fatal(err)
	}
	if equal := reflect.DeepEqual(accountRaw, testAccRaw); !equal {
		t.Fatal("account keys don't match")
	}
	clean()
}

func Test_OpenWallet(t *testing.T) {
	wm, err := createWallet(t, true)
	if err != nil {
		t.Fatal(err)
	}
	err = wm.OpenWallet("test_wallet", testPass)
	if err != nil {
		t.Fatal(err)
	}
	secret, err := wm.GetSecret()
	if err != nil {
		t.Fatal(err)
	}

	if equal := bytes.Compare(secret.Marshal(), serTestPriv); equal != 0 {
		t.Fatal("secret keys don't match")
	}
	account, err := wm.GetAccount()
	if err != nil {
		t.Fatal(err)
	}
	testAccount := secret.PublicKey().ToAccount()

	if equal := reflect.DeepEqual(account, testAccount); !equal {
		t.Fatal("account keys don't match")
	}
	public, err := wm.GetPublic()
	if err != nil {
		t.Fatal(err)
	}
	if equal := reflect.DeepEqual(secret.PublicKey().Marshal(), public.Marshal()); !equal {
		t.Fatal("account keys don't match")
	}
	accountRaw, err := wm.GetAccountRaw()
	if err != nil {
		t.Fatal(err)
	}
	testAccRaw, err := secret.PublicKey().Hash()
	if err != nil {
		t.Fatal(err)
	}
	if equal := reflect.DeepEqual(accountRaw, testAccRaw); !equal {
		t.Fatal("account keys don't match")
	}
	err = wm.CloseWallet()
	if err != nil {
		t.Fatal(err)
	}
	clean()
}

func Test_OpenWalletWithWrongPassword(t *testing.T) {
	wm, err := createWallet(t, true)
	if err != nil {
		t.Fatal(err)
	}
	err = wm.OpenWallet("test_wallet", "wrong_pass")
	if err == nil {
		t.Fatal(err)
	}
	clean()
}

func createWallet(t *testing.T, close bool) (wallet.Wallet, error) {
	ctrl := gomock.NewController(t)

	ch := chain.NewMockBlockchain(ctrl)
	hn := hostnode.NewMockHostNode(ctrl)
	hn.EXPECT().Topic(p2p.MsgTxCmd)
	hn.EXPECT().Topic(p2p.MsgDepositCmd)
	hn.EXPECT().Topic(p2p.MsgDepositsCmd)
	hn.EXPECT().Topic(p2p.MsgExitCmd)
	hn.EXPECT().Topic(p2p.MsgExitsCmd)

	coinspool := mempool.NewMockCoinsMempool(ctrl)
	actionpool := mempool.NewMockActionMempool(ctrl)

	walletMan, err := wallet.NewWallet(context.Background(), nil, "./", &params.TestNet, ch, hn, coinspool, actionpool)
	if err != nil {
		return nil, err
	}
	priv, err := bls.SecretKeyFromBytes(serTestPriv)
	if err != nil {
		return nil, err
	}
	err = walletMan.NewWallet("test_wallet", priv, testPass)
	if err != nil {
		return nil, err
	}
	if close {
		err = walletMan.CloseWallet()
		if err != nil {
			return nil, err
		}
	}
	return walletMan, nil
}

func clean() {
	_ = os.RemoveAll("./wallets")
}
