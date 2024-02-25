package cpu

import "encoding/json"

type Register struct {
	value uint16
	mask  uint16
}

func (r *Register) Set(v uint16) {
	r.value = v & r.mask
}

func (r *Register) Val() uint16 {
	return r.value
}

func (r *Register) SetLo(v byte) {
	r.Set((r.value & 0xFF00) | uint16(v))
}

func (r *Register) SetHi(v byte) {
	r.Set((uint16(v) << 8) | (r.value & 0x00FF))
}

func (r *Register) Lo() byte {
	return byte(r.value & 0xFF)
}

func (r *Register) Hi() byte {
	return byte((r.value & 0xFF00) >> 8)
}

func (r *Register) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value uint16
		Mask  uint16
	}{
		Value: r.value,
		Mask:  r.mask,
	})
}

func (r *Register) UnmarshalJSON(bs []byte) error {
	var reg struct {
		Value uint16
		Mask  uint16
	}

	err := json.Unmarshal(bs, &reg)

	if err != nil {
		return err
	}

	r.value = reg.Value
	r.mask = reg.Mask

	return nil
}
