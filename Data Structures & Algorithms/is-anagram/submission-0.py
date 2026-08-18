from collections import defaultdict
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        cache = defaultdict(int)
        for i in range(len(s)):
            cache[s[i]] += 1
            cache[t[i]] -= 1
    
        for v in cache.values():
            if v != 0:
                return False
        
        return True
