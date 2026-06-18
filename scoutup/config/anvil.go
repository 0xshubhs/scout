package config

func PrepareDefaultAnvilConfig() *NetworkConfig {
	return &NetworkConfig{
		Chains: []*ChainConfig{
			{
				Name:       "Anvil",
				RPCUrl:     "http://localhost:8545",
				FirstBlock: 0,
				ChainID:    31337,
			},
		},
	}
}
