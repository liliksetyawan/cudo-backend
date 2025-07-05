package dto

type CheckFraud struct {
	Transaction     []interface{}   `json:"transaction"`
	ProcessMetadata ProcessMetadata `json:"process_metadata"`
}

type ProcessMetadata struct {
	TotalTransactionAnalyzed int64         `json:"total_transaction_analyzed"`
	DurationMs               int64         `json:"duration_ms"`
	ParallelTasks            ParallelTasks `json:"parallel_tasks"`
}

type ParallelTasks struct {
	FrequencyAnalysisDuration int64 `json:"frequency_analysis_duration"`
	AmountAnalysisDuration    int64 `json:"amount_analysis_duration"`
	PatternAnalysisDuration   int64 `json:"pattern_analysis_duration"`
}
