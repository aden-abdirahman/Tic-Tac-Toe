package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 600
	screenHeight = 600
	cellSize     = 200
)

// setting our global variables
var (
	gameOver    = false
	draw        = false
	playersMove = true
	aiMove      = false
	x           rl.Texture2D
	o           rl.Texture2D
)

// setting cell struct for our board(t is a texture struct, rect is a rectangle struct, marked is a boolean for monitoring marked cells, and charType is for monitoring if a cell holds an x or o texture)
type Cell struct {
	t        rl.Texture2D
	rect     rl.Rectangle
	marked   bool
	charType string
}

// game over screen
func gameOverFunc() {
	rl.DrawText("GAME OVER", screenWidth/2-rl.MeasureText("GAME OVER", 40)/2, screenHeight/2-rl.MeasureText("GAME OVER", 40)/2+100, 40, rl.Red)
}

// screen for draws
func drawFunc() {
	rl.DrawText("Draw", screenWidth/2-rl.MeasureText("Draw", 40)/2, screenHeight/2-rl.MeasureText("Draw", 40)/2, 40, rl.Red)
}

// quit function that unloads textures from gpu memory and closes the game
func quit() {
	rl.UnloadTexture(x)
	rl.UnloadTexture(o)
	rl.CloseWindow()
}

func main() {
	//	initializing the raylib game window
	rl.InitWindow(screenWidth, screenHeight, "Tic Tac Toe")
	rl.SetMouseScale(1.0, 1.0)

	// loading textures for 'x' and 'o'
	x := rl.LoadTexture("assets/x.png")
	o := rl.LoadTexture("assets/o.png")

	// initializing the game board and setting positions to zero
	var positions int = 0
	board := [3][3]Cell{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j].rect.Width = cellSize
			board[i][j].rect.Height = cellSize

			board[i][j].rect.X = float32(j * cellSize)
			board[i][j].rect.Y = float32(i * cellSize)
		}
	}

	//  setting a var for monitoring if the mouse is clicked
	var mouseButtonPressed bool = false

	//	main game loop
	for !rl.WindowShouldClose() {
		if !gameOver && !draw {

			// getting mouse position
			var mousePos rl.Vector2 = rl.GetMousePosition()

			// if mouse is pressed, check if cell is empty or occupied by 'x' or 'o'
			if mouseButtonPressed {
				for i := 0; i < len(board); i++ {
					for j := 0; j < len(board[i]); j++ {
						// wordy conditionals that basically check if mouse is within the cell and if the cell is empty
						if mousePos.X > board[i][j].rect.X && mousePos.Y > board[i][j].rect.Y && mousePos.X < board[i][j].rect.X+cellSize && mousePos.Y < board[i][j].rect.Y+cellSize && !board[i][j].marked {
							if playersMove {
								board[i][j].t = x
								board[i][j].charType = "x"
							} else {
								board[i][j].t = o
								board[i][j].charType = "o"
							}
							board[i][j].marked = true
							playersMove = !playersMove
							positions++
						}

					}

				}

				mouseButtonPressed = false
			}

			// setting up logic to check if game is over(not the most optimal way but it works :D)
			if board[0][0].marked && board[0][1].marked && board[0][2].marked {
				if board[0][0].charType == board[0][1].charType && board[0][1].charType == board[0][2].charType {
					gameOver = true
				}
			}
			if board[1][0].marked && board[1][1].marked && board[1][2].marked {
				if board[1][0].charType == board[1][1].charType && board[1][1].charType == board[1][2].charType {
					gameOver = true
				}
			}
			if board[2][0].marked && board[2][1].marked && board[2][2].marked {
				if board[2][0].charType == board[2][1].charType && board[2][1].charType == board[2][2].charType {
					gameOver = true
				}
			}
			if board[0][0].marked && board[1][0].marked && board[2][0].marked {
				if board[0][0].charType == board[1][0].charType && board[1][0].charType == board[2][0].charType {
					gameOver = true
				}
			}
			if board[0][1].marked && board[1][1].marked && board[2][1].marked {
				if board[0][1].charType == board[1][1].charType && board[1][1].charType == board[2][1].charType {
					gameOver = true
				}
			}
			if board[0][2].marked && board[1][2].marked && board[2][2].marked {
				if board[0][2].charType == board[1][2].charType && board[1][2].charType == board[2][2].charType {
					gameOver = true
				}
			}
			if board[0][0].marked && board[1][1].marked && board[2][2].marked {
				if board[0][0].charType == board[1][1].charType && board[1][1].charType == board[2][2].charType {
					gameOver = true
				}
			}
			if board[0][2].marked && board[1][1].marked && board[2][0].marked {
				if board[0][2].charType == board[1][1].charType && board[1][1].charType == board[2][0].charType {
					gameOver = true
				}
			}

			if positions == 9 && !gameOver {
				draw = true
			}

		}

		// dealing with mouse events
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsMouseButtonPressed(rl.MouseRightButton) {
			mouseButtonPressed = true
		}

		//  drawing screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// drawing grid on board using lines(used vectors earlier but it was extra code)
		rl.DrawLine(0, screenHeight/3, screenWidth, screenHeight/3, rl.White)
		rl.DrawLine(0, screenHeight*2/3, screenWidth, screenHeight*2/3, rl.White)
		rl.DrawLine(screenWidth/3, 0, screenWidth/3, screenHeight, rl.White)
		rl.DrawLine(screenWidth*2/3, 0, screenWidth*2/3, screenHeight, rl.White)

		// logic for drawing rectangles for each cell and adding textures
		// if !gameOver && !draw { commenting out condition because i want users to be able to see the last move, will update game over screen
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				// this function creates a color filled rectangle for each cell, it takes in the position of the cell, the size of the cell and the color as params
				rl.DrawRectangle(int32(board[i][j].rect.X), int32(board[i][j].rect.Y), int32(board[i][j].rect.Width-5), int32(board[i][j].rect.Height-5), rl.DarkBlue)
				// this function draws a texture on the rectangle, it takes in the position of the cell, the size of the cell and the texture as params. made sure it was centered
				rl.DrawTexture(board[i][j].t, int32(board[i][j].rect.X+board[i][j].rect.Width/2-float32(board[i][j].t.Width)/2), int32(board[i][j].rect.Y+board[i][j].rect.Height/2-float32(board[i][j].t.Height)/2), rl.White)

			}
		}
		// }

		if gameOver {
			gameOverFunc()
		} else if draw {
			drawFunc()
		}

		rl.EndDrawing()

	}

	quit()

}
