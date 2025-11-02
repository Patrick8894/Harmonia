package engine

type PiDTO struct {
	Samples int64 `json:"samples" binding:"required,min=1"`
}

type MatrixDTO struct {
	Rows int32     `json:"rows" binding:"required,min=1"`
	Cols int32     `json:"cols" binding:"required,min=1"`
	Data []float64 `json:"data" binding:"required"`
}

type MatMulDTO struct {
	A MatrixDTO `json:"a" binding:"required"`
	B MatrixDTO `json:"b" binding:"required"`
}

type StatsDTO struct {
	Data   []float64 `json:"data" binding:"required"`
	Sample *bool     `json:"sample"` // optional; default to true if nil
}
