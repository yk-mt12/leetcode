# 目的: 脳内の思考を書き起こすイメージ
ref. https://github.com/huyfififi/coding-challenges/pull/51/files#diff-71cb86653e9eef6d89d5d3172df491d0c6a39e6f209dda437afa0197ea7bb385

# Step1
linked list
ポインタを使うのはわかる
循環すると言うのはど言うこと？
リストが2以上の場合循環できると言うこと？

Input: head: list<num>[], pos: num

5分経過

https://github.com/Shunii85/arai60/pull/1/changes#diff-78a8cc2f12531874b33e124f4df2f7a1f7303c9f419c2bff1cbcd53c44be047e を参考にした

見てみて、質問の意味を理解していないことがわかった
https://qiita.com/KueharX/items/20f411ebf1c53cff2208をみてみる

連結リストに循環があるかどうかを判断する。そもそも循環とは何か？

> リストにサイクルがあるというのは、次のポインタを辿り続けることで再び到達できるノードが存在する場合を指す。
ref. https://note.com/toppy_taiwan/n/nb238a8a3ae76

つまり, head = [9, 2, 3, 4, 5], pos = 2の場合
head[2]は3なので、5の次は3でサイクルがあると言うことになる。ただしposは関数には渡されない

解法としては2つあり
hashmap, slow/fast方式（Floyd's Algorithm）があるらしい
hashmapはポインタアドレスを保持するhashmapを持っておき、headを順番に見ていく。ポインタアドレスがhashmapに入っていればサイクルがある状態

slow/fast方式の場合、fastが2個先を見る、slowは一個先を見る。

slowがfastを追いかけていって、何週目かでslowとfastが一致するタイミングがある。
もしサイクルがない場合は、fastがループせず先にnilに到達するため、ループせずに終了する

どちらも時間計算量はO(n), 空間計算量はO(1)
Hashmapの場合、訪問したノードを全て記録しておく必要があるため、ノード数に比例してメモリが増える。
一方で、fast & slowは2つの変数のみで一致したらOKなのでメモリ領域を使わないでOK

ちなみにhashmap[1]と言うindex: 1自体のメモリ領域はどれくらいなのか？いまいちメモリという概念を理解していない。
メモリとは？
- OSからRAMの領域を借りて使う
- RAM
- Goの場合
  - int: 8bytes
  - bool: 1byte
  - ポインタ: 8 bytes
  - string "hello": 5byte + 16 bytes
- Hashmapのエントリ一個あたりのメモリ60byte

InputはListNodeのheadだけ。

```go
type ListNode struct {
  Val Int
  Next *ListNode
}

func hashCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

```
# Step2
特になし

# Step3
なんで`for fast != nil && fast.Next != nil {`でないとダメなんだっけ？
終了条件が`fast != nil`だけではダメ？
→ fast.Nextをチェックしないと、fast.Nextがnilの場合panicになる

```go
type ListNode struct {
  Val Int
  Next *ListNode
}

func hasCycle(head *ListNode) bool {
  if head == nil {
    return false
  }

  slow := head
  fast := head.Next

  for fast != nil && fast.Next != nil {
    if slow == fast {
      return true
    }

    slow := slow.Next
    fast := fast.Next.Next
  }
  return false
}
```
