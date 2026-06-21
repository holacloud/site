package store

type Id struct {
	Id      string `json:"id" bson:"_id"`
	Version int64  `json:"version" bson:"version"`
}

func NewId(id string) *Id {
	return &Id{
		Id:      id,
		Version: 0,
	}
}

func (i *Id) GetId() string {
	return i.Id
}

func (i *Id) GetVersion() int64 {
	return i.Version
}

func (i *Id) SetVersion(version int64) {
	i.Version = version
}
