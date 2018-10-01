package imgui

// #include <stdlib.h>
import "C"
import "unsafe"

// AllocatedGlyphRanges are GlyphRanges dynamically allocated by the application.
// Such ranges need to be freed when they are no longer in use to avoid resource leak.
type AllocatedGlyphRanges GlyphRanges

// Free releases the underlying memory of the ranges.
// Call this method when the ranges are no longer in use.
func (ranges *AllocatedGlyphRanges) Free() {
	C.free(unsafe.Pointer(*ranges))
	*ranges = 0
}

// GlyphRangesBuilder can be used to create a new, combined, set of ranges.
type GlyphRangesBuilder struct {
	ranges []glyphRange
}

// Build combines all the currently registered ranges and creates a new instance.
// The returned ranges object needs to be explicitly freed in order to release resources.
func (builder *GlyphRangesBuilder) Build() AllocatedGlyphRanges {
	raw := C.malloc(C.size_t(2 * ((len(builder.ranges) * 2) + 1)))
	rawSlice := (*[1 << 30]uint16)(unsafe.Pointer(raw))[:]
	outIndex := 0
	for _, r := range builder.ranges {
		rawSlice[outIndex+0] = r.from
		rawSlice[outIndex+1] = r.to
		outIndex += 2
	}
	rawSlice[outIndex] = 0
	return AllocatedGlyphRanges(uintptr(raw))
}

// AddExisting adds the given set of ranges to the builder.
// The provided ranges are immediately extracted.
func (builder *GlyphRangesBuilder) AddExisting(ranges ...GlyphRanges) {
	for _, rawRange := range ranges {
		builder.ranges = append(builder.ranges, rawRange.extract()...)
	}
}

// Add extends the builder with the given range (inclusive).
// from must be smaller, or equal to, to - otherwise the range is ignored.
func (builder *GlyphRangesBuilder) Add(from, to rune) {
	if from > to {
		return
	}
	builder.ranges = append(builder.ranges, glyphRange{from: uint16(from), to: uint16(to)})
}
