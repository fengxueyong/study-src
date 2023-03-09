package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

// 在workspace根目录下：git clone https://github.com/golang/example.git
// 将golang.org/x/example包下载到根目录下，使用workspace模式，自行修改example并依赖。
// example下的修改，将自行被依赖他的其他module所感知，只要这些module都在workspace的use名单里面
// 是一种依赖本地其他module或依赖其他包并需要定制的一种依赖方法，
// module也可以做到。但是需要使用go mod edit -replace .... 然后再go mod tidy。 每个需要依赖本地其他module的module都要这么搞一下，
// 感觉没必要，就大家相互感知就好，想要依赖谁，就把谁添加到go work use的大家庭好了。
// 这个go workspace或go mod edit -replace有点类似IDEA中java项目的工程依赖，如果你直接通过maven依赖（类似于go module不通过replace的方式），
// 每个被依赖都得先打包才可以。
// go.work不需要上传仓库，而go.sum和go.mod需要
// 更多请参考：https://go.dev/ref/mod#workspaces
func main() {
	fmt.Println(stringutil.Reverse("Hello"))
	// stringutil包本地已做修改。
	fmt.Println(stringutil.ToUpper("Hello"))
}
