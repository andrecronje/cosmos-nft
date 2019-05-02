package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/spf13/cobra"
)

// GetCmdEditNFTMetadata is the CLI command for a EditNFTMetadata transaction
func GetCmdEditNFTMetadata(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-metadata [denom] [tokenID]",
		Short: "transfer a token of some denom with some tokenID to some recipient",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			msg, err := parseEditMetadataFlags()
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			tokenID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg = NewMsgEditNFTMetadata(cliCtx.GetFromAddress(),
				nft.Denom(args[0]),
				nft.TokenID(tokenID),
				msg.EditName,
				msg.EditDescription,
				msg.EditImage,
				msg.EditTokenURI,
				msg.Name,
				msg.Description,
				msg.Image,
				msg.TokenURI,
			)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagName, "", "Name/title of nft")
	cmd.Flags().String(flagDescription, "", "Description of nft")
	cmd.Flags().String(flagImage, "", "Image uri of nft")
	cmd.Flags().String(flagTokenURI, "", "URI for supplemental off-chain metadata (should return a JSON object)")

	return cmd
}

// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdMintNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [recipient] [denom] [tokenID]",
		Short: "mints a token of some denom with some tokenID to some recipient",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			metadata, err := parseEditMetadataFlags()
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			tokenID, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := nfts.NewMsgMintNFT(cliCtx.GetFromAddress(),
				sdk.AccAddress(args[0]),
				nfts.TokenID(tokenID),
				nfts.Denom(args[1]),
				metadata.Name,
				metadata.Description,
				metadata.Image,
				metadata.TokenURI,
			)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagName, "", "Name/title of nft")
	cmd.Flags().String(flagDescription, "", "Description of nft")
	cmd.Flags().String(flagImage, "", "Image uri of nft")
	cmd.Flags().String(flagTokenURI, "", "URI for supplemental off-chain metadata (should return a JSON object)")

	return cmd
}

// GetCmdBurnNFT is the CLI command for sending a BurnNFT transaction
func GetCmdBurnNFT(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "burn [denom] [tokenID]",
		Short: "burn a token of some denom with some tokenID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			// TODO: Does this need to be true? What does it mean to have an account that doesn't exist?
			// If it just means having a balance in some token then no, an account doens't need to "exist".
			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			tokenID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := nfts.NewMsgBurnNFT(cliCtx.GetFromAddress(), nfts.TokenID(tokenID), nfts.Denom(args[1]))
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
}
