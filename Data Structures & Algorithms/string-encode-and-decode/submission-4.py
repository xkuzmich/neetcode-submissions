class Solution:

    def encode(self, strs: List[str]) -> str:
        enc_str = []
        for word in strs:
            enc_str.append(str(len(word)))
            enc_str.append("#")
            enc_str.append(word)
        return "".join(enc_str)

    def decode(self, s: str) -> List[str]:
        print(s)
        result = []
        i = 0
        while i < len(s):
            if s[i+1] == "#":
                str_len = int(s[i])
                i += 1
            else:
                n = []
                while s[i] != "#":
                    n.append(s[i])
                    i += 1
                str_len = int("".join(n))

            result.append(s[i + 1: i + 1 + str_len])
            i += str_len + 1
        
        return result
