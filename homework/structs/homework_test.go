package main

import (
	"bytes"
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		// копируем не более 42 байт имени прямо в структуру,
		// чтобы символы хранились рядом с остальными данными
		n := len(name)
		if n > len(person.name) {
			n = len(person.name)
		}
		copy(person.name[:], name[:n])
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
		person.gold = int32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		v := uint16(mana)
		// мана занимает 10 бит: младшие 8 бит в flags[3], старшие 2 — в flags[4](0-1)
		person.flags[3] = byte(v)
		person.flags[4] = person.flags[4] & 0xFC | byte(v>>8)&0x03
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		v := uint16(health)
		// здоровье занимает 10 бит: младшие 6 бит в flags[4](2-7), старшие 4 — в flags[5](0-3)
		person.flags[4] = person.flags[4]&0x03 | byte(v)&0x3F<<2
		person.flags[5] = person.flags[5]&0xF0 | byte(v>>6)&0x0F
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		// уважение занимает 4 бита: flags[0](0-3)
		person.flags[0] = person.flags[0]&0xF0 | byte(respect)&0x0F
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		// сила занимает 4 бита: flags[0](4-7)
		person.flags[0] = person.flags[0]&0x0F | byte(strength)&0x0F<<4
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		// опыт занимает 4 бита: flags[1](0-3)
		person.flags[1] = person.flags[1]&0xF0 | byte(experience)&0x0F
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		// уровень занимает 4 бита: flags[1](4-7)
		person.flags[1] = person.flags[1]&0x0F | byte(level)&0x0F<<4
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		// флаг дома: бит 0 в flags[2]
		person.flags[2] |= 0x01
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		// флаг оружия: бит 1 в flags[2]
		person.flags[2] |= 0x02
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		// флаг семьи: бит 2 в flags[2]
		person.flags[2] |= 0x04
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		// тип занимает 2 бита: flags[2](3-4)
		person.flags[2] = person.flags[2]&0x1F | (byte(personType)&0x03)<<3
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

// GamePerson занимает ровно 64 байта (размер кэш-линии):
//
//	x, y, z, gold int32      16 байт (смещения 0-15, выравнивание 4)
//	flags         [6]byte     6 байт (смещения 16-21, побитовая упаковка):
//	  [0]: уважение(0-3) | сила(4-7)
//	  [1]: опыт(0-3) | уровень(4-7)
//	  [2]: дом(0) | оружие(1) | семья(2) | тип(3-4)
//	  [3]: мана(0-7, младшие 8 из 10 бит)
//	  [4]: мана(0-1, старшие 2 бита) | здоровье(2-7, младшие 6 из 10 бит)
//	  [5]: здоровье(0-3, старшие 4 бита)
//	name          [42]byte    42 байта (смещения 22-63)
//
// Поля int32 идут первыми, чтобы не было выравнивающего padding между
// байтовым массивом имени и координатами.
type GamePerson struct {
	x    int32
	y    int32
	z    int32
	gold int32
	flags [6]byte
	name  [42]byte
}

func NewGamePerson(options ...Option) GamePerson {
	person := GamePerson{}
	for _, option := range options {
		option(&person)
	}
	return person
}

func (p *GamePerson) Name() string {
	// имя латиницей, terminates null-байтом; если имя заняло все 42 байта — берем их все
	n := bytes.IndexByte(p.name[:], 0)
	if n < 0 {
		n = len(p.name)
	}
	return unsafe.String(unsafe.SliceData(p.name[:]), n)
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
	return int(uint16(p.flags[3]) | uint16(p.flags[4]&0x03)<<8)
}

func (p *GamePerson) Health() int {
	return int(uint16(p.flags[4]>>2)&0x3F) | int(uint16(p.flags[5]&0x0F))<<6
}

func (p *GamePerson) Respect() int {
	return int(p.flags[0] & 0x0F)
}

func (p *GamePerson) Strength() int {
	return int(p.flags[0] >> 4)
}

func (p *GamePerson) Experience() int {
	return int(p.flags[1] & 0x0F)
}

func (p *GamePerson) Level() int {
	return int(p.flags[1] >> 4)
}

func (p *GamePerson) HasHouse() bool {
	return p.flags[2]&0x01 != 0
}

func (p *GamePerson) HasGun() bool {
	return p.flags[2]&0x02 != 0
}

func (p *GamePerson) HasFamilty() bool {
	return p.flags[2]&0x04 != 0
}

func (p *GamePerson) Type() int {
	return int((p.flags[2] >> 3) & 0x03)
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

func TestGamePersonRoundsToCacheLine(t *testing.T) {
	// структура должна умещаться ровно в одну кэш-линию,
	// чтобы несколько объектов помещались в одну строку без пересечений
	size := unsafe.Sizeof(GamePerson{})
	assert.Equal(t, uintptr(64), size)
}

func TestGamePersonDefaults(t *testing.T) {
	person := NewGamePerson()
	assert.Equal(t, "", person.Name())
	assert.Equal(t, 0, person.X())
	assert.Equal(t, 0, person.Y())
	assert.Equal(t, 0, person.Z())
	assert.Equal(t, 0, person.Gold())
	assert.Equal(t, 0, person.Mana())
	assert.Equal(t, 0, person.Health())
	assert.Equal(t, 0, person.Respect())
	assert.Equal(t, 0, person.Strength())
	assert.Equal(t, 0, person.Experience())
	assert.Equal(t, 0, person.Level())
	assert.False(t, person.HasHouse())
	assert.False(t, person.HasGun())
	assert.False(t, person.HasFamilty())
	assert.Equal(t, 0, person.Type())
}

func TestGamePersonEmptyNameAndGun(t *testing.T) {
	person := NewGamePerson(
		WithName(""),
		WithGun(),
		WithType(WarriorGamePersonType),
		WithLevel(1),
	)
	assert.Equal(t, "", person.Name())
	assert.True(t, person.HasGun())
	assert.False(t, person.HasHouse())
	assert.False(t, person.HasFamilty())
	assert.Equal(t, WarriorGamePersonType, person.Type())
	assert.Equal(t, 1, person.Level())
}

func TestGamePersonMaxNameLength(t *testing.T) {
	// ровно 42 символа — максимум, который помещается в [42]byte
	name := make([]byte, 42)
	for i := range name {
		name[i] = 'a' + byte(i%26)
	}
	lastName := string(name)

	person := NewGamePerson(WithName(lastName))
	assert.Equal(t, lastName, person.Name())

	// имя длиннее 42 символов обрезается, не ломая структуру
	longName := lastName + "extra"
	overflown := NewGamePerson(WithName(longName))
	assert.Equal(t, lastName, overflown.Name())
}

func TestGamePersonNegativeCoordinatesAndMinGold(t *testing.T) {
	person := NewGamePerson(
		WithCoordinates(-2_000_000_000, -2_000_000_000, -2_000_000_000),
		WithGold(0),
	)
	assert.Equal(t, -2_000_000_000, person.X())
	assert.Equal(t, -2_000_000_000, person.Y())
	assert.Equal(t, -2_000_000_000, person.Z())
	assert.Equal(t, 0, person.Gold())
}

func TestGamePersonBoundaryValues(t *testing.T) {
	// проверяем границы всех полей, чтобы убедиться, что побитовая
	// упаковка не переполняется (10 бит маны/здоровья, 4 бита мелких значений)
	person := NewGamePerson(
		WithCoordinates(math.MinInt32, math.MinInt32, math.MinInt32),
		WithGold(math.MaxInt32),
		WithMana(1000),
		WithHealth(1000),
		WithRespect(10),
		WithStrength(10),
		WithExperience(10),
		WithLevel(10),
		WithHouse(),
		WithGun(),
		WithFamily(),
		WithType(BlacksmithGamePersonType),
	)
	assert.Equal(t, math.MinInt32, person.X())
	assert.Equal(t, math.MinInt32, person.Y())
	assert.Equal(t, math.MinInt32, person.Z())
	assert.Equal(t, math.MaxInt32, person.Gold())
	assert.Equal(t, 1000, person.Mana())
	assert.Equal(t, 1000, person.Health())
	assert.Equal(t, 10, person.Respect())
	assert.Equal(t, 10, person.Strength())
	assert.Equal(t, 10, person.Experience())
	assert.Equal(t, 10, person.Level())
	assert.True(t, person.HasHouse())
	assert.True(t, person.HasGun())
	assert.True(t, person.HasFamilty())
	assert.Equal(t, BlacksmithGamePersonType, person.Type())
}

func TestGamePersonOptionsOrderIndependence(t *testing.T) {
	// опции не должны зависеть от порядка применения
	first := NewGamePerson(
		WithName("abc"),
		WithCoordinates(1, 2, 3),
		WithGold(5),
		WithMana(7),
		WithHealth(9),
		WithRespect(3),
		WithStrength(4),
		WithExperience(5),
		WithLevel(6),
		WithHouse(),
		WithGun(),
		WithFamily(),
		WithType(WarriorGamePersonType),
	)
	second := NewGamePerson(
		WithType(WarriorGamePersonType),
		WithFamily(),
		WithGun(),
		WithHouse(),
		WithLevel(6),
		WithExperience(5),
		WithStrength(4),
		WithRespect(3),
		WithHealth(9),
		WithMana(7),
		WithGold(5),
		WithCoordinates(1, 2, 3),
		WithName("abc"),
	)
	assert.Equal(t, first.Name(), second.Name())
	assert.Equal(t, first.Gold(), second.Gold())
	assert.Equal(t, first.Mana(), second.Mana())
	assert.Equal(t, first.Health(), second.Health())
	assert.Equal(t, first.Type(), second.Type())
	assert.Equal(t, first.HasHouse(), second.HasHouse())
}
