// WARNING: this file is generated. DO NOT EDIT

package altsrc

import (
	"flag"

	"github.com/urfave/cli/v2"
)

// BoolFlag is the flag type that wraps cli.BoolFlag to allow
// for other values to be specified
type BoolFlag struct {
	*cli.BoolFlag
	set *flag.FlagSet
}

// NewBoolFlag creates a new BoolFlag
func NewBoolFlag(fl *cli.BoolFlag) *BoolFlag {
	return &BoolFlag{BoolFlag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped BoolFlag.Apply
func (f *BoolFlag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.BoolFlag.Apply(set)
}

// DurationFlag is the flag type that wraps cli.DurationFlag to allow
// for other values to be specified
type DurationFlag struct {
	*cli.DurationFlag
	set *flag.FlagSet
}

// NewDurationFlag creates a new DurationFlag
func NewDurationFlag(fl *cli.DurationFlag) *DurationFlag {
	return &DurationFlag{DurationFlag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped DurationFlag.Apply
func (f *DurationFlag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.DurationFlag.Apply(set)
}

// Float64Flag is the flag type that wraps cli.Float64Flag to allow
// for other values to be specified
type Float64Flag struct {
	*cli.Float64Flag
	set *flag.FlagSet
}

// NewFloat64Flag creates a new Float64Flag
func NewFloat64Flag(fl *cli.Float64Flag) *Float64Flag {
	return &Float64Flag{Float64Flag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped Float64Flag.Apply
func (f *Float64Flag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.Float64Flag.Apply(set)
}

// Float64SliceFlag is the flag type that wraps cli.Float64SliceFlag to allow
// for other values to be specified
type Float64SliceFlag struct {
	*cli.Float64SliceFlag
	set *flag.FlagSet
}

// NewFloat64SliceFlag creates a new Float64SliceFlag
func NewFloat64SliceFlag(fl *cli.Float64SliceFlag) *Float64SliceFlag {
	return &Float64SliceFlag{Float64SliceFlag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped Float64SliceFlag.Apply
func (f *Float64SliceFlag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.Float64SliceFlag.Apply(set)
}

// GenericFlag is the flag type that wraps cli.GenericFlag to allow
// for other values to be specified
type GenericFlag struct {
	*cli.GenericFlag
	set *flag.FlagSet
}

// NewGenericFlag creates a new GenericFlag
func NewGenericFlag(fl *cli.GenericFlag) *GenericFlag {
	return &GenericFlag{GenericFlag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped GenericFlag.Apply
func (f *GenericFlag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.GenericFlag.Apply(set)
}

// IntFlag is the flag type that wraps cli.IntFlag to allow
// for other values to be specified
type IntFlag struct {
	*cli.IntFlag
	set *flag.FlagSet
}

// NewIntFlag creates a new IntFlag
func NewIntFlag(fl *cli.IntFlag) *IntFlag {
	return &IntFlag{IntFlag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped IntFlag.Apply
func (f *IntFlag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.IntFlag.Apply(set)
}

// Int64Flag is the flag type that wraps cli.Int64Flag to allow
// for other values to be specified
type Int64Flag struct {
	*cli.Int64Flag
	set *flag.FlagSet
}

// NewInt64Flag creates a new Int64Flag
func NewInt64Flag(fl *cli.Int64Flag) *Int64Flag {
	return &Int64Flag{Int64Flag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped Int64Flag.Apply
func (f *Int64Flag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.Int64Flag.Apply(set)
}

// Int64SliceFlag is the flag type that wraps cli.Int64SliceFlag to allow
// for other values to be specified
type Int64SliceFlag struct {
	*cli.Int64SliceFlag
	set *flag.FlagSet
}

// NewInt64SliceFlag creates a new Int64SliceFlag
func NewInt64SliceFlag(fl *cli.Int64SliceFlag) *Int64SliceFlag {
	return &Int64SliceFlag{Int64SliceFlag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped Int64SliceFlag.Apply
func (f *Int64SliceFlag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.Int64SliceFlag.Apply(set)
}

// IntSliceFlag is the flag type that wraps cli.IntSliceFlag to allow
// for other values to be specified
type IntSliceFlag struct {
	*cli.IntSliceFlag
	set *flag.FlagSet
}

// NewIntSliceFlag creates a new IntSliceFlag
func NewIntSliceFlag(fl *cli.IntSliceFlag) *IntSliceFlag {
	return &IntSliceFlag{IntSliceFlag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped IntSliceFlag.Apply
func (f *IntSliceFlag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.IntSliceFlag.Apply(set)
}

// PathFlag is the flag type that wraps cli.PathFlag to allow
// for other values to be specified
type PathFlag struct {
	*cli.PathFlag
	set *flag.FlagSet
}

// NewPathFlag creates a new PathFlag
func NewPathFlag(fl *cli.PathFlag) *PathFlag {
	return &PathFlag{PathFlag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped PathFlag.Apply
func (f *PathFlag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.PathFlag.Apply(set)
}

// StringFlag is the flag type that wraps cli.StringFlag to allow
// for other values to be specified
type StringFlag struct {
	*cli.StringFlag
	set *flag.FlagSet
}

// NewStringFlag creates a new StringFlag
func NewStringFlag(fl *cli.StringFlag) *StringFlag {
	return &StringFlag{StringFlag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped StringFlag.Apply
func (f *StringFlag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.StringFlag.Apply(set)
}

// StringSliceFlag is the flag type that wraps cli.StringSliceFlag to allow
// for other values to be specified
type StringSliceFlag struct {
	*cli.StringSliceFlag
	set *flag.FlagSet
}

// NewStringSliceFlag creates a new StringSliceFlag
func NewStringSliceFlag(fl *cli.StringSliceFlag) *StringSliceFlag {
	return &StringSliceFlag{StringSliceFlag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped StringSliceFlag.Apply
func (f *StringSliceFlag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.StringSliceFlag.Apply(set)
}

// UintFlag is the flag type that wraps cli.UintFlag to allow
// for other values to be specified
type UintFlag struct {
	*cli.UintFlag
	set *flag.FlagSet
}

// NewUintFlag creates a new UintFlag
func NewUintFlag(fl *cli.UintFlag) *UintFlag {
	return &UintFlag{UintFlag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped UintFlag.Apply
func (f *UintFlag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.UintFlag.Apply(set)
}

// Uint64Flag is the flag type that wraps cli.Uint64Flag to allow
// for other values to be specified
type Uint64Flag struct {
	*cli.Uint64Flag
	set *flag.FlagSet
}

// NewUint64Flag creates a new Uint64Flag
func NewUint64Flag(fl *cli.Uint64Flag) *Uint64Flag {
	return &Uint64Flag{Uint64Flag: fl, set: nil}
}

// Apply saves the flagSet for later usage calls, then calls
// the wrapped Uint64Flag.Apply
func (f *Uint64Flag) Apply(set *flag.FlagSet) error {
	f.set = set
	return f.Uint64Flag.Apply(set)
}

// vim:ro
