class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if not nums:
            return 0

        nums_s = set(nums)

        res = 0 

        for num in nums:
            if num - 1 in nums_s:
                continue
            lenght = 0
            current = num
            while current in nums_s:
                current += 1
                lenght += 1
            
            res = max(res, lenght)
        
        return res
