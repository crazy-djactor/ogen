package cli

import (
	"encoding/base64"
	"fmt"
	"github.com/fatih/color"
	"github.com/olympus-protocol/ogen/bls"
	"github.com/olympus-protocol/ogen/miner/keystore"
	"github.com/spf13/cobra"
	"path"
	"strconv"
)

func init() {
	rootCmd.AddCommand(generateKeysCmd)
}

var generateKeysCmd = &cobra.Command{
	Use:   "generate <numkeys>",
	Short: "Generates validator keys and saves them to your key store",
	Long:  `Generates validator keys and saves them to your key store`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		numKeys := 1
		if len(args) > 0 {
			numKeys, err = strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("invalid argument: %s\n", args[0])
				return
			}
		}

		keystorePath := path.Join(DataFolder, "wallet")
		k, err := keystore.NewBadgerKeystore(keystorePath)
		if err != nil {
			panic(err)
		}

		keys := make([]*bls.SecretKey, numKeys)
		for i := 0; i < numKeys; i++ {
			key, err := k.GenerateNewKey()
			if err != nil {
				fmt.Printf("error generating key: %s\n", err)
				return
			}
			keys[i] = key
		}

		err = k.Close()
		if err != nil {
			fmt.Printf("error closing wallet: %s\n", err)
		}

		colorHeader := color.New(color.FgCyan, color.Bold)
		colorSecret := color.New(color.FgRed)
		colorPubkey := color.New(color.FgGreen)
		//colorNormal := color.New(color.Fg)
		for i, k := range keys {
			colorHeader.Printf("Validator #%d\n", i)
			kBytes := k.Serialize()
			pkBytes := k.DerivePublicKey().Serialize()
			keyb64 := base64.StdEncoding.EncodeToString(kBytes[:])
			pkb64 := base64.StdEncoding.EncodeToString(pkBytes[:])
			colorSecret.Printf("Secret Key: ")
			fmt.Printf("%s\n", keyb64)
			colorPubkey.Printf("Public Key: ")
			fmt.Printf("%s\n", pkb64)
		}
	},
}