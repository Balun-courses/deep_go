// GamePerson binary layout (64 bytes):
// [0..41]    Name (42 bytes)
// [42..45]   X coordinate (int32, 4 bytes, big-endian)
// [46..49]   Y coordinate (int32, 4 bytes, big-endian)
// [50..53]   Z coordinate (int32, 4 bytes, big-endian)
// [54..57]   Gold (int32, 4 bytes, big-endian)
// [58..59]   Mana (10 bits: 8-bit low in byte 58, 2-bit high in bits 4-5 of byte 59)
// [60..61]   Health (10 bits: 8-bit low in byte 60, 2-bit high in bits 4-5 of byte 61)
// [62]       Respect (4 bits high nibble) + Strength (4 bits low nibble)
// [63]       Experience (4 bits high nibble) + Level (4 bits low nibble)
// [61]       also holds flags and type:
//              bit 0 = HasHouse, bit 1 = HasGun, bit 2 = HasFamily
//              bits 6-7 = Type (0=Builder,1=Blacksmith,2=Warrior)

package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// Player types (2-bit):
const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

// Layout sizes
const (
	NameMaxLength = 42 // bytes allocated for name
)

// Byte offsets within data buffer
const (
	dX          = NameMaxLength // 42: X coordinate (int32)
	dY          = dX + 4        // 46: Y coordinate (int32)
	dZ          = dY + 4        // 50: Z coordinate (int32)
	dGold       = dZ + 4        // 54: Gold (int32)
	dMana       = dGold + 4     // 58: Mana (10‑bit)
	dHealth     = dMana + 2     // 60: Health (10‑bit)
	dRespect    = dHealth + 2   // 62: Respect/Strength (4+4 bits)
	dExperience = dRespect + 1  // 63: Experience/Level (4+4 bits)
)

