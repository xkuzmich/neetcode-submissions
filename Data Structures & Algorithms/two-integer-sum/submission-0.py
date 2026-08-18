from collections import defaultdict
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        cache = defaultdict(int)

        for i in range(len(nums)):
            second = target - nums[i]

            if second in cache:
                return [cache[second], i]
            else:
                cache[nums[i]] = i