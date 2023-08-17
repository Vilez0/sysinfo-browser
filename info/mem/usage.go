package mem

func UsageMB() int {
	return Total() - Available()
}

func UsagePercent() int {
	return int((float64(UsageMB()) / float64(Total())) * 100) 
}
