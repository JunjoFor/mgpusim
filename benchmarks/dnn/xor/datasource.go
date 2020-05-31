package xor

import "gitlab.com/akita/dnn/tensor"

// DataSource defines the training dataset for the xor operation.
type DataSource struct {
	allData   []float64
	allLabel  []int
	imageSize int
	currPtr   int
}

// NewDataSource creates a new XOR datasource
func NewDataSource() *DataSource {
	ds := &DataSource{
		imageSize: 2,
	}
	ds.allData = []float64{
		0, 0,
		0, 1,
		1, 0,
		1, 1,
	}
	ds.allLabel = []int{
		0, 1, 1, 0,
	}
	return ds
}

// NextBatch returns the next batch data.
func (ds *DataSource) NextBatch(batchSize int) (
	data tensor.Tensor,
	label []int,
) {
	start := ds.currPtr
	end := start + batchSize

	if end > len(ds.allLabel) {
		end = len(ds.allLabel)
	}

	rawData := ds.allData[start*ds.imageSize : end*ds.imageSize]
	data = &tensor.SimpleTensor{}
	data.Init(rawData, []int{end - start, ds.imageSize})

	label = ds.allLabel[start:end]

	ds.currPtr = end

	return data, label
}

// Rewind moves the pointer to the beginning of the training set.
func (ds *DataSource) Rewind() {
	ds.currPtr = 0
}
