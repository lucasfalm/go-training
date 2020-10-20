// https://www.hackerrank.com/challenges/drawing-book/problem
func pageCount(n int32, p int32) int32 {
    if n - p >= -1 && n - p <= 1 {
        return 0
    }

    if p == 1 {
        return 0
    }

    var (
        frontCount = int32(1)
        backCount = n
        totalFront int32
        totalBack int32
    )
   
    for frontCount < p {
        if frontCount + 1 == p {
            frontCount = p + 1
            break
        }

        totalFront++
        frontCount += 2
    }
    
    for backCount > p {
        if backCount - 1 == p {
            frontCount = p - 1
            break
        }

        totalBack++
        backCount -= 2
    }

    if totalBack < totalFront {
        return totalBack
    } else {
        return totalFront
    }
}
