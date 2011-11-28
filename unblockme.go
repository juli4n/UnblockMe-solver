package main

import "fmt";
import "strconv";

type Block struct {
	vertical   bool 
	x          int
	y          int
	lenght     int
	special    bool
	number     int
	parentGrid *Grid
}

type Grid struct {
	state [6][6]*Block
}

func (self *Grid) RemoveBlock(b *Block) {
    b.parentGrid = nil;
    self.state[b.x][b.y] = nil;
    var i int;
    for i = 1; i < b.lenght; i++ {
        if b.vertical {
            self.state[b.x][b.y + i] = nil;
        } else {
            self.state[b.x +i][b.y] = nil;   
        }
    }    
}

func (self *Grid) Solved() bool {
    return self.state[5][2] != nil && self.state[5][2].special;
}

func (self *Grid) AddBlock(b *Block) {
    b.parentGrid = self;
    self.AssertEmpty(b.x, b.y);   
    self.state[b.x][b.y] = b;
    var i int;
    for i = 1; i < b.lenght; i++ {
        if b.vertical {
            self.AssertEmpty(b.x, b.y +i);
            self.state[b.x][b.y + i] = b;
        } else {
            self.AssertEmpty(b.x + i, b.y);
            self.state[b.x +i][b.y] = b;   
        }
    }       
}

func (self *Grid) Move(b *Block, clockwise bool) {
    self.RemoveBlock(b);
    b.Move(clockwise);
    self.AddBlock(b);    
}

func (self *Grid) SolveIt(visited map[string]bool, changes int) (solved bool, total_changes int) {

    var doMove = func (clockwise bool, b *Block) (solved bool, total_changes int) {        
        self.Move(b, clockwise);
        if !visited[self.String()] {
            visited[self.String()] = true;
            fmt.Printf("%s\n", self);
            var found, ch = self.SolveIt(visited, changes);
            if found {
                return true, ch + 1;
            } else {
                self.Move(b, !clockwise);
            } 
        } else {
            self.Move(b, !clockwise);
        }
        return false, 0;
    };
    
    if self.Solved() {
        return true, changes;
    }
    
    var i,j int;
    for i = 0; i < 6; i++ {
        for j = 0; j < 6; j++ {
            var b = self.state[i][j];
            if b == nil {
                continue;
            }
            if b.HasMovement(true) {
                var solved, moves = doMove(true, b);
                if solved {
                    return solved, moves;
                }
            }
            if b.HasMovement(false) {
                var solved, moves = doMove(false, b);
                if solved {
                    return solved, moves;
                }
            }          
        }
    }
    return false, 0;
}

func (self *Grid) AssertEmpty(x,y int) bool {
    if self.state[x][y] != nil {
        fmt.Printf("A block already exists [%d,%d]\n", x, y);
        return false;    
    }
    return true;
}

func (self *Block) IsExitPosition() bool {
    if (!self.vertical && self.special) {
        if self.y == 2 && self.x + self.lenght == 6 {
            return true;
        }
    }
    return false;
}

func (self *Block) HasMovement(clockwise bool)  bool {

    if self.vertical {
        if !clockwise && self.y > 0 && self.parentGrid.state[self.x][self.y - 1] == nil {
            return true;
        }
        if clockwise && self.y + self.lenght  < 6 && self.parentGrid.state[self.x][self.y + self.lenght] == nil {
            return true;
        }
        return false;
    } else {
        if !clockwise && self.x > 0 && self.parentGrid.state[self.x -1][self.y] == nil {
            return true;
        }
        if clockwise && self.x  + self.lenght < 6 && self.parentGrid.state[self.x + self.lenght][self.y] == nil {
            return true;
        }
        return false;
    }    
    return false;    
}

func (self *Block) Move(clockwise bool) {
     if self.vertical {
        if clockwise {
            self.y += 1;
        } else {
            self.y -= 1
        }
     } else {
        if clockwise {
            self.x += 1;
        } else {
            self.x -= 1
        }        
     }
}

func (self *Grid) String() string {
	var result = ""
	var i int;
	var j int;
	for i = 0; i < 6; i++ {
	   for j = 0; j < 6; j++ {
	       if self.state[j][i] != nil {
	           if self.state[j][i].special {
                result += "X"    
	           } else {
                result += strconv.Itoa(self.state[j][i].number)
	           }
			} else {
				result += "0"
			}
		}
		result += "\n"
	}
	return result
}

