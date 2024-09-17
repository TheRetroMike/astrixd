package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// AstrixMainnetPrivate is the version that is used for
// astrix mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var AstrixMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// AstrixMainnetPublic is the version that is used for
// astrix mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var AstrixMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// AstrixTestnetPrivate is the version that is used for
// astrix testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var AstrixTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// AstrixTestnetPublic is the version that is used for
// astrix testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var AstrixTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// AstrixDevnetPrivate is the version that is used for
// astrix devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var AstrixDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// AstrixDevnetPublic is the version that is used for
// astrix devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var AstrixDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// AstrixSimnetPrivate is the version that is used for
// astrix simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var AstrixSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// AstrixSimnetPublic is the version that is used for
// astrix simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var AstrixSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case AstrixMainnetPrivate:
		return AstrixMainnetPublic, nil
	case AstrixTestnetPrivate:
		return AstrixTestnetPublic, nil
	case AstrixDevnetPrivate:
		return AstrixDevnetPublic, nil
	case AstrixSimnetPrivate:
		return AstrixSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case AstrixMainnetPrivate:
		return true
	case AstrixTestnetPrivate:
		return true
	case AstrixDevnetPrivate:
		return true
	case AstrixSimnetPrivate:
		return true
	}

	return false
}
