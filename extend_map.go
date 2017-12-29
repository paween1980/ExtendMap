package extend_map

// ExtendMap - extend map a, b to new map, b has more priority
func ExtendMap(a, b map[string]interface{}) (out map[string]interface{}) {
	out = make(map[string]interface{})
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		if out[k] == nil {
			out[k] = v
		} else {
			outV, _ := out[k].(map[string]interface{})
			vv, ok := v.(map[string]interface{})
			if ok {
				out[k] = ExtendMap(outV, vv)
			} else {
				out[k] = v
			}
		}
	}
	return
}
