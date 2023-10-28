package ki

import (
	"github.com/google/uuid"
	"github.com/neonnetwork/ki/pkg/structure"
)

type Window interface {
	Id() uuid.UUID
	Position() structure.Vector2[int32]
	Size() structure.Vector2[int32]
	Render() error
	Split() structure.Pair[structure.Vector2[int32], structure.Vector2[int32]]
	SetPosition(vector structure.Vector2[int32])
	SetSize(vector structure.Vector2[int32])
}
