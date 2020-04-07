package profile

func Merge(global map[string]string, active map[string]string) map[string]string {
	if len(global) == 0 {
		return active
	}

	if len(active) == 0 {
		return global
	}

	for k, v := range active {
		global[k] = v
	}

	return global
}
