const numBuckets = 1048576

type MyHashSet struct {
	buckets [][]int
}

func Constructor() MyHashSet {
    return MyHashSet{buckets: make([][]int, numBuckets)}
}

func (this *MyHashSet) hash(key int) int {
	return key & (numBuckets - 1)
}
func (this *MyHashSet) Add(key int) {
    bucket := this.hash(key)
	if !this.Contains(key) {
		this.buckets[bucket] = append(this.buckets[bucket], key)
	}
	
}

func (this *MyHashSet) Remove(key int) {
   bucket := this.hash(key)
   	for i, k := range this.buckets[bucket] {
		if k == key {
			b := this.buckets[bucket]
			b[i] = b[len(b)-1]
			this.buckets[bucket] = b[:len(b)-1]
			break
		}
	}
	
}

func (this *MyHashSet) Contains(key int) bool {
	bucket := this.hash(key)
	for _, k := range this.buckets[bucket] {
		if k == key {
			return true
		}
	}
    return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 