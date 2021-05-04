package main

func index(vs []string, item string) int {
	for i, v := range vs {
		if v == item {
			return i
		}
	}

	return -1
}

func includes(vs []string, item string) bool {
	return index(vs, item) >= 0
}

func any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}

	return false
}

func all(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}

	return true
}

func filter(vs []string, f func(string) bool) []string {
	results := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			results = append(results, v)
		}
	}

	return results
}

func mapItems(vs []string, f func(string) string) []string {
	results := make([]string, len(vs))
	for i, v := range vs {
		results[i] = f(v)
	}
	return results
}

func main() {

}
