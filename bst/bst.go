package bst

import (
	"errors"
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

func (n *NodeList) ModifyMovieOrVenue(newMovie, newVenue string) error {

	// this support only change movie or venue, date hour will be the same

	if len(*n) == 0 {
		log.Println("no item to modify")
		return errors.New("no item to modify")
	}
	if len(*n) > 2 {
		log.Println("too many items to modify")
		return errors.New("too many items to modify")
	}

	// no change on datehour, which is easier
	(*n)[0].Movie = newMovie
	(*n)[0].Venue = newVenue

	return nil
}

type Bst struct {
	Root   *Node
	Length int
}

func (d *Bst) Test() *NodeList {
	return &d.Root.Same
}

func (d *Bst) SearchSingleDateHour(DateHour int) *NodeList {
	var res NodeList
	d.searchHelper(d.Root, DateHour, &res)
	return &res
}

func (d *Bst) searchHelper(node *Node, value int, res *NodeList) {
	if node == nil {
		return
	}

	if node.DateHour == value {
		*res = node.Same
		// temp := &Schedule{node.DateHour, node.Venue, node.Movie}
		// *res = append(*res, temp)
		// // fmt.Println(node)
		// if len(node.Same) > 0 {
		// 	*res = append(*res, node.Same...)
		// }
		return

	} else if node.DateHour < value {
		d.searchHelper(node.Right, value, res)
		return
	} else {
		d.searchHelper(node.Left, value, res)
		return
	}
}

func (d *Bst) SearchRangeDateHour(start, end int) *NodeList {
	var res NodeList
	d.searchRangeHelper(d.Root, start, end, &res)
	return &res
}

func (d *Bst) searchRangeHelper(node *Node, start, end int, res *NodeList) {
	if node == nil {
		return
	}

	if node.DateHour >= start && node.DateHour <= end {
		// fmt.Println(node, "hit")
		// for _, node := range node.Same {
		// 	*res = append(*res, node)
		// }
		// temp := &Schedule{node.DateHour, node.Venue, node.Movie}
		// *res = append(*res, temp)
		// fmt.Println(node)
		// if len(node.Same) > 0 {
		*res = append(*res, node.Same...)
		// }
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

func (d *Bst) AddNode(da int, ve, mo string) error {
	if d.Root == nil {
		temp := Schedule{da, ve, mo}
		tempSlice := NodeList{&temp}
		d.Root = &Node{temp, nil, nil, tempSlice}
		d.Length += 1
		return nil
	} else {
		tmp, err := d.addHelper(d.Root, da, ve, mo)
		if err != nil {
			return err
		}
		d.Root = tmp
		d.Length += 1
		return nil
	}

}

func (d *Bst) addHelper(node *Node, da int, ve, mo string) (*Node, error) {
	if node == nil {
		// any node with same date hour will be added to same slice
		// including the first node
		temp := Schedule{da, ve, mo}
		tempSlice := NodeList{&temp}
		return &Node{temp, nil, nil, tempSlice}, nil
	}
	if da == node.DateHour {
		// fmt.Println(node)

		// check if the same node already exists in the slice
		for _, node := range node.Same {
			if node.Venue == ve {
				log.Printf("already have movie '%s' in the same DateHour hour %d in venue '%s', adding '%s' to the list failed", node.Movie, node.DateHour, node.Venue, mo)
				return nil, errors.New("already have movie in the same DateHour hour in venue")
			}
		}
		// otherwise append
		node.Same = append(node.Same, &Schedule{da, ve, mo})
	} else if da > node.DateHour {
		tmp, err := d.addHelper(node.Right, da, ve, mo)
		if err != nil {
			return nil, err
		}
		node.Right = tmp

	} else {
		tmp, err := d.addHelper(node.Left, da, ve, mo)
		if err != nil {
			return nil, err
		}
		node.Left = tmp
	}
	return node, nil
}

func (d *Bst) ModifyDateHour(oldDate, newDate int,
	oldMovie, newMovie string,
	oldVenue, newVenue string) error {

	if oldDate != newDate {
		// this means we need to remove and insert as the date has changed
		err := d.RemoveOneEntry(oldDate, oldVenue, oldMovie)

		if err != nil {
			return err
		}
		d.AddNode(newDate, newVenue, newMovie)
		return nil
	} else {
		return errors.New("you must change the date")
	}
}

func (d *Bst) PrintLevelOrder() {
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

// removeDateHour will remove that node with the given value,
// including all same schedule in the node slice, if not found, return nil
func (d *Bst) RemoveDateHour(date int) error {
	n, err := d.removeHelper(d.Root, date)
	d.Root = n
	return err
}

func (d *Bst) removeHelper(node *Node, value int) (*Node, error) {
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

func (d *Bst) RemoveOneEntry(date int, venue, movie string) error {
	// tmp := d.Root
	temp, err := d.removeOneEntryHelper(d.Root, date, movie, venue)
	if err != nil {
		return err
	} else {
		d.Root = temp
		return nil
	}
}

func (d *Bst) removeOneEntryHelper(node *Node, date int, venue, movie string) (*Node, error) {

	if node == nil {
		return nil, fmt.Errorf("unable to find node to remove")
	}
	if node.DateHour == date {
		var tempNode *Node
		for i := range node.Same {
			// fmt.Println("hit", node.Same[i], date, movie, venue, i)
			if node.Same[i].Movie == movie && node.Same[i].Venue == venue {

				//check if the node to be removed is the first node in the slice
				if i == 0 {

					// check if the same slice has other nodes or not
					if len(node.Same) == 1 {
						// no other nodes in the same slice, we follow bst removal
						if node.Left == nil {
							return node.Right, nil
						}

						if node.Right == nil {
							return node.Left, nil
						}

						left := node.Left
						for left.Right != nil {
							left = left.Right
						}
						left.Right = node.Right
						node = node.Left
						return node, nil
					} else {
						// there is some nodes left in the same slice
						// we pop it to be come the new node
						nextNodeAfterRemove := node.Same[i+1]
						tempSchedule := Schedule{
							nextNodeAfterRemove.DateHour,
							nextNodeAfterRemove.Venue,
							nextNodeAfterRemove.Movie,
						}
						tempNode = &Node{tempSchedule, node.Left, node.Right, node.Same[1:]}
						tempNode.Left = node.Left
						tempNode.Right = node.Right
						// fmt.Println(tempNode)
						return tempNode, nil
					}

				} else {
					// the matching node is not the first node in the slice
					// we need to remove the node from the slice
					// bst level no change
					node.Same = append(node.Same[:i], node.Same[i+1:]...)
					return node, nil
				}

			}

		}

		// reached to the end of same slice, and no matching node found
		return node, fmt.Errorf("unable to find node in slice to remove")

	} else if node.DateHour < date {
		n, err := d.removeOneEntryHelper(node.Right, date, movie, venue)
		if err != nil {
			return nil, err
		}
		node.Right = n
	} else {
		n, err := d.removeOneEntryHelper(node.Left, date, movie, venue)
		if err != nil {
			return nil, err
		}
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
