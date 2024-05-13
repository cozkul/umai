package models

type Resource struct {
	ID             uint `gorm:"primaryKey"`
	ParentID       uint
	ParentType     string
	Food           float32
	Metal          float32
	Crystal        float32
	QuantumEssence float32
	Energy         float32
}

func (r *Resource) AddValue(v Resource) {
	r.Food += v.Food
	r.Metal += v.Metal
	r.Crystal += v.Crystal
	r.QuantumEssence += v.QuantumEssence
}

func (r *Resource) Clamp(max Resource) {
	r.Food = clamp(r.Food, max.Food)
	r.Metal = clamp(r.Metal, max.Metal)
	r.Crystal = clamp(r.Crystal, max.Crystal)
	r.QuantumEssence = clamp(r.QuantumEssence, max.QuantumEssence)
}

func clamp(val float32, max float32) float32 {
	if val > max {
		return max
	} else {
		return val
	}
}
