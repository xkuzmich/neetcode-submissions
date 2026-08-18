class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        
        all_power, zero_cnt = 1, 0

        for elem in nums:
            if elem == 0:
                zero_cnt += 1
            else:
                all_power *= elem
        
        if zero_cnt > 1:
            return [0] * len(nums)
        
        result = [0] * len(nums) 

        for i, elem in enumerate(nums):
            if zero_cnt: 
                result[i] = 0 if elem else all_power
            else:
                result[i] = all_power // elem

        return result