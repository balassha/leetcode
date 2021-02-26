func romanToInt(s string) int {
    num := make(map[string]int)
    num["I"] = 1
    num["V"] = 5
    num["X"] = 10
    num["L"] = 50
    num["C"] = 100
    num["D"] = 500
    num["M"] = 1000
    
    el := strings.Split(s,"")
    count := 0
    if len(el)>0{
        for c:=len(el)-1;c>=0;c--{
            if (el[c]=="V" || el[c]=="X") && (c-1>=0) && (el[c-1]=="I") {
                if el[c]=="V"{
                    count += 4
                    c--
                    continue
                }else{
                    count += 9
                    c--
                    continue
                }
            }
            if (el[c]=="L" || el[c]=="C") && (c-1>=0) && (el[c-1]=="X") {
                if el[c]=="L"{
                    count += 40
                    c--
                    continue
                }else{
                    count += 90
                    c--
                    continue
                }
            }
            if (el[c]=="D" || el[c]=="M") && (c-1>=0) && (el[c-1]=="C") {
                if el[c]=="D"{
                    count += 400
                    c--
                    continue
                }else{
                    count += 900
                    c--
                    continue
                }
            }
            count +=num[el[c]]
        }   
    }else{
        return 0
    }
    return count
}
