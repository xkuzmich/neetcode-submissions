const numBuckets = 16384

type entry struct {
    key, value int
}
type MyHashMap struct {
	buckets [][]entry
}

func Constructor() MyHashMap {
    return MyHashMap{buckets: make([][]entry, numBuckets)}
}

func (this *MyHashMap) hash(key int) int {
	return key & (numBuckets - 1)
}

func (this *MyHashMap) Put(key int, value int) {
    bucket := this.hash(key)
	b := this.buckets[bucket]
	for i := range b {
		if b[i].key == key {
			b[i].value= value
			return
		}
	}
	this.buckets[bucket] = append(b, entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
    bucket := this.hash(key)
	for _, e := range this.buckets[bucket] {
		if e.key == key {
			return e.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
    bucket := this.hash(key)
   	for i, e := range this.buckets[bucket] {
		if e.key == key {
			b := this.buckets[bucket]
			b[i] = b[len(b)-1]
			this.buckets[bucket] = b[:len(b)-1]
			break
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */