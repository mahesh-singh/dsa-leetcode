class Solution:
    def myPow(self, x: float, n: int) -> float:
        
        def pow(x, n):
            if n == 0:
               return 1
            if n == 1:
                return x
            
            tmp = pow(x, n//2)
            res = tmp * tmp 
            res = res * x if n % 2 else res
            return res
        
        res = pow(x, abs(n))

        if n <0:
            res = 1/res
        
        return res