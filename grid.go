package main

import ( 
"math/rand"
"time"
)

type Grid struct {
 width, height int
 cells [][]string
}

func grid(width, height int) *Grid {
    cells := make([][]string, height)   
    for i:= 0; i<height; i++ {
     cells[i] = make([]string, width)
    }
    return &Grid{width: width, height: height, cells: cells}
}

func initGrid(width, height int,start bool) *Grid{
    grid := grid(width, height)
    cells := grid.cells
    if start {
        for i:= 0; i<height; i++ {
            for j:=0;j<width;j++{
                rand.Seed(time.Now().UnixNano())
                randNum := rand.Float64()
                if randNum < 0.5 {
                    cells[i][j]="*"   
                }else{
                    cells[i][j]=" "
                }
            }   
    }
    }else{
        grid.updateGrid()
    } 

    return grid
}

func (g *Grid) liveNeighbors(row, col int) int {
    directions := [][2]int{
        {-1, -1}, {-1, 0}, {-1, 1},
        {0, -1},         {0, 1},
        {1, -1}, {1, 0}, {1, 1},
    }
    count := 0
    for _, d := range directions {
        newRow, newCol := row+d[0], col+d[1]
        if newRow >= 0 && newRow < g.height && newCol >= 0 && newCol < g.width && g.cells[newRow][newCol] == "*" {
            count++
        }
    }
    return count
}
func (g *Grid) updateGrid() {
    newCells := make([][]string, g.height)
    for i := range newCells {
        newCells[i] = make([]string, g.width)
        for j := range newCells[i] {
            liveNeighbors := g.liveNeighbors(i, j)

            if g.cells[i][j] == "*" {
                if liveNeighbors == 2 || liveNeighbors == 3 {
                    newCells[i][j] = "*"
                } else {
                    newCells[i][j] = " "
                }
            } else {
                if liveNeighbors == 3 {
                    newCells[i][j] = "*"
                } else {
                    newCells[i][j] = " "
                }
            }
        }
    }
    g.cells = newCells
}
