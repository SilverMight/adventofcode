from typing import List
import sys

sys.setrecursionlimit(1000000)

directions = [(-1, 0), (0, 1), (1, 0), (0, -1)]

def isValid(next : tuple, puzzle : List[str]):
    (x, y) = next
    return x >= 0 and x < len(puzzle) and y >= 0 and y < len(puzzle[0]) and puzzle[x][y] != '#'

def dfs(x: int, y: int, end: tuple[int, int], puzzle : List[str], visited : set, steps :int, results: List[int]):

    if (x, y) == end:
        results.append(steps)
        return 
        
    visited.add((x, y))

    for dx, dy in directions:
        next = (x + dx, y + dy)
        if isValid(next, puzzle) and next not in visited:
            dfs(next[0], next[1], end, puzzle, visited, steps + 1, results)


    visited.remove((x, y))

def main(): 
    puzzle = open("input.txt", "r").read().splitlines()

    rows = len(puzzle)

    start = (0, puzzle[0].index('.'))
    end = (rows - 1, puzzle[rows - 1].index('.'))

    results = []
    dfs(*start, end, puzzle, set(), 0, results)
    print(max(results))

if __name__ == "__main__":
    main()
