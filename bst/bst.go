package bst

import (
	"fmt"
	"log"
)

type Schedule struct {
	DateHour int
	Venue    string
	Movie    string
}

type NodeList []*Schedule

type Node struct {
	Schedule
	Left  *Node
	Right *Node
	Same  NodeList
}

func (s Schedule) String() string {
	return fmt.Sprintf("%d %s %s", s.DateHour, s.Venue, s.Movie)
}

func (n *Node) String() string {
	return fmt.Sprintf("%d %s %s, %s", n.DateHour, n.Venue, n.Movie, n.Same)
}

func (n *NodeList) ByVenue(v string) *NodeList {
	var res NodeList
	for _, node := range *n {
		if node.Venue == v {
			res = append(res, node)
		}
	}
	return &res
}

func (n *NodeList) ByMovie(m string) *NodeList {
	var res NodeList
	for _, node := range *n {
		if node.Movie == m {
			res = append(res, node)
		}
	}
	return &res
}

type NodeBst struct {
	Root   *Node
	Length int
}

func (d *NodeBst) SearchSingleDateHour(DateHour int) *NodeList {
	var res NodeList
	d.searchHelper(d.Root, DateHour, &res)
	return &res
}

func (d *NodeBst) searchHelper(node *Node, value int, res *NodeList) {

	if node == nil {
		return
	}

	if node.DateHour == value {
		temp := &Schedule{node.DateHour, node.Venue, node.Movie}
		*res = append(*res, temp)
		// fmt.Println(node)
		if len(node.Same) > 0 {
			*res = append(*res, node.Same...)
		}
		return

	} else if node.DateHour < value {
		d.searchHelper(node.Right, value, res)
		return
	} else {
		d.searchHelper(node.Left, value, res)
		return
	}
}

func (d *NodeBst) SearchRangeDate(start, end int) *NodeList {
	var res NodeList
	d.searchRangeHelper(d.Root, start, end, &res)
	return &res
}

func (d *NodeBst) searchRangeHelper(node *Node, start, end int, res *NodeList) {
	if node == nil {
		return
	}

	if node.DateHour >= start && node.DateHour <= end {
		temp := &Schedule{node.DateHour, node.Venue, node.Movie}
		*res = append(*res, temp)
		// fmt.Println(node)
		if len(node.Same) > 0 {
			*res = append(*res, node.Same...)
		}
		d.searchRangeHelper(node.Left, start, end, res)
		d.searchRangeHelper(node.Right, start, end, res)
		return
	} else if node.DateHour <= start {
		d.searchRangeHelper(node.Right, start, end, res)
		return
	} else if node.DateHour > end {
		d.searchRangeHelper(node.Left, start, end, res)
		return
	} else {

	}
}

func (d *NodeBst) AddNode(da int, ve, mo string) {
	if d.Root == nil {
		d.Root = &Node{Schedule{da, ve, mo}, nil, nil, nil}
	} else {
		d.Root = d.addHelper(d.Root, da, ve, mo)
	}
	d.Length += 1
}

func (d *NodeBst) addHelper(node *Node, da int, ve, mo string) *Node {
	if node == nil {
		return &Node{Schedule{da, ve, mo}, nil, nil, nil}
	}
	if da == node.DateHour {
		// fmt.Println(node)
		// check self
		if node.Venue == ve {
			log.Printf("already have movie '%s' in the same DateHour hour %d in venue '%s', adding '%s' to the list failed", node.Movie, node.DateHour, node.Venue, mo)
			return nil
		}

		// check if the same node already exists in the slice
		for _, node := range node.Same {
			if node.Venue == ve {
				log.Printf("already have movie '%s' in the same DateHour hour %d in venue '%s', adding '%s' to the list failed", node.Movie, node.DateHour, node.Venue, mo)
				return nil
			}
		}
		// otherwise append
		node.Same = append(node.Same, &Schedule{da, ve, mo})
	} else if da > node.DateHour {
		node.Right = d.addHelper(node.Right, da, ve, mo)
	} else {
		node.Left = d.addHelper(node.Left, da, ve, mo)
	}
	return node
}

func (d *NodeBst) PrintLevelOrder() {
	// this is for troubleshooting purpose, print the nodes under tree from
	// root level

	var queue []*Node
	var level int
	queue = append(queue, d.Root)

	for len(queue) > 0 {
		length := len(queue)
		fmt.Printf("level %d: \n", level)
		for i := 0; i < length; i++ {
			var popNodeVal Schedule

			popNode := queue[0]
			if popNode != nil {
				popNodeVal = popNode.Schedule
				queue = append(queue, popNode.Left)
				// there is no need to append same, as it's not a bst
				// queue = append(queue, popNode.Same...)
				fmt.Println("current level Same Slice:", popNode.Same)
				queue = append(queue, popNode.Right)
			}
			queue = queue[1:]
			fmt.Printf("current level BST: %s  \n", popNodeVal)
		}
		level++
		fmt.Println("....")
	}

}

// SearchDateHour will find the node with the given value, if not found, return nil
// the return node can be further search by other critieria

// func (d *NodeBst) ModifyByDateHour(DateHour int) *Node {
// 	found := d.SearchDateHour(DateHour)
// 	if found == nil {
// 		log.Println("no such DateHour")
// 		return nil
// 	}
// 	return found
// }

// removeDateHour will remove that node with the given value,
// including all same schedule in the node slice, if not found, return nil
func (d *NodeBst) RemoveDateHour(date int) error {
	n, err := d.removeHelper(d.Root, date)
	d.Root = n
	return err
}

func (d *NodeBst) removeHelper(node *Node, value int) (*Node, error) {
	if node == nil {
		return nil, fmt.Errorf("unable to find node to remove")
	}
	if node.DateHour == value {
		// fmt.Println("found", node)
		// found the node to be removed

		// if node.Left == nil && node.Right == nil {
		// 	fmt.Println("hit")
		// 	return nil, nil
		// }

		// the node to be removed doesn't have left, so we just return right
		if node.Left == nil {
			return node.Right, nil
		}

		if node.Right == nil {
			return node.Left, nil
		}

		// the node to be removed have both left and right
		// we found the largest left
		left := node.Left
		// fmt.Println(left, "test1")
		// fmt.Println(left)
		for left.Right != nil {
			left = left.Right
		}
		// fmt.Println(left, "test2")

		// after the loop, this is the largest node on left (of the node to be removed)
		// we concat the node right to the right side.
		left.Right = node.Right
		node = node.Left

	} else if node.DateHour < value {
		// fmt.Println(value, node)
		n, _ := d.removeHelper(node.Right, value)
		// fmt.Println(value, node, "after")
		node.Right = n

	} else if node.DateHour > value {
		// fmt.Println(value, node, "...")
		n, _ := d.removeHelper(node.Left, value)
		node.Left = n
	}

	return node, nil

}

// func (n *NodeList) SubModifyByVenue(oldVenue, newVenue string) {

// 	// check others
// 	for _, node := range *n {
// 		if node.Venue == oldVenue {
// 			node.Venue = newVenue
// 			log.Printf("Node has been modified frpm %s to %s", oldVenue, newVenue)
// 			return
// 		}
// 	}
// }

// func (n *Node) SubModifyByMovie(oldMovie, newMovie string) {

// 	// check self
// 	if n.Movie == oldMovie {
// 		n.Movie = newMovie
// 		return
// 	}

// 	// check others
// 	for _, node := range n.Same {
// 		if node.Movie == oldMovie {
// 			node.Movie = newMovie
// 			log.Printf("Node has been modified frpm %s to %s", oldMovie, newMovie)
// 			return
// 		}
// 	}
// }
