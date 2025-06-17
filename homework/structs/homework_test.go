package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		name = name[:42]
		dataNamePtr := (*[42]byte)(unsafe.Pointer(&person.name))
		for i := 0; i < len(name); i++ {
			dataNamePtr[i] = name[i]
		}
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		// option 1
		person.x = *(*[4]byte)(unsafe.Pointer(&x))
		person.y = *(*[4]byte)(unsafe.Pointer(&y))
		person.z = *(*[4]byte)(unsafe.Pointer(&z))

		// option 2
		person.x = [4]byte{byte(x), byte(x >> 8), byte(x >> 16), byte(x >> 24)}
		person.y = [4]byte{byte(y), byte(y >> 8), byte(y >> 16), byte(y >> 24)}
		person.z = [4]byte{byte(z), byte(z >> 8), byte(z >> 16), byte(z >> 24)}
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		// option 1
		person.gold = *(*[4]byte)(unsafe.Pointer(&gold))

		// option 2
		person.gold = [4]byte{byte(gold), byte(gold >> 8), byte(gold >> 16), byte(gold >> 24)}
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.manaNHealth[0] = byte(mana)
		person.manaNHealth[1] &= 0b11110000
		person.manaNHealth[1] |= byte(mana>>8) & 0b1111
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.manaNHealth[1] &= 0b00001111
		person.manaNHealth[1] |= byte(health<<4) & 0b11110000
		person.manaNHealth[2] = byte(health >> 4)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.respectNStrength &= 0b0000
		person.respectNStrength |= byte(respect) & 0b1111
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.respectNStrength &= 0b00001111
		person.respectNStrength |= byte(strength) & 0b1111 << 4
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.experienceNLevel &= 0b0000
		person.experienceNLevel |= byte(experience) & 0b1111
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.experienceNLevel &= 0b00001111
		person.experienceNLevel |= byte(level) & 0b1111 << 4
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags |= 0b1
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags |= 0b10
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags |= 0b100
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		switch personType {
		case BuilderGamePersonType:
			person.flags |= 0b1000
		case BlacksmithGamePersonType:
			person.flags |= 0b10000
		case WarriorGamePersonType:
			person.flags |= 0b100000
		}
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	name             [42]byte // 42 byte     mana                  free                      free
	x, y, z          [4]byte  // 32 bit * 3   |                     |                         |
	gold             [4]byte  // 31 bit       v                     v                         V
	manaNHealth      [3]byte  // 10 bit * 2: |*,*,*,*,*,*,*,*|,|*,*,_,_,x,x,x,x|,|x,x,x,x,x,x,_,_|
	respectNStrength byte     // 10 bit * 2                             ^
	experienceNLevel byte     // 10 bit * 2                             |
	flags            byte     // 6 bit:      |*,*,*,*,*,*,_,_|        health
} //                                          ^ ^ ^ ^ ^ ^
//                                            | | | | | |
//                                            | | | | | warrior
//                                            | | | | blacksmith
//                                            | | | builder
//                                            | | family
//                                            | weapon
//                                            home
//

func NewGamePerson(options ...Option) GamePerson {
	person := GamePerson{}
	for _, option := range options {
		option(&person)
	}
	return person
}

func (p *GamePerson) Name() string {
	if p == nil {
		return ""
	}
	return unsafe.String(&p.name[0], 42)
}

func (p *GamePerson) X() int {
	if p == nil {
		return 0
	}

	// option 1
	return int(*(*int32)(unsafe.Pointer(&p.x)))

	// option 2
	return int(int32(p.x[0]) | int32(p.x[1])<<8 | int32(p.x[2])<<16 | int32(p.x[3])<<24)
}

func (p *GamePerson) Y() int {
	if p == nil {
		return 0
	}

	// option 1
	return int(*(*int32)(unsafe.Pointer(&p.y)))

	// option 2
	return int(int32(p.y[0]) | int32(p.y[1])<<8 | int32(p.y[2])<<16 | int32(p.y[3])<<24)
}

func (p *GamePerson) Z() int {
	if p == nil {
		return 0
	}

	// option 1
	return int(*(*int32)(unsafe.Pointer(&p.z)))

	// option 2
	return int(int32(p.z[0]) | int32(p.z[1])<<8 | int32(p.z[2])<<16 | int32(p.z[3])<<24)
}

func (p *GamePerson) Gold() int {
	if p == nil {
		return 0
	}

	// option 1
	return int(*(*uint32)(unsafe.Pointer(&p.gold)))

	// option 2
	return int(uint32(p.gold[0]) | uint32(p.gold[1])<<8 | uint32(p.gold[2])<<16 | uint32(p.gold[3])<<24)
}

func (p *GamePerson) Mana() int {
	if p == nil {
		return 0
	}

	return int(p.manaNHealth[0]) | int(p.manaNHealth[1]&0b1111)<<8
}

func (p *GamePerson) Health() int {
	if p == nil {
		return 0
	}

	return int(p.manaNHealth[1]&0b11110000)>>4 | int(p.manaNHealth[2])<<4
}

func (p *GamePerson) Respect() int {
	if p == nil {
		return 0
	}

	return int(p.respectNStrength & 0b1111)
}

func (p *GamePerson) Strength() int {
	if p == nil {
		return 0
	}

	return int(p.respectNStrength & 0b11110000 >> 4)
}

func (p *GamePerson) Experience() int {
	if p == nil {
		return 0
	}

	return int(p.experienceNLevel & 0b1111)
}

func (p *GamePerson) Level() int {
	if p == nil {
		return 0
	}

	return int(p.experienceNLevel & 0b11110000 >> 4)
}

func (p *GamePerson) HasHouse() bool {
	return p.flags&0b1 != 0
}

func (p *GamePerson) HasGun() bool {
	return p.flags&0b10 != 0
}

func (p *GamePerson) HasFamilty() bool {
	return p.flags&0b100 != 0
}

func (p *GamePerson) Type() int {
	if p.flags&0b1000 != 0 {
		return BuilderGamePersonType
	} else if p.flags&0b10000 != 0 {
		return BlacksmithGamePersonType
	} else if p.flags&0b100000 != 0 {
		return WarriorGamePersonType
	}
	return BuilderGamePersonType
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
	assert.True(t, unsafe.Sizeof(person) <= 64)
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
