package patterns

import "github.com/motivewc/wowpatch/internal/binary"

var (
	PortalPattern            = binary.StringToPattern(".actual.battle.net")
	ConnectToModulusPattern  = binary.Pattern{0x91, 0xD5, 0x9B, 0xB7, 0xD4, 0xE1, 0x83, 0xA5}
	CryptoEdPublicKeyPattern = binary.Pattern{0x15, 0xD6, 0x18, 0xBD, 0x7D, 0xB5, 0x77, 0xBD}
)
