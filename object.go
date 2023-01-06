package kDrive

type ObjectType string

func (ot ObjectType) String() string {
	return string(ot)
}

type Cursor string

func (c Cursor) String() string {
	return string(c)
}
