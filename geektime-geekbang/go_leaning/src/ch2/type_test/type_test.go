package type_test

import "testing"
// 别名
type MyInt int64
func TestType(t *testing.T){
  var a int32 = 1
  var b int64
  // 显示转化
  b = int64(a)
  var c MyInt
  c = MyInt(b)
  t.Log(a,b,c)
}

func TestPoint(t *testing.T){
  a := 1
  aPtr := &a
  //aPtr = aPtr + 1
  t.Log(a,aPtr)
  t.Logf("%T, %T", a, aPtr)
}

func TestString(t *testing.T){
  var a string
  t.Log("*" + a + "*")

  //
  if a==""{

  }
}

func TestOperator(t *testing.T){
  a :=[...]int{1,2,3,4}
  b:=[...]int{1,4,3,2}
  c:=[...]int{1,2,3,4,5}
  d:=[...]int{1,2,3,4}
  // 返回false
  t.Log(a==b)
  // 有编译错误，因为长度不一样，不能比较
  t.Log(a==c)
  // 返回true
  t.Log(a==d)
}
