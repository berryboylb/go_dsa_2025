package dsa_11_11_25

func solution(s string, p string) bool {
    slen, plen := len(s), len(p)
    dot, star := byte('.'), byte('*')

    matrix := make([][]bool, slen+1)
    for i := range matrix {
        matrix[i] = make([]bool, plen+1)
    }

    matrix[0][0] = true

    // Initialize for patterns like a*, a*b*, etc.
    for j := 1; j <= plen; j++ {
        if p[j-1] == star && j >= 2 {
            matrix[0][j] = matrix[0][j-2]
        }
    }

    for i := 1; i <= slen; i++ {
        for j := 1; j <= plen; j++ {
            if p[j-1] == s[i-1] || p[j-1] == dot {
                matrix[i][j] = matrix[i-1][j-1]
            } else if p[j-1] == star {
                // zero occurrence
                matrix[i][j] = matrix[i][j-2]
                // one or more occurrence
                if p[j-2] == s[i-1] || p[j-2] == dot {
                    matrix[i][j] = matrix[i][j] || matrix[i-1][j]
                }
            }
        }
    }

    return matrix[slen][plen]
}


func Main() {

}
