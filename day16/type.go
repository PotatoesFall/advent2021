package main

type TypeID int

const (
	TypeIDLiteral TypeID = 4

	TypeIDSum         TypeID = 0
	TypeIDProduct     TypeID = 1
	TypeIDMinimum     TypeID = 2
	TypeIDMaximum     TypeID = 3
	TypeIDGreaterThan TypeID = 5
	TypeIDLessThan    TypeID = 6
	TypeIDEqualTo     TypeID = 7
)

type Packet struct {
	Version  int
	Type     TypeID
	Literal  int
	Children []Packet
}