// Flags and type packed into bits of byte at flagsOffset
const (
	flagsOffset = dHealth + 1 // 61: shared with high 2 bits of Health

	flagHouseBit  = 0 // bit 0: HasHouse
	flagGunBit    = 1 // bit 1: HasGun
	flagFamilyBit = 2 // bit 2: HasFamily

	typeShift = 6 // bits 6-7: player type
	typeMask  = 0x03 << typeShift
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		if len(name) > NameMaxLength {
			panic("NameMaxLength exeeded")
		}
		copy(person.data[:], name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.writeInt32Raw(dX, x)
		person.writeInt32Raw(dY, y)
		person.writeInt32Raw(dZ, z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.writeInt32Raw(dGold, gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.write10bitValue(dMana, mana)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.write10bitValue(dHealth, health)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.write4bitValue(dRespect, respect, true)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.write4bitValue(dRespect, strength, false)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.write4bitValue(dExperience, experience, true)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.write4bitValue(dExperience, level, false)
	}
}

func WithHouse() func(*GamePerson) {
	return func(p *GamePerson) {
		p.data[flagsOffset] |= 1 << flagHouseBit
	}
}

func WithGun() func(*GamePerson) {
	return func(p *GamePerson) {
		p.data[flagsOffset] |= 1 << flagGunBit
	}
}

func WithFamily() func(*GamePerson) {
	return func(p *GamePerson) {
		p.data[flagsOffset] |= 1 << flagFamilyBit
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(p *GamePerson) {
		if personType < BuilderGamePersonType || personType > WarriorGamePersonType {
			panic("invalid person type")
		}

		p.data[flagsOffset] &^= typeMask
		p.data[flagsOffset] |= byte(personType) << typeShift
	}
}

type GamePerson struct {
	data [64]byte
}

func NewGamePerson(options ...Option) GamePerson {
	person := GamePerson{}

	for _, option := range options {
		option(&person)
	}

	return person
}

func (p *GamePerson) Name() string {
	var nameLength int
	for i := range NameMaxLength {
		if p.data[i] == 0 {
			break
		}
		nameLength++
	}
	return unsafe.String(&p.data[0], nameLength)
}

func (p *GamePerson) X() int {
	return int(p.readInt32Raw(dX))
}

func (p *GamePerson) Y() int {
	return int(p.readInt32Raw(dY))
}

func (p *GamePerson) Z() int {
	return int(p.readInt32Raw(dZ))
}

func (p *GamePerson) Gold() int {
	return int(p.readInt32Raw(dGold))
}

func (p *GamePerson) Mana() int {
	return p.read10bitValue(dMana)
}

func (p *GamePerson) Health() int {
	return p.read10bitValue(dHealth)
}

func (p *GamePerson) Respect() int {
	return p.read4bitValue(dRespect, true)
}

func (p *GamePerson) Strength() int {
	return p.read4bitValue(dRespect, false)
}

func (p *GamePerson) Experience() int {
	return p.read4bitValue(dExperience, true)
}

func (p *GamePerson) Level() int {
	return p.read4bitValue(dExperience, false)
}

func (p *GamePerson) HasHouse() bool {
	return (p.data[flagsOffset]>>flagHouseBit)&1 != 0
}

func (p *GamePerson) HasGun() bool {
	return (p.data[flagsOffset]>>flagGunBit)&1 != 0
}

func (p *GamePerson) HasFamilty() bool {
	return (p.data[flagsOffset]>>flagFamilyBit)&1 != 0
}

func (p *GamePerson) Type() int {
	return int((p.data[flagsOffset] & typeMask) >> typeShift)
}

func (p *GamePerson) writeInt32Raw(offset int, value int) {
	if offset+4 > len(p.data) {
		panic("offset out of range")
	}
	if value < math.MinInt32 || value > math.MaxInt32 {
		panic("value out of int32 range")
	}
	p.data[offset] = byte(value >> 24)
	p.data[offset+1] = byte(value >> 16)
	p.data[offset+2] = byte(value >> 8)
	p.data[offset+3] = byte(value)
}

func (p *GamePerson) readInt32Raw(offset int) int32 {
	if offset+4 > len(p.data) {
		panic("offset out of range")
	}
	value := int32(p.data[offset])<<24 |
		int32(p.data[offset+1])<<16 |
		int32(p.data[offset+2])<<8 |
		int32(p.data[offset+3])
	return value
}

func (p *GamePerson) write10bitValue(offset, value int) {
	if value < 0 || value > 0x3FF {
		panic("value out of 10-bit range")
	}
	p.data[offset] = byte(value & 0xFF)
	p.data[offset+1] = (p.data[offset+1] & 0x0F) | byte((value>>8)&0x03)<<4
}

func (p *GamePerson) read10bitValue(offset int) int {
	low := int(p.data[offset])
	high := int((p.data[offset+1] >> 4) & 0x03)
	return (high << 8) | low
}

func (p *GamePerson) write4bitValue(offset, value int, highNibble bool) {
	if value < 0 || value > 0xF {
		panic("value out of 4-bit range")
	}

	if highNibble {
		p.data[offset] = (p.data[offset] & 0x0F) | byte(value<<4)
	} else {
		p.data[offset] = (p.data[offset] & 0xF0) | byte(value&0x0F)
	}
}

func (p *GamePerson) read4bitValue(offset int, highNibble bool) int {
	if highNibble {
		return int(p.data[offset] >> 4)
	}
	return int(p.data[offset] & 0x0F)
}

func TestGamePerson(t *testing.T) {
	assert.LessOrEqual(t, unsafe.Sizeof(GamePerson{}), uintptr(64))

	const x, y, z = math.MinInt32, math.MaxInt32, 0
	const name = "aaaaaaaaaaaaa_bbbbbbbbbbbbb_cccccccccccccc"
	const personType = BuilderGamePersonType
	const gold = math.MaxInt32
	const mana = 1000
	const health = 1000
	const respect = 10
	const strength = 10
	const experience = 10
	const level = 10

	options := []Option{
		WithName(name),
		WithCoordinates(x, y, z),
		WithGold(gold),
		WithMana(mana),
		WithHealth(health),
		WithRespect(respect),
		WithStrength(strength),
		WithExperience(experience),
		WithLevel(level),
		WithHouse(),
		WithFamily(),
		WithType(personType),
	}

	person := NewGamePerson(options...)
	assert.Equal(t, name, person.Name())
	assert.Equal(t, x, person.X())
	assert.Equal(t, y, person.Y())
	assert.Equal(t, z, person.Z())
	assert.Equal(t, gold, person.Gold())
	assert.Equal(t, mana, person.Mana())
	assert.Equal(t, health, person.Health())
	assert.Equal(t, respect, person.Respect())
	assert.Equal(t, strength, person.Strength())
	assert.Equal(t, experience, person.Experience())
	assert.Equal(t, level, person.Level())
	assert.True(t, person.HasHouse())
	assert.True(t, person.HasFamilty())
	assert.False(t, person.HasGun())
	assert.Equal(t, personType, person.Type())
}
