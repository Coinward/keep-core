package cmd

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/keep-network/keep-core/config"
	"github.com/keep-network/keep-core/pkg/beacon"
	"github.com/keep-network/keep-core/pkg/chain/ethereum"
	"github.com/keep-network/keep-core/pkg/net/key"
	"github.com/keep-network/keep-core/pkg/net/libp2p"
	libp2pcrypto "github.com/libp2p/go-libp2p-crypto"
	"github.com/urfave/cli"
)

// StartCommand contains the definition of the start command-line subcommand.
var StartCommand cli.Command

const (
	bootstrapFlag = "bootstrap"
	portFlag      = "port"
	portShort     = "p"
)

const startDescription = `Starts the Keep client in the foreground. Currently this only consists of the
   threshold relay client for the Keep random beacon.`

func init() {
	StartCommand =
		cli.Command{
			Name:        "start",
			Usage:       `Starts the Keep client in the foreground`,
			Description: startDescription,
			Action:      Start,
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name: portFlag + "," + portShort,
				},
			},
		}
}

// Start starts a node; if it's not a bootstrap node it will get the Node.URLs
// from the config file
func Start(c *cli.Context) error {
	config, err := config.ReadConfig(c.GlobalString("config"))
	if err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	var port int
	if c.Int(portFlag) > 0 {
		config.LibP2P.Port = c.Int(portFlag)
	}

	isBootstrapNode := config.LibP2P.Seed != 0

	var (
		staticKey *key.StaticNetworkKey
	)
	if isBootstrapNode {
		staticKey, err = loadStaticKey(config.Ethereum.Account)
		if err != nil {
			return fmt.Errorf("error loading static peer's key [%v]", err)
		}
	} else {
		r := rand.Reader

		privKey, _, err := libp2pcrypto.GenerateSecp256k1Key(r)
		if err != nil {
			return err

		}
		staticKey = privKey.(*key.StaticNetworkKey)
	}

	ctx := context.Background()
	netProvider, err := libp2p.Connect(ctx, config.LibP2P, staticKey)
	if err != nil {
		return err
	}

	nodeHeader(isBootstrapNode, netProvider.AddrStrings(), port)

	chainProvider, err := ethereum.Connect(config.Ethereum)
	if err != nil {
		return fmt.Errorf("error connecting to Ethereum node: [%v]", err)
	}

	blockCounter, err := chainProvider.BlockCounter()
	if err != nil {
		return fmt.Errorf("error initializing blockcounter: [%v]", err)
	}

	err = beacon.Initialize(
		ctx,
		chainProvider.ThresholdRelay(),
		blockCounter,
		netProvider,
	)
	if err != nil {
		return fmt.Errorf("error initializing beacon: [%v]", err)
	}

	fmt.Println("Beacon initialized")
	select {
	case <-ctx.Done():
		if err != nil {
			return err
		}

		return fmt.Errorf("uh-oh, we went boom boom for no reason")
	}
}

func loadStaticKey(account ethereum.Account) (*key.StaticNetworkKey, error) {
	ethereumKey, err := ethereum.DecryptKeyFile(
		account.KeyFile,
		account.KeyFilePassword,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read KeyFile: %s [%v]", account.KeyFile, err,
		)
	}

	return key.EthereumKeyToNetworkKey(ethereumKey), nil
}
