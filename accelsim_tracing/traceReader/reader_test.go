package tracereader

import (
	"strings"
	"testing"

	"github.com/sarchlab/accelsimtracing/nvidia"
)

func TestBuildExecFromText(t *testing.T) {
	reader := new(TraceReader)

	// format:  H2D or D2H, start, length
	t.Run("h2d", func(t *testing.T) {
		e := reader.buildExecFromText("MemcpyHtoD,0x00007fc841500000,32768")
		if e.Direction != nvidia.H2D {
			t.Errorf("Expected H2D, got %v", e.Direction)
		}
		if e.Address != 0x00007fc841500000 {
			t.Errorf("Expected 0x00007fc841500000, got %v", e.Address)
		}
		if e.Length != 32768 {
			t.Errorf("Expected 32768, got %v", e.Length)
		}
	})

	// // format:  kernel_name
	// t.Run("kernel", func(t *testing.T) {
	// 	e := reader.generateExecFromText("kernel-1.traceg")
	// 	if e.(*execKernel).filename != "kernel-1.traceg" {
	// 		t.Errorf("Expected kernel-1.traceg, got %v", e.filename)
	// 	}
	// })
}

var threadblockExample string = `#BEGIN_TB

thread block = 1,0,0

warp = 0
insts = 12
0000 ffffffff 1 R1 IMAD.MOV.U32 2 R255 R255 0 
0010 ffffffff 1 R6 S2R 0 0 
0020 ffffffff 1 R3 S2R 0 0 
0030 ffffffff 1 R6 IMAD 2 R6 R3 0 
0040 ffffffff 0 ISETP.GE.AND 1 R6 0 
0050 00000000 0 EXIT 0 0 
0060 ffffffff 1 R4 IADD3 2 R6 R255 0 
0070 ffffffff 0 ULDC.64 0 0 
0080 ffffffff 1 R5 LEA.HI.X.SX32 1 R6 0 
0090 ffffffff 1 R0 LDG.E.U8 1 R4 1 1 0x7fc841520200 1 
00a0 ffffffff 0 ISETP.NE.AND 2 R0 R255 0 
00b0 ffffffff 0 EXIT 0 0 

warp = 1
insts = 2
0000 ffffffff 1 R1 IMAD.MOV.U32 2 R255 R255 0 
0010 ffffffff 1 R6 S2R 0 0 

#END_TB`

func TestBuildThreadblock(t *testing.T) {
	lines := strings.Split(threadblockExample, "\n")
	tb := buildThreadblock(lines)

	if len(tb.warps) != 2 {
		t.Errorf("Expected 2 warps, got %v", len(tb.warps))
	}
	if tb.threadblockDim != [3](int){1, 0, 0} {
		t.Errorf("Expected (1,0,0), got %v", tb.threadblockDim)
	}
	if tb.warps[0].instsCount != 12 || len(tb.warps[0].instructions) != 12 {
		t.Errorf("Expected 12 instructions, got %v", len(tb.warps[0].instructions))
	}
	if tb.warps[1].instsCount != 2 || len(tb.warps[1].instructions) != 2 {
		t.Errorf("Expected 2 instructions, got %v", len(tb.warps[1].instructions))
	}
	if tb.warps[0].warpID != 0 || tb.warps[1].warpID != 1 {
		t.Errorf("Expected 0 and 1, got %v and %v", tb.warps[0].warpID, tb.warps[1].warpID)
	}
}