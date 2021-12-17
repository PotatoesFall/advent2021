package main

const (
	LengthTypeID15BitSubpacketBitLength = 0
	LengthTypeID11BitSubpacketCount     = 1
)

func readPacket(input []bool) ([]bool, Packet) {
	var packet Packet

	input, packet.Version = readInt(input, 3)
	var typeID int
	input, typeID = readInt(input, 3)
	packet.Type = TypeID(typeID)

	if packet.Type == TypeIDLiteral {
		input, packet.Literal = readLiteral(input)
		return input, packet
	}

	input, packet.Children = readSubPackets(input)

	return input, packet
}

func readSubPackets(input []bool) ([]bool, []Packet) {
	var lengthTypeID int
	input, lengthTypeID = readInt(input, 1)

	if lengthTypeID == LengthTypeID11BitSubpacketCount {
		var subPacketCount int
		input, subPacketCount = readInt(input, 11)

		var subPackets []Packet
		for i := 0; i < subPacketCount; i++ {
			var subPacket Packet
			input, subPacket = readPacket(input)
			subPackets = append(subPackets, subPacket)
		}

		return input, subPackets
	}

	if lengthTypeID == LengthTypeID15BitSubpacketBitLength {
		var subPacketLength int
		input, subPacketLength = readInt(input, 15)

		subPackets := parsePackets(input[:subPacketLength])
		input = input[subPacketLength:]

		return input, subPackets
	}

	panic(lengthTypeID)
}

func parsePackets(input []bool) []Packet {
	var packets []Packet

	for len(input) != 0 {
		var packet Packet
		input, packet = readPacket(input)
		packets = append(packets, packet)
	}

	return packets
}

func readInt(input []bool, n int) ([]bool, int) {
	bools := input[:n]

	return input[n:], parseInt(bools)
}

func parseInt(bools []bool) int {
	var num int
	for _, b := range bools {
		num *= 2
		if b {
			num++
		}
	}
	return num
}

func readLiteral(input []bool) ([]bool, int) {
	keepReading := 1

	var literal int
	for keepReading == 1 {
		input, keepReading = readInt(input, 1)
		var n int
		input, n = readInt(input, 4)

		literal <<= 4
		literal += n
	}

	return input, literal
}
