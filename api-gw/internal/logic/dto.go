package logic

// REST request bodies (clean DTOs for Gin binding)

type EvalDTO struct {
	Expression string             `json:"expression" binding:"required"`
	Variables  map[string]float64 `json:"variables"` // optional
}

type TransformDTO struct {
	Data    []float64 `json:"data"     binding:"required"`
	Expr    string    `json:"expression"`                         // optional
	VarName string    `json:"var_name"`                           // optional
	Op      string    `json:"operation"       binding:"required"` // "MAP" | "FILTER" | "SUM" (case-insensitive) or number
}

type PlanDTO struct {
	Goal     string   `json:"goal"      binding:"required"`
	Hints    []string `json:"hints"`     // optional
	MaxSteps int32    `json:"max_steps"` // optional
}
