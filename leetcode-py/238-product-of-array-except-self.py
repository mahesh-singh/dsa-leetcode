class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        prefix = [1] * n
        postfix = [1] * n
        output = [1] * n
        for i in range(n):
            if i == 0:
                prefix[i] = 1
            else:
                prefix[i] = prefix[i-1] * nums[i-1]

        for i in range(n    - 1, -1, -1):
            if i == len(nums)-1:
                postfix[i] = 1
            else:
                postfix[i] = postfix[i+1] * nums[i+1]
        
        for i in range(n):
            output[i] = prefix[i] * postfix[i]
        
        return output