package questions

import "math"

/**
while (cntPerfectSquares.size() <= n)
        {
            int m = cntPerfectSquares.size();
            int cntSquares = INT_MAX;
            for (int i = 1; i*i <= m; i++)
            {
                cntSquares = min(cntSquares, cntPerfectSquares[m - i*i] + 1);
            }

            cntPerfectSquares.push_back(cntSquares);
        }

        return cntPerfectSquares[n];
*/
func numSquares(n int) int {
	if n <= 0 {
		return 0
	}
	countPerfectSquares := make([]int, 0)

	for len(countPerfectSquares) <= n {
		m := len(countPerfectSquares)
		countSquares := math.MaxInt32
		for i := 1; i*i <= m; i++ {
			countSquares = int(math.Min(float64(countSquares), float64(countPerfectSquares[m-i*i]-1)))
		}
		countPerfectSquares = append(countPerfectSquares, countSquares)
	}
	return countPerfectSquares[n]
}
