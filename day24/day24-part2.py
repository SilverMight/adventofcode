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

def getIntersections(vertices : list[Vertex]) -> tuple[int, int, int]:
    #rangeStart, rangeEnd = 7, 27
    rangeStart, rangeEnd = 200000000000000, 400000000000000
    
    xSet = None
    ySet = None
    zSet = None 

    for v1, v2 in itertools.combinations(vertices, r=2):
        # only check x and y

        if v1.dx == v2.dx:
            distanceDifference = v2.x - v1.x
            curXSet = set()
            for hailstoneV in range(-1000, 1000):
                if hailstoneV - v1.dx == 0:
                    continue
                if distanceDifference % (hailstoneV - v1.dx) == 0:
                    curXSet.add(hailstoneV)
            if xSet:
                xSet &= curXSet
            else:
                xSet = curXSet.copy()

        if v1.dy == v2.dy:
            distanceDifference = v2.y - v1.y
            curYSet = set()
            for hailstoneV in range(-1000, 1000):
                if hailstoneV - v1.dy == 0:
                    continue
                if distanceDifference % (hailstoneV - v1.dy) == 0:
                    curYSet.add(hailstoneV)
            if ySet:
                ySet &= curYSet
            else:
                ySet = curYSet.copy()
        if v1.dz == v2.dz:
            distanceDifference = v2.z - v1.z
            curZSet = set()
            for hailstoneV in range(-1000, 1000):
                if hailstoneV - v1.dz == 0:
                    continue
                if distanceDifference % (hailstoneV - v1.dz) == 0:
                    curZSet.add(hailstoneV)
            if zSet:
                zSet &= curZSet
            else:
                zSet = curZSet.copy()

     
    return xSet.pop(), ySet.pop(), zSet.pop()


# A little confusing. Need to review this.
def part2(vertices : list[Vertex]):
    rockDx, rockDy, rockDz = getIntersections(vertices)
    
    slopeA = (vertices[0].dy - rockDy) / (vertices[0].dx - rockDx)
    slopeB = (vertices[1].dy - rockDy) / (vertices[1].dx - rockDx)

    interceptA = vertices[0].y - (slopeA * vertices[0].x) 
    interceptB = vertices[1].y - (slopeB * vertices[1].x) 




    xIntercept = round((interceptB - interceptA) / (slopeA - slopeB))
    yIntercept = int(xIntercept * slopeA + interceptA)
    time = round((xIntercept - vertices[0].x) / (vertices[0].dx - rockDx))

    zIntercept = vertices[0].z + (vertices[0].dz - rockDz) * time

    print(xIntercept + yIntercept + zIntercept)

def parse(puzzle : list[str]) -> list[Vertex]:
    verticesList : list[Vertex] = []
    verticesText : list[list[int]] = []
    for line in puzzle:
        numberStr = line.replace('@', ',').split(',')

        args = [int(num) for num in numberStr]
        assert(len(args) == 6)
        verticesText.append(args)

    verticesText.sort()
    verticesList = [Vertex(*args) for args in verticesText]

    return verticesList




def main():
    puzzle = open("input.txt", "r").read().splitlines()
    vertices = parse(puzzle)

    part2(vertices)



if __name__ == "__main__":
    main()
