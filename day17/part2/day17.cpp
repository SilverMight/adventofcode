#include <bits/stdc++.h>
#include <iostream>

static inline int mod(int a, int b) { 
    return (a % b + b) % b;
}

using Matrix = std::vector<std::vector<int>>;

enum Direction {
    Up,
    Right,
    Down,
    Left,
};


struct Point {
    int X, Y;

    Point(int x, int y) : X{x}, Y{y} {

    }

    Point operator+ (const Point& p) const {
        return Point{this->X + p.X, this->Y + p.Y};
    }
};

static const std::array<Point, 4> directions = {
    Point{-1, 0},
    Point{0, 1},
    Point{1, 0},
    Point{0, -1},
};

struct Tile {
    Point point = {0, 0};
    Direction dir = Up;
    int curMovesInSameDirection = 0;
    int heatLoss = 0;

    bool operator< (const Tile& p) const {
        return this->heatLoss < p.heatLoss;
    }
    bool operator> (const Tile& p) const {
        return p < *this; 
    }
    friend bool operator== (const Tile& p1, const Tile& p2) {
        return (p1.point.X == p2.point.X) && (p1.point.Y == p2.point.Y) && (p1.dir == p2.dir) && (p1.curMovesInSameDirection == p2.curMovesInSameDirection);
    }
};

struct TileHash {
    std::size_t operator()(const Tile& tile) const {
        std::size_t hashValue = 17;  
        hashValue = hashValue * 31 + static_cast<std::size_t>(tile.point.X);
        hashValue = hashValue * 31 + static_cast<std::size_t>(tile.point.Y);
        hashValue = hashValue * 31 + static_cast<std::size_t>(tile.dir);
        hashValue = hashValue * 31 + static_cast<std::size_t>(tile.curMovesInSameDirection);
        return hashValue;
    }
};



Tile moveStraight(const Tile& curTile) {
    Tile next;
    next.dir = curTile.dir;
    next.point = curTile.point + directions[next.dir];
    next.curMovesInSameDirection = curTile.curMovesInSameDirection + 1;

    return next;
}

Tile moveRight(const Tile& curTile) {
    Tile next;
    next.dir = static_cast<Direction>(mod(curTile.dir + 1, 4));
    next.point = curTile.point + directions[next.dir];
    next.curMovesInSameDirection = 1;

    return next;
}

Tile moveLeft(const Tile& curTile) {
    Tile next;
    next.dir = static_cast<Direction>(mod(curTile.dir - 1, 4));
    next.point = curTile.point + directions[next.dir];
    next.curMovesInSameDirection = 1;

    return next;
}

int djikstras(const Matrix& puzzle) {
    std::priority_queue<Tile, std::vector<Tile>, std::greater<Tile>> queue;
    std::unordered_set<Tile, TileHash> visited;

    Tile tileStart;
    tileStart.point = Point{0, 0};
    tileStart.dir = Down;
    queue.push(tileStart);
    tileStart.dir = Right;
    queue.push(tileStart);
    



    auto isValid = [&](const Tile& tile) -> bool {
        if (tile.point.X < 0 || tile.point.X >= puzzle.size() ||
            tile.point.Y < 0 || tile.point.Y >= puzzle[0].size()) {
            return false;
        }
        return true;
    };

    while (!queue.empty()) {
        Tile curr = queue.top();
        queue.pop();
        if(visited.count(curr) != 0) {
            continue;
        }

        //std::printf("Curr: %d, %d, Heat: %d, Same dir: %d\n", curr.point.X, curr.point.Y, curr.heatLoss, curr.curMovesInSameDirection);
        visited.insert(curr);
        if((curr.point.X == (puzzle.size() - 1)) && curr.point.Y == (puzzle[0].size() - 1)) {
            if(curr.curMovesInSameDirection > 3) {
                return curr.heatLoss; 
            }
        }

        // Straight (same direction)
        if(curr.curMovesInSameDirection < 10) {
            Tile next = moveStraight(curr);
            if(isValid(next)) {
                next.heatLoss = curr.heatLoss +  puzzle[next.point.X][next.point.Y];
                queue.push(next);
            }
        }


        // Left
        if(curr.curMovesInSameDirection > 3) {
            Tile nextRight = moveRight(curr);
            if(isValid(nextRight)) {
                nextRight.heatLoss = curr.heatLoss +  puzzle[nextRight.point.X][nextRight.point.Y];
                queue.push(nextRight);
            }

            Tile nextLeft = moveLeft(curr);
            if(isValid(nextLeft)) {
                nextLeft.heatLoss = curr.heatLoss +  puzzle[nextLeft.point.X][nextLeft.point.Y];
                queue.push(nextLeft);
            }
        }
    }
    return 0;

}

int main() { 
    Matrix puzzle;

    std::string line;
    while(std::getline(std::cin, line)) {
        puzzle.push_back({});
        for(char c : line) {
            puzzle.back().push_back(c - '0');
        }
    }

    std::cout << djikstras(puzzle) << "\n";

}
