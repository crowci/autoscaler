package engine

import "github.com/crowci/crow/v3/crow-go/crow"

func countTasksByLabel(jobs []crow.Task, labelKey, labelValue string) int {
	count := 0
	for _, job := range jobs {
		val, exists := job.Labels[labelKey]
		if exists && val == labelValue {
			count++
		}
	}
	return count
}
