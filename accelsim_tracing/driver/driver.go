package driver

import (
	"errors"

	"github.com/sarchlab/akita/v3/tracing"
	"github.com/sarchlab/mgpusim/v3/accelsim_tracing/gpu"
	"github.com/sarchlab/mgpusim/v3/accelsim_tracing/nvidia"
	"github.com/sarchlab/mgpusim/v3/accelsim_tracing/trace"
)

type driver struct {
	gpu             *gpu.GPU
	mericsCollector *collector

	flagReportInstCount bool
	instCountTracers    []instCountTracer
}

func (d *driver) Exec(kl *nvidia.KernelList) error {
	if kl == nil {
		return errors.New("no trace specified")
	}

	for _, e := range kl.TraceExecs {
		if e.Type == "memcopy" {
			// [todo] implement
		} else if e.Type == "kernel" {
			tgReader := trace.NewTGReader(e.FilePath)

			tg := tgReader.ReadAll()
			for _, tb := range tg.ThreadBlocks {
				d.gpu.RunThreadBlock(tb)
			}

			tgReader.CloseFile()
		}
	}

	return nil
}


func (d *driver) reportInstCount() {
	for _, t := range d.instCountTracers {
		d.mericsCollector.Collect(t.cu.Name(), "inst_count", float64(t.tracer.Count))
	}
}

