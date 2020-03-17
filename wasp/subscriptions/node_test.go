package subscriptions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTree_Insert_Remove(t *testing.T) {
	myTree := &tree{
		root: newNode(),
	}
	require.NoError(t, myTree.Insert([]byte("test/a"), 1, "test"))
	require.NoError(t, myTree.Insert([]byte("test/b"), 1, "test2"))
	require.NoError(t, myTree.Insert([]byte("test/b/c"), 1, "test3"))
	require.NotNil(t, myTree.root.children["test"])
	require.NotNil(t, myTree.root.children["test"].children["a"])
	require.NotNil(t, myTree.root.children["test"].children["b"])
	require.Equal(t, "test", myTree.root.children["test"].children["a"].recipients[0])
	require.Equal(t, "test2", myTree.root.children["test"].children["b"].recipients[0])
	require.Equal(t, "test3", myTree.root.children["test"].children["b"].children["c"].recipients[0])
	require.NoError(t, myTree.Remove([]byte("test/b/c"), "test3"))
	require.Nil(t, myTree.root.children["test"].children["b"].children["c"])
}
func TestTree_Match(t *testing.T) {
	myTree := &tree{
		root: newNode(),
	}
	require.NoError(t, myTree.Insert([]byte("test/a"), 1, "test"))
	require.NoError(t, myTree.Insert([]byte("test/+"), 1, "test2"))
	require.NoError(t, myTree.Insert([]byte("test/b/c"), 1, "test3"))
	recipientIds := []string{}
	recipientQos := []int32{}
	require.NoError(t, myTree.Match([]byte("test/b/c"), 0, &recipientIds, &recipientQos))
	require.Equal(t, "test3", recipientIds[0])
	require.NoError(t, myTree.Match([]byte("test/a"), 0, &recipientIds, &recipientQos))
	require.Equal(t, "test", recipientIds[0])
	require.Equal(t, "test2", recipientIds[1])
}
