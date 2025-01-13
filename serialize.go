package ion

type Serialize interface {
	Serialize(b *Buffer) (err error)
	Deserialize(b *Buffer) (err error)
}
