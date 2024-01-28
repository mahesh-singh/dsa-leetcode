class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        maxSum = nums[0]
        startIndex = 0
        endIndex = 0
        sum = 0

        for i in range(len(nums)):
            if sum ==0:
                startIndex =i
            sum += nums[i]
            if sum > maxSum:
                maxSum = sum
                endIndex = i

            if sum < 0:
                sum = 0



        return maxSum
