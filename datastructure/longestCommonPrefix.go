func longestCommonPrefix(strs []string) string {
    index := 0
    var res []string
    for {
        m := make(map[string]int)
        for _,v := range strs{
	if index == len(v){
		continue
	}
            m[string(v[index])]++
        }
        var val []int
        for _,v := range m{
            val = append(val,v)
        }
        sort.Ints(val)
        if len(val)==0 || val[len(val)-1]<len(strs){
            break
        }
        var maxMatching string
        for k,v := range m {
            if v == val[len(val)-1]{
                maxMatching = k
            }
        }
        res =append(res,maxMatching)
	var up []string
	for _,v := range strs{
		if index == len(v){
			continue
		}
		if string(v[index])==maxMatching {
			up = append(up,v)
		}
	}
	strs = up
	index++
    }
    return strings.Join(res[:], "")
}