func buildStage1() (grid *Grid) {

    var b1 = &Block{false, 1, 2, 2, true, 1, nil};
    var b2 = &Block{true, 3, 1, 3, false, 2, nil};
    var b3 = &Block{true, 4, 1, 3, false, 3, nil};
    var b4 = &Block{false, 2, 5, 2, false, 4, nil};
    var b5 = &Block{true, 1, 4, 2, false, 5, nil};
    var b6 = &Block{true, 2, 3, 2, false, 6, nil};
    var b7 = &Block{false, 0, 3, 2, false, 7, nil};
    
    grid = new(Grid);
    grid.AddBlock(b1);
    grid.AddBlock(b2);
    grid.AddBlock(b3);
    grid.AddBlock(b4);
    grid.AddBlock(b5);
    grid.AddBlock(b6);
    grid.AddBlock(b7);

    return grid;    
}
/*
func buildStage2() (grid *Grid) {

    var b1 = &Block{false, 0, 0, 3, false, nil};
    var b2 = &Block{true, 2, 1, 2, false, nil};
    var b3 = &Block{true, 5, 0, 3, false, nil};
    var b4 = &Block{false, 3, 2, 2, true, nil};
    var b5 = &Block{false, 0, 3, 3, false, nil};
    var b6 = &Block{true, 2, 4, 2, false, nil};
    var b7 = &Block{true, 3, 3, 3, false, nil};
    var b8 = &Block{false, 4, 5, 2, false, nil};

    grid = new(Grid);
    grid.AddBlock(b1);
    grid.AddBlock(b2);
    grid.AddBlock(b3);
    grid.AddBlock(b4);
    grid.AddBlock(b5);
    grid.AddBlock(b6);
    grid.AddBlock(b7);
    grid.AddBlock(b8);

    return grid;    
}

func buildStage19() (grid *Grid) {

    var b1 = &Block{false, 0, 1, 2, false, nil};
    var b2 = &Block{false, 0, 2, 2, true, nil};
    var b3 = &Block{false, 0, 3, 3, false, nil};
    var b4 = &Block{false, 1, 4, 2, false, nil};
    var b5 = &Block{false, 1, 5, 2, false, nil};

    var b6 = &Block{true, 2, 0, 3, false, nil};
    var b7 = &Block{true, 3, 0, 2, false, nil};
    var b8 = &Block{true, 3, 2, 2, false, nil};

    var b9 = &Block{false, 3, 4, 2, false, nil};
    var b10 = &Block{false, 3, 5, 2, false, nil};

    var b11 = &Block{true, 4, 3, 2, false, nil};
    var b12 = &Block{true, 5, 2, 2, false, nil};
    
    var b13 = &Block{false, 4, 0, 2, false, nil};
    grid = new(Grid);
    grid.AddBlock(b1);
    grid.AddBlock(b2);
    grid.AddBlock(b3);
    grid.AddBlock(b4);
    grid.AddBlock(b5);
    grid.AddBlock(b6);
    grid.AddBlock(b7);
    grid.AddBlock(b8);
    grid.AddBlock(b9);
    grid.AddBlock(b10);
    grid.AddBlock(b11);
    grid.AddBlock(b12);
    grid.AddBlock(b13);

    return grid;    

}
*/
func buildStage400() (grid *Grid) {
    var b1 = &Block{false, 0, 0, 2, false, 1, nil};
    var b2 = &Block{true, 2, 0, 2, false, 2, nil};
    var b3 = &Block{true, 5, 0, 2, false, 3, nil};
    var b4 = &Block{true, 0, 1, 3, false, 4, nil};
    var b5 = &Block{true, 1, 1, 2, false, 5, nil};
    var b6 = &Block{true, 4, 1, 2, false, 6, nil};
    var b7 = &Block{false, 2, 2, 2, true, 7, nil};
    var b8 = &Block{true, 5, 2, 2, false, 8, nil};
    var b9 = &Block{false, 1, 3, 2, false, 9, nil};
    var b10 = &Block{true, 3, 3, 3, false, 1, nil};
    var b11 = &Block{false, 4, 4, 2, false, 2, nil};
    var b12 = &Block{false, 0, 5, 2, false, 3, nil};
    grid = new(Grid);
    grid.AddBlock(b1);
    grid.AddBlock(b2);
    grid.AddBlock(b3);
    grid.AddBlock(b4);
    grid.AddBlock(b5);
    grid.AddBlock(b6);
    grid.AddBlock(b7);
    grid.AddBlock(b8);
    grid.AddBlock(b9);
    grid.AddBlock(b10);
    grid.AddBlock(b11);
    grid.AddBlock(b12);
    return grid;    
}

func main() {

    var stage2 *Grid = buildStage400();
    fmt.Printf("%s\n", stage2);

    var _, changes = stage2.SolveIt(make(map[string]bool), 0);

    fmt.Printf("Solved with %d changes\n", changes);

}
