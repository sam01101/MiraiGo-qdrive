package tlv

import "github.com/sam01101/MiraiGo-qdrive/binary"

func T193(ticket string) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x193)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.Write([]byte(ticket))
		}))
	})
}
