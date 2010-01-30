//	TODO:	Storing and retrieving of pointers

package vm

import . "container/vector"
import "unsafe"

type Buffer struct {
	IntVector
}

func (b *Buffer) Init(length int)						{ b.Resize(length, length) }
func (b *Buffer) Len() int								{ return b.Len() }
func (b *Buffer) Slice(i, j int) *Buffer				{ return &Buffer{*b.IntVector.Slice(0, b.Len() - 1)} }
func (b *Buffer) Clone() *Buffer						{ return b.Slice(0, b.Len() - 1) }
func (b *Buffer) At(i int) int							{ return b.IntVector.At(i) }
func (b *Buffer) Set(i, x int)							{ b.Set(i, x) }
func (b *Buffer) Add(i, x int)							{ b.Set(i, b.At(i) + x) }
func (b *Buffer) Subtract(i, x int)						{ b.Set(i, b.At(i) - x) }
func (b *Buffer) Multiply(i, x int)						{ b.Set(i, b.At(i) * x) }
func (b *Buffer) Divide(i, x int)						{ b.Set(i, b.At(i) / x) }
func (b *Buffer) And(i, x int)							{ b.Set(i, b.At(i) & x) }
func (b *Buffer) Or(i, x int)							{ b.Set(i, b.At(i) | x) }
func (b *Buffer) Xor(i, x int)							{ b.Set(i, b.At(i) ^ x) }
func (b *Buffer) Increment(i int)						{ b.Add(i, 1) }
func (b *Buffer) Decrement(i int)						{ b.Subtract(i, 1) }
func (b *Buffer) Equal(i, x int) bool					{ return b.At(i) == x }
func (b *Buffer) LessThan(i, x int)	bool				{ return b.At(i) < x }
func (b *Buffer) GreaterThan(i, x int) bool				{ return b.At(i) > x }
func (b *Buffer) Copy(i, j int)							{ b.Set(i, b.At(j)) }
func (b *Buffer) GetBuffer(i int) *Buffer				{ return (*Buffer)(unsafe.Pointer(uintptr(b.At(i)))) }
func (b *Buffer) PutBuffer(i int, p *Buffer)			{ b.Set(i, int(uintptr(unsafe.Pointer(p)))) }