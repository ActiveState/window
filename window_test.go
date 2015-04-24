package window

import "testing"
import "container/list"
import "container/ring"

func BenchmarkListS1000(b *testing.B) {
    l(1000,b)
}

func BenchmarkRingS1000(b *testing.B) {
    r(1000,b)
}

func BenchmarkS1000M1(b *testing.B) {
    m(1000,1,b)
}

func BenchmarkS1000M10(b *testing.B) {
    m(1000,10,b)
}

func BenchmarkS1000M100(b *testing.B) {
    m(1000,100,b)
}

func BenchmarkS1000M500(b *testing.B) {
    m(1000,500,b)
}

func BenchmarkSlicifyList(b *testing.B) {
    b.StopTimer()
    l := getList(1000)
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        slicifyList(l)
    }
}

func BenchmarkSlicifyRing(b *testing.B) {
    b.StopTimer()
    l := getRing(1000)
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        slicifyRing(l)
    }
}

func BenchmarkSlicifyWindow(b *testing.B) {
    b.StopTimer()
    w := getWindow(1000)
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        slicifyWindow(w)
    }
}

// m will, given the size and multiple, run
// 1000 times the size worth of data through
// the moving window
func m(size, multiple int, b *testing.B) {
    b.StopTimer()
    w := New(size, multiple)
    TIMES := 1000*size
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        for j := 0; j < TIMES; j++ {
            w.PushBack(string(i))
        }
    }
}

func l(size int, b *testing.B) {
    b.StopTimer()
    lst := list.New()
    TIMES := 1000*size
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        for j := 0; j < TIMES; j++ {
            lst.PushBack(string(j))
            if (lst.Len() > size) {
                lst.Remove(lst.Front())
            }
        }
    }
}

// contributed by: Dustin
func r(size int, b *testing.B) {
    b.StopTimer()
    rng := ring.New(size)
    TIMES := 1000 * size
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        for j := 0; j < TIMES; j++ {
            rng.Value = string(j)
            rng = rng.Prev()
        }
    }
}

func getList(size int) (l *list.List){
    l = list.New()
    for i := 0; i < size; i++ {
        l.PushBack(string(i))
    }
    return
}

func getRing(size int) (r *ring.Ring) {
    r = ring.New(size)
    for i := 0; i < size; i++ {
        r.Value = string(i)
        r = r.Prev()
    }
    return
}

func getWindow(size int) (w *MovingWindow) {
    w = New(size, 1)
    for i := 0; i < size; i++ {
        w.PushBack(string(i))
    }
    return
}

func slicifyList(lst *list.List) {
    s := make([]string, 0, lst.Len())
    for e := lst.Front(); e != nil; e = e.Next() {
        s = append(s, e.Value.(string))
    }
}

func slicifyRing(r *ring.Ring) {
    l := r.Len()
    s := make([]string, 0, l)
    for i := 0; i < l; i++ {
        s = append(s, r.Value.(string))
        r = r.Prev()
    }
}

// not necessary, but to be completely fair
func slicifyWindow(w *MovingWindow) {
    s := w.Slice()
    s[0] = string(0)
}
