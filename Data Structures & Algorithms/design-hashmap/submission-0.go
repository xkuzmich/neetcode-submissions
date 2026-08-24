const numBuckets = 16384

type MyHashMap struct {
	buckets [][][]int
}

func Constructor() MyHashMap {
    return MyHashMap{buckets: make([][][]int, numBuckets)}
}

func (this *MyHashMap) hash(key int) int {
	return key & (numBuckets - 1)
}

func (this *MyHashMap) Put(key int, value int) {
    bucket := this.hash(key)
	for _, k := range this.buckets[bucket] {
		if k[0] == key {
			k[1] = value
			return
		}
	}
	this.buckets[bucket] = append(this.buckets[bucket], []int{key, value})
}

func (this *MyHashMap) Get(key int) int {
    bucket := this.hash(key)
	for _, k := range this.buckets[bucket] {
		if k[0] == key {
			return k[1]
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
    bucket := this.hash(key)
   	for i, k := range this.buckets[bucket] {
		if k[0] == key {
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