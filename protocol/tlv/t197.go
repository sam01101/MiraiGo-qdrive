package tlv

import "github.com/sam01101/MiraiGo-qdrive/binary"

func T197() []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x197)
		w.WriteTlv([]byte{0})
	})
}
