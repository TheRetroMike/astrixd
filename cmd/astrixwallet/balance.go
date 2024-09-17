package main

import (
	"context"
	"fmt"

	"github.com/astrix-network/astrixd/cmd/astrixwallet/daemon/client"
	"github.com/astrix-network/astrixd/cmd/astrixwallet/daemon/pb"
	"github.com/astrix-network/astrixd/cmd/astrixwallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatAix(addressBalance.Available), utils.FormatAix(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, AIX %s %s%s\n", utils.FormatAix(response.Available), utils.FormatAix(response.Pending), pendingSuffix)

	return nil
}
