package problems

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	record := make(map[[26]int8][]string)

	for _, str := range strs {
		temp := string2ints(str)
		record[temp] = append(record[temp], str)
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}

func string2ints(s string) [26]int8 {
	temp := [26]int8{}
	for _, v := range s {
		temp[v-'a'] += 1
	}
	return temp
}
