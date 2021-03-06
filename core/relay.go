package core

import (
"fmt"
	"github.com/sirupsen/logrus"

	"github.com/Polkadex-Substrate/eth-relayer/workers"
// 	"github.com/Polkadex-Substrate/eth-relayer/workers/beefyrelayer"
	"github.com/Polkadex-Substrate/eth-relayer/workers/ethrelayer"
// 	"github.com/Polkadex-Substrate/eth-relayer/workers/parachaincommitmentrelayer"
)

type Relay struct{}

func (re *Relay) Run() error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	var pool workers.WorkerPool
	fmt.Println("Running the worker")
	if config.Workers.EthRelayer.Enabled {
		ethrelayerFactory := func() (workers.Worker, *workers.WorkerConfig, error) {
			return ethrelayer.NewWorker(
				&config.Eth,
				&config.Parachain,
				logrus.WithField("worker", ethrelayer.Name),
			), &config.Workers.EthRelayer, nil
		}
		pool = append(pool, ethrelayerFactory)
		fmt.Println("Creating the Eth Relayer Worker")
	}

// 	if config.Workers.BeefyRelayer.Enabled {
// 		beefyrelayerFactory := func() (workers.Worker, *workers.WorkerConfig, error) {
// 			beefyRelayer, err := beefyrelayer.NewWorker(
// 				&config.Relaychain,
// 				&config.Eth,
// 				&config.BeefyRelayerDatabase,
// 				logrus.WithField("worker", beefyrelayer.Name),
// 			)
// 			if err != nil {
// 				return nil, nil, err
// 			}
// 			return beefyRelayer, &config.Workers.BeefyRelayer, nil
// 		}
// 		pool = append(pool, beefyrelayerFactory)
// 	}

// 	if config.Workers.ParachainCommitmentRelayer.Enabled {
// 		parachaincommitmentrelayerFactory := func() (workers.Worker, *workers.WorkerConfig, error) {
// 			parachainCommitmentRelayer, err := parachaincommitmentrelayer.NewWorker(
// 				&config.Parachain,
// 				&config.Relaychain,
// 				&config.Eth,
// 				logrus.WithField("worker", parachaincommitmentrelayer.Name),
// 			)
// 			if err != nil {
// 				return nil, nil, err
// 			}
// 			return parachainCommitmentRelayer, &config.Workers.ParachainCommitmentRelayer, nil
// 		}
// 		pool = append(pool, parachaincommitmentrelayerFactory)
// 	}

	return pool.Run()
}
