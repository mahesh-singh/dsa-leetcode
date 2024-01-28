class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        
        result  = float('-inf')
        
        prod = 1
        for n in nums:
            prod *=  n
            result = max(result, prod)
            if prod == 0:
                prod = 1
        
        prod = 1
        for i in range(len(nums)-1, -1, -1):
            
            prod *= nums[i]
            result = max(result, prod)
            if prod == 0:
                prod = 1

        return result


