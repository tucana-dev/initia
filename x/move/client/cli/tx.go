package cli

import (
	"fmt"
	"os"
	"strings"

	"cosmossdk.io/core/address"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/initia-labs/initia/x/move/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(ac address.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Move transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		PublishCmd(ac),
		ExecuteCmd(ac),
		ScriptCmd(ac),
	)
	return txCmd
}

// PublishCmd will publish move binary files
func PublishCmd(ac address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "publish [move file1] [move file2] [...]",
		Short: "Publish move binary files",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Publish move binary files. allowed to upload up to 100 files

Example:
$ %s tx move publish \
    module1.mv \
	module2.mv \
	./dir/module3.mv
`, version.AppName,
			),
		),
		Args:    cobra.RangeArgs(1, 100),
		Aliases: []string{"p"},
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			bundle := make([][]byte, len(args))
			for i, arg := range args {
				// check arg has .mv suffix
				if !strings.HasSuffix(arg, ".mv") {
					return fmt.Errorf("invalid move file: %s", arg)
				}
				bundle[i], err = os.ReadFile(arg)
				if err != nil {
					return err
				}
			}

			upgradePolicyStr, err := cmd.Flags().GetString(FlagUpgradePolicy)
			if err != nil {
				return err
			}

			upgradePolicy, found := types.UpgradePolicy_value[upgradePolicyStr]
			if !found {
				return fmt.Errorf("invalid upgrade-policy `%s`", upgradePolicyStr)
			}

			sender, err := ac.BytesToString(clientCtx.FromAddress)
			if err != nil {
				return err
			}

			msg := types.MsgPublish{
				Sender:        sender,
				CodeBytes:     bundle,
				UpgradePolicy: types.UpgradePolicy(upgradePolicy),
			}

			if err = msg.Validate(ac); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetUpgradePolicy())
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// ExecuteCmd will execute an entry function of a published module.
func ExecuteCmd(ac address.Codec) *cobra.Command {
	bech32PrefixAccAddr := sdk.GetConfig().GetBech32AccountAddrPrefix()
	cmd := &cobra.Command{
		Use:   "execute [module address] [module name] [function name]",
		Short: "Execute an entry function of a published module",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Execute an entry function of a published module

Supported types : u8, u16, u32, u64, u128, u256, bool, string, address, raw_hex, raw_base64,
	vector<inner_type>, option<inner_type>, decimal128, decimal256, fixed_point32, fixed_point64
Example of args: address:0x1 bool:true u8:0 string:hello vector<u32>:a,b,c,d

Example:
$ %s tx move execute \
    %s1lwjmdnks33xwnmfayc64ycprww49n33mtm92ne \
	ManagedCoin \
	mint_to \
	--type-args '0x1::native_uinit::Coin 0x1::native_uusdc::Coin' \
 	--args 'u8:0 address:0x1 string:"hello world"'
`, version.AppName, bech32PrefixAccAddr,
			),
		),
		Aliases: []string{"ex", "e"},
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if _, err := types.AccAddressFromString(ac, args[0]); err != nil {
				return err
			}

			var typeArgs []string
			flagTypeArgs, err := cmd.Flags().GetString(FlagTypeArgs)
			if err != nil {
				return err
			}
			if flagTypeArgs != "" {
				typeArgs = strings.Split(flagTypeArgs, " ")
			}

			flagArgs, err := cmd.Flags().GetString(FlagArgs)
			if err != nil {
				return err
			}

			moveArgTypes, moveArgs := parseArguments(flagArgs)
			if len(moveArgTypes) != len(moveArgs) {
				return fmt.Errorf("invalid argument format len(moveArgTypes) != len(moveArgs)")
			}

			bcsArgs := [][]byte{}
			for i := range moveArgTypes {
				serializer := NewSerializer()
				bcsArg, err := BcsSerializeArg(moveArgTypes[i], moveArgs[i], serializer, ac)
				if err != nil {
					return err
				}

				bcsArgs = append(bcsArgs, bcsArg)
			}

			sender, err := ac.BytesToString(clientCtx.FromAddress)
			if err != nil {
				return err
			}

			msg := types.MsgExecute{
				Sender:        sender,
				ModuleAddress: args[0],
				ModuleName:    args[1],
				FunctionName:  args[2],
				TypeArgs:      typeArgs,
				Args:          bcsArgs,
			}

			if err = msg.Validate(ac); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetTypeArgs())
	cmd.Flags().AddFlagSet(FlagSetArgs())
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// ScriptCmd will execute a given script.
func ScriptCmd(ac address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "script [script-file]",
		Short: "Execute a given script",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Execute a given script

Supported types : u8, u16, u32, u64, u128, u256, bool, string, address, raw_hex, raw_base64,
	vector<inner_type>, option<inner_type>, decimal128, decimal256, fixed_point32, fixed_point64
Example of args: address:0x1 bool:true u8:0 string:hello vector<u32>:a,b,c,d

Example:
$ %s tx move script \
    ./script.mv \
	--type-args '0x1::native_uinit::Coin 0x1::native_uusdc::Coin' \
	--args 'u8:0 address:0x1'
`, version.AppName,
			),
		),
		Aliases: []string{"s"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			codeBytes, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			var typeArgs []string
			flagTypeArgs, err := cmd.Flags().GetString(FlagTypeArgs)
			if err != nil {
				return err
			}
			if flagTypeArgs != "" {
				typeArgs = strings.Split(flagTypeArgs, " ")
			}

			var flagArgsList []string
			flagArgs, err := cmd.Flags().GetString(FlagArgs)
			if err != nil {
				return err
			}
			if flagArgs != "" {
				flagArgsList = strings.Split(flagArgs, " ")
			}

			bcsArgs := make([][]byte, len(flagArgsList))
			for i, arg := range flagArgsList {
				argSplit := strings.Split(arg, ":")
				if len(argSplit) != 2 {
					return fmt.Errorf("invalid argument format: %s", arg)
				}

				serializer := NewSerializer()
				bcsArg, err := BcsSerializeArg(argSplit[0], argSplit[1], serializer, ac)
				if err != nil {
					return err
				}

				bcsArgs[i] = bcsArg
			}

			sender, err := ac.BytesToString(clientCtx.FromAddress)
			if err != nil {
				return err
			}

			msg := types.MsgScript{
				Sender:    sender,
				CodeBytes: codeBytes,
				TypeArgs:  typeArgs,
				Args:      bcsArgs,
			}

			if err = msg.Validate(ac); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	cmd.Flags().AddFlagSet(FlagSetTypeArgs())
	cmd.Flags().AddFlagSet(FlagSetArgs())
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
