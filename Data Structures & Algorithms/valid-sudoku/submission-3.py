class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
    

        for r in board:
            row = set()
            for c in r:
                if c != ".":
                    if c in row:
                        return False
                    row.add(c)
        
        for c in range(len(board[0])):
            colum = set()
            for r in range(len(board)):
                if board[r][c] == ".":
                    continue
                if board[r][c] != ".":
                    if board[r][c] in colum:
                        return False
                    colum.add(board[r][c])
        
        for box in range(9):
            square = set()
            for i in range(3):
                for j in range(3):
                    r = (box // 3) * 3 + i
                    c = (box % 3) * 3 + j
                    if board[r][c] == ".":
                        continue
                    if board[r][c] in square:
                        return False
                    square.add(board[r][c])

        return True

