package main

// TODO:
//func addGenesisWasmMsgCmd(defaultNodeHome string) *cobra.Command {
//	txCmd := &cobra.Command{
//		Use:                        "add-wasm-message",
//		Short:                      "Wasm genesis subcommands",
//		DisableFlagParsing:         true,
//		SuggestionsMinimumDistance: 2,
//		RunE:                       client.ValidateCmd,
//	}
//
//	genesisIO := wasmcli.NewDefaultGenesisIO()
//
//	txCmd.AddCommand(
//		wasmcli.GenesisStoreCodeCmd(defaultNodeHome, genesisIO),
//		wasmcli.GenesisInstantiateContractCmd(defaultNodeHome, genesisIO),
//		wasmcli.GenesisExecuteContractCmd(defaultNodeHome, genesisIO),
//		wasmcli.GenesisListContractsCmd(defaultNodeHome, genesisIO),
//		wasmcli.GenesisListCodesCmd(defaultNodeHome, genesisIO),
//	)
//
//	return txCmd
//}
