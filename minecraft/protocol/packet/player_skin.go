package packet

import (
	"bytes"
	"encoding/binary"
	"github.com/google/uuid"
	"phoenixbuilder/minecraft/protocol"
)

// PlayerSkin is sent by the client to the server when it updates its own skin using the in-game skin picker.
// It is relayed by the server, or sent if the server changes the skin of a player on its own accord. Note
// that the packet can only be sent for players that are in the player list at the time of sending.
type PlayerSkin struct {
	// UUID is the UUID of the player as sent in the Login packet when the client joined the server. It must
	// match this UUID exactly for the skin to show up on the player.
	UUID uuid.UUID
	// Skin is the new skin to be applied on the player with the UUID in the field above. The skin, including
	// its animations, will be shown after sending it.
	Skin protocol.Skin
	// NewSkinName no longer has a function: The field can be left empty at all times.
	NewSkinName string
	// OldSkinName no longer has a function: The field can be left empty at all times.
	OldSkinName string
}

// ID ...
func (*PlayerSkin) ID() uint32 {
	return IDPlayerSkin
}

// Marshal ...
func (pk *PlayerSkin) Marshal(buf *bytes.Buffer) {
	_ = protocol.WriteUUID(buf, pk.UUID)
	_ = protocol.WriteSerialisedSkin(buf, pk.Skin)
	_ = protocol.WriteString(buf, pk.NewSkinName)
	_ = protocol.WriteString(buf, pk.OldSkinName)
	_ = binary.Write(buf, binary.LittleEndian, pk.Skin.Trusted)
}

// Unmarshal ...
func (pk *PlayerSkin) Unmarshal(buf *bytes.Buffer) error {
	return chainErr(
		protocol.UUID(buf, &pk.UUID),
		protocol.SerialisedSkin(buf, &pk.Skin),
		protocol.String(buf, &pk.NewSkinName),
		protocol.String(buf, &pk.OldSkinName),
		binary.Read(buf, binary.LittleEndian, &pk.Skin.Trusted),
	)
}
