var res []string

func restoreIpAddresses(s string) []string {
    if len(s) < 4 || len(s) > 12 {
        return []string{}
    }
    res = []string{}
    genIpAddress(s, 0, 0, "")
    return res
}

func genIpAddress(s string, index, part int, ip string) {

    if part == 3 {
        si, err := strconv.Atoi(s[index:])
        ss := strconv.Itoa(si)
        if err != nil || si > 255 || len(ss) != len(s) - index {
            return
        }
        res = append(res, ip + s[index:])
        return
    }
    for i := 1; i < 4; i++{
        if index + i < len(s) {
            si, err := strconv.Atoi(s[index:index+i])
            ss := strconv.Itoa(si)
            if err != nil || si > 255 || len(ss) != i {
                return
            }
            genIpAddress(s, index + i, part + 1, ip  + s[index:index+i] + ".")
        }
    }
    return

}
