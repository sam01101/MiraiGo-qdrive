package tlv

import "github.com/sam01101/MiraiGo-qdrive/binary"

func T108(imei string) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x108)
		w.WriteTlv([]byte(imei))
	})
}
