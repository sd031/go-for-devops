package main

import (
	"fmt"
	"strings"
	"time"
)

// Tree struct to hold tree data
type Tree struct {
	height int
	layers []string
}

// Animator interface for animating the tree
type Animator interface {
	Animate(tree *Tree)
}

// TreeAnimator struct to implement Animator interface
type TreeAnimator struct {
	frameCount int
	frameDelay time.Duration
}

// Animate method for TreeAnimator
func (ta *TreeAnimator) Animate(tree *Tree) {
	for frame := 0; frame < ta.frameCount; frame++ {
		fmt.Print("\033[H\033[2J") // Clear the screen (UNIX terminals)
		for _, layer := range tree.layers {
			fmt.Println(layer)
		}
		time.Sleep(ta.frameDelay * time.Millisecond)
		updateTreeLayers(tree, frame)
	}
}

// Function to update tree layers
func updateTreeLayers(tree *Tree, frame int) {
	tree.layers = tree.layers[:tree.height] // Reset to original tree height
	for i := 0; i < tree.height; i++ {
		var layer strings.Builder
		for j := 0; j < tree.height-i-1; j++ {
			layer.WriteString(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			if j%2 == frame%2 {
				layer.WriteString("*")
			} else {
				layer.WriteString("o")
			}
		}
		tree.layers[i] = layer.String()
	}
	// Add the trunk
	trunk := strings.Repeat(" ", tree.height-1) + "| |"
	for i := 0; i < tree.height/2; i++ {
		tree.layers = append(tree.layers, trunk)
	}
}

// Function to create a new tree
func newTree(height int) *Tree {
	tree := &Tree{
		height: height,
		layers: make([]string, height+height/2), // additional space for the trunk
	}
	updateTreeLayers(tree, 0)
	return tree
}

func main() {
	tree := newTree(10)
	animator := &TreeAnimator{frameCount: 10, frameDelay: 200}

	// Using a goroutine and channel for animation
	done := make(chan bool)
	go func() {
		animator.Animate(tree)
		done <- true
	}()

	<-done
}
