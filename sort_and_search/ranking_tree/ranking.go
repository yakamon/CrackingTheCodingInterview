package ranking

var root *RankingTreeNode

func Track(n int) {
	if root == nil {
		root = New(n)
	} else {
		root.Insert(n)
	}
}

func GetRankOfNumber(n int) int {
	return root.GetRank(n)
}

type RankingTreeNode struct {
	LeftSize int
	Left     *RankingTreeNode
	Right    *RankingTreeNode
	Data     int
}

func New(data int) *RankingTreeNode {
	return &RankingTreeNode{0, nil, nil, data}
}

func (node *RankingTreeNode) Insert(data int) {
	if data <= node.Data {
		if node.Left != nil {
			node.Left.Insert(data)
		} else {
			node.Left = New(data)
			node.LeftSize++
		}
	} else {
		if node.Right != nil {
			node.Right.Insert(data)
		} else {
			node.Right = New(data)
		}
	}
}

func (node *RankingTreeNode) GetRank(data int) {
	if data == node.Data {
		return node.LeftSize
	}
	if data < node.Data {
		if node.Left == nil {
			return -1
		}
		return node.Left.GetRank(data)
	}
	if node.Right == nil {
		return -1
	}
	rank := node.Right.GetRank(data)
	if rank == -1 {
		return -1
	}
	return node.LeftSize + 1 + rank
}
