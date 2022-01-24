//
// Copyright 2021, Offchain Labs, Inc. All rights reserved.
//

package precompiles

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

// This precompile provides owners with tools for managing the rollup.
// All calls to this precompile are authorized by the OwnerPrecompile wrapper,
// which ensures only a chain owner can access these methods. For methods that
// are safe for non-owners to call, see ArbOwnerOld
type ArbOwner struct {
	Address addr // 0x70
}

// Promotes the user to chain owner
func (con ArbOwner) AddChainOwner(c ctx, evm mech, newOwner addr) error {
	return c.state.ChainOwners().Add(newOwner)
}

// Demotes the user from chain owner
func (con ArbOwner) RemoveChainOwner(c ctx, evm mech, addr addr) error {
	member, _ := con.IsChainOwner(c, evm, addr)
	if !member {
		return errors.New("tried to remove non-owner")
	}
	return c.state.ChainOwners().Remove(addr)
}

// See if the user is a chain owner
func (con ArbOwner) IsChainOwner(c ctx, evm mech, addr addr) (bool, error) {
	return c.state.ChainOwners().IsMember(addr)
}

// Retrieves the list of chain owners
func (con ArbOwner) GetAllChainOwners(c ctx, evm mech) ([]common.Address, error) {
	return c.state.ChainOwners().AllMembers()
}

// Sets the L1 gas price estimate directly, bypassing the autoregression
func (con ArbOwner) SetL1GasPriceEstimate(c ctx, evm mech, priceInWei huge) error {
	return c.state.L1PricingState().SetL1GasPriceEstimateWei(priceInWei)
}

// Sets the L2 gas price directly, bypassing the pool calculus
func (con ArbOwner) SetL2GasPrice(c ctx, evm mech, priceInWei huge) error {
	return c.state.L2PricingState().SetGasPriceWei(priceInWei)
}

func (con ArbOwner) SetMinimumGasPrice(c ctx, evm mech, priceInWei huge) error {
	return c.state.L2PricingState().SetGasPriceWei(priceInWei)
}

func (con ArbOwner) SetSpeedLimit(c ctx, evm mech, limit uint64) error {
	return c.state.L2PricingState().SetSpeedLimitPerSecond(limit)
}

func (con ArbOwner) SetGasPoolSeconds(c ctx, evm mech, seconds uint64) error {
	return c.state.L2PricingState().SetGasPoolSeconds(seconds)
}

func (con ArbOwner) SetSmallGasPoolSeconds(c ctx, evm mech, seconds uint64) error {
	return c.state.L2PricingState().SetSmallGasPoolSeconds(seconds)
}

func (con ArbOwner) SetMaxTxGasLimit(c ctx, evm mech, limit uint64) error {
	return c.state.L2PricingState().SetMaxPerBlockGasLimit(limit)
}

func (con ArbOwner) GetNetworkFeeAccount(c ctx, evm mech) (addr, error) {
	return c.state.NetworkFeeAccount()
}

func (con ArbOwner) SetNetworkFeeAccount(c ctx, evm mech, newNetworkFeeAccount addr) error {
	return c.state.SetNetworkFeeAccount(newNetworkFeeAccount)
}
