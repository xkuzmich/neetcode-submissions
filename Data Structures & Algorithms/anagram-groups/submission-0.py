from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = []

        def counting_letters(word: str) -> tuple(int):
            counts = [0] * 26

            for letter in word:
                counts[ord(letter.lower()) - ord('a')] += 1
            
            return tuple(counts)
        
        result = defaultdict(list)

        for word in strs:
            char_count = counting_letters(word)

            result[char_count].append(word)
        
        return [value for value in result.values()]
        
