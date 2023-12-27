from dataclasses import dataclass
import itertools

@dataclass()
class Vertex:
    x : int; y: int; z: int
    dx : int; dy: int; dz: int


def parametricToCartesian(x: int, y: int, dx: int, dy: int) -> tuple[float, float]:
    slope = dy/dx
    intercept = y + (slope * -x)

    return slope, intercept
    

def timeCheck(xInt:float, x0: int, dx: int) -> bool:
    return ((xInt - x0) / dx) > 0

def getIntersections(vertices : list[Vertex]) -> int:
    numOfIntersections = 0

    #rangeStart, rangeEnd = 7, 27
    rangeStart, rangeEnd = 200000000000000, 400000000000000
    
    for v1, v2 in itertools.combinations(vertices, r=2):
        # only check x and y
        v1_slope, v1_intercept = parametricToCartesian(v1.x, v1.y, v1.dx, v1.dy)
        v2_slope, v2_intercept = parametricToCartesian(v2.x, v2.y, v2.dx, v2.dy)

        
        # solve for interception

        if v1_slope - v2_slope == 0:
            continue


        xIntercept = (v2_intercept - v1_intercept) / (v1_slope - v2_slope)
        yIntercept = xIntercept * v1_slope + v1_intercept

        if not (timeCheck(xIntercept, v1.x, v1.dx) and timeCheck(xIntercept, v2.x, v2.dx)):
            continue
        if rangeStart <= xIntercept <= rangeEnd and rangeStart <= yIntercept <= rangeEnd:
            numOfIntersections += 1

    return numOfIntersections

def parse(puzzle : list[str]) -> list[Vertex]:
    verticesList : list[Vertex] = []
    for line in puzzle:
        numberStr = line.replace('@', ',').split(',')

        args = [int(num) for num in numberStr]
        assert(len(args) == 6)

        verticesList.append(Vertex(*args))

    return verticesList




def main():
    puzzle = open("input.txt", "r").read().splitlines()
    vertices = parse(puzzle)

    print(getIntersections(vertices))


if __name__ == "__main__":
    main()
