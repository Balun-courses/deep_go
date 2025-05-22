package main

import (
	"encoding/json"
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		for i := range name {
			person.name[i] = name[i]
		}
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		higherBits := mana & (oneByteMask << fourBitsShift) >> fourBitsShift
		lowerBits := (mana & fourBitsMask) << fourBitsShift
		person.manaHealth[1] |= uint8(lowerBits)
		person.manaHealth[2] |= uint8(higherBits)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		higherBits := (health & (fourBitsMask << oneByteShift)) >> oneByteShift
		lowerBits := health & oneByteMask
		person.manaHealth[1] |= uint8(higherBits)
		person.manaHealth[0] |= uint8(lowerBits)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.respectExperience |= uint8(respect << fourBitsShift)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.strengthLevel |= uint8(strength << fourBitsShift)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.respectExperience |= uint8(experience)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.strengthLevel |= uint8(level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.typePropertyMask |= hasHouse
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.typePropertyMask |= hasGun
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.typePropertyMask |= hasFamily
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		var mask uint8
		switch personType {
		case BuilderGamePersonType:
			mask = typeBuilder
		case BlacksmithGamePersonType:
			mask = typeBlacksmith
		case WarriorGamePersonType:
			mask = typeWarrior
		default:
			return
		}
		person.typePropertyMask |= mask
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

const (
	fourBitsMask  = 0xF
	oneByteMask   = 0xFF
	fourBitsShift = 4
	oneByteShift  = 8

	typeBuilder    = 1  // 0000 0001
	typeBlacksmith = 2  // 0000 0010
	typeWarrior    = 4  // 0000 0100
	hasHouse       = 8  // 0000 1000
	hasGun         = 16 // 0001 0000
	hasFamily      = 32 // 0010 0000
)

type GamePerson struct {
	x, y, z           int32
	gold              uint32
	manaHealth        [3]uint8
	respectExperience uint8
	strengthLevel     uint8
	name              [42]byte
	typePropertyMask  uint8
}

func NewGamePerson(options ...Option) GamePerson {
	gp := GamePerson{}
	for _, option := range options {
		option(&gp)
	}
	return gp
}

func (p *GamePerson) Name() string {
	z := 0
	for z < len(p.name) && p.name[z] != 0 {
		z++
	}
	return string(p.name[:z])
}

func (p *GamePerson) X() int {
	return int(p.x)
}

func (p *GamePerson) Y() int {
	return int(p.y)
}

func (p *GamePerson) Z() int {
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int(p.manaHealth[2])<<fourBitsShift + int(p.manaHealth[1])>>fourBitsShift
}

func (p *GamePerson) Health() int {
	return int(p.manaHealth[1]&fourBitsMask)<<oneByteShift + int(p.manaHealth[0])
}

func (p *GamePerson) Respect() int {
	return int((p.respectExperience & (fourBitsMask << fourBitsShift)) >> fourBitsShift)
}

func (p *GamePerson) Strength() int {
	return int((p.strengthLevel & (fourBitsMask << fourBitsShift)) >> fourBitsShift)
}

func (p *GamePerson) Experience() int {
	return int(p.respectExperience & fourBitsMask)
}

func (p *GamePerson) Level() int {
	return int(p.strengthLevel & fourBitsMask)
}

func (p *GamePerson) HasHouse() bool {
	return p.typePropertyMask&hasHouse != 0
}

func (p *GamePerson) HasGun() bool {
	return p.typePropertyMask&hasGun != 0
}

func (p *GamePerson) HasFamily() bool {
	return p.typePropertyMask&hasFamily != 0
}

func (p *GamePerson) Type() int {
	if p.typePropertyMask&typeBuilder != 0 {
		return BuilderGamePersonType
	}
	if p.typePropertyMask&typeBlacksmith != 0 {
		return BlacksmithGamePersonType
	}
	if p.typePropertyMask&typeWarrior != 0 {
		return WarriorGamePersonType
	}
	return 0
}

type GamePersonForMarshal struct {
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Z          int    `json:"z"`
	Name       string `json:"name"`
	Gold       int    `json:"gold"`
	Health     int    `json:"health"`
	Respect    int    `json:"respect"`
	Strength   int    `json:"strength"`
	Experience int    `json:"experience"`
	Level      int    `json:"level"`
	HasHouse   bool   `json:"has_house"`
	HasFamily  bool   `json:"has_family"`
	HasGun     bool   `json:"has_gun"`
	Type       string `json:"type"`
}

func (p *GamePerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		X          int    `json:"x"`
		Y          int    `json:"y"`
		Z          int    `json:"z"`
		Name       string `json:"name"`
		Gold       int    `json:"gold"`
		Health     int    `json:"health"`
		Respect    int    `json:"respect"`
		Strength   int    `json:"strength"`
		Experience int    `json:"experience"`
		Level      int    `json:"level"`
		HasHouse   bool   `json:"has_house"`
		HasFamily  bool   `json:"has_family"`
		HasGun     bool   `json:"has_gun"`
		Type       string `json:"type"`
	}{
		X:          p.X(),
		Y:          p.Y(),
		Z:          p.Z(),
		Name:       p.Name(),
		Gold:       p.Gold(),
		Health:     p.Health(),
		Respect:    p.Respect(),
		Strength:   p.Strength(),
		Experience: p.Experience(),
		Level:      p.Level(),
		HasHouse:   p.HasHouse(),
		HasFamily:  p.HasFamily(),
		HasGun:     p.HasGun(),
		Type: func() string {
			switch p.Type() {
			case BuilderGamePersonType:
				return "Builder"
			case BlacksmithGamePersonType:
				return "Blacksmith"
			case WarriorGamePersonType:
				return "Warrior"
			default:
				return "Unknown"
			}
		}(),
	})
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
	assert.True(t, person.HasFamily())
	assert.False(t, person.HasGun())
	assert.Equal(t, personType, person.Type())
	rawData, err := json.Marshal(&person)
	assert.NoError(t, err)
	assert.Equal(t,
		string(rawData),
		"{\"x\":-2147483648,\"y\":2147483647,\"z\":0,\"name\":\"aaaaaaaaaaaaa_bbbbbbbbbbbbb_cccccccccccccc\",\"gold\":2147483647,\"health\":1000,\"respect\":10,\"strength\":10,\"experience\":10,\"level\":10,\"has_house\":true,\"has_family\":true,\"has_gun\":false,\"type\":\"Builder\"}")
}
