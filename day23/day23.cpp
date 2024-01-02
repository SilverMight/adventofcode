#include <string>
#include <vector>
#include <array>
#include <iostream>
#include <fstream>
using std::vector;


struct Pair {
    int x, y;

    friend bool operator==(Pair a, Pair b) {
        if(a.x == b.x && a.y == b.y) {
            return true;
        }
        return false;
    }
    friend Pair operator+(Pair a, Pair b) {
        return Pair{a.x + b.x, a.y + b.y};
    }
};

auto isValid(const vector<std::string>& puzzle, Pair check) -> bool {
    if ((check.x >= 0 && check.x < puzzle.size()) && (check.y >= 0 && check.y < puzzle[0].size())) {
        if(puzzle[check.x][check.y] != '#')
            return true;
    }
    return false;
}

constexpr std::array directions = {Pair{1, 0}, Pair{-1, 0}, Pair{0, 1}, Pair {0, -1}};

auto dfs(const vector<std::string>& puzzle, Pair start, Pair end, std::size_t steps, vector<vector<bool>>& visited, std::size_t& maxResult) -> void {
    if (start == end) {
        maxResult = std::max(steps, maxResult);
    }

    visited[start.x][start.y] = true;

    for(const auto& dir : directions) {
        Pair next = start + dir;

        if (isValid(puzzle, next) && !visited[next.x][next.y]) {
            dfs(puzzle, next, end, steps + 1, visited, maxResult);
        }
    }

    visited[start.x][start.y] = false;
}

int main() {
    vector<std::string> puzzle;
    std::ifstream file{"input.txt"};

    std::string line;
    while (std::getline(file, line)) {
        puzzle.push_back(line);
    }

    auto startIndex = Pair{0, int(puzzle[0].find('.'))};
    auto endIndex = Pair{int(puzzle.size() - 1), int(puzzle[puzzle.size() - 1].find('.'))};

    auto visited = vector<vector<bool>>(puzzle.size(), vector<bool>(puzzle[0].size()));
    std::size_t maxResult = 0;
    dfs(puzzle, startIndex, endIndex, 0, visited, maxResult);

    std::cout << "Max: " << maxResult << '\n';

}
