module github.com/wbing441282413/gotest/case/map

go 1.17

//每个有go.mod文件的就是一个模块，正常来说一个项目应该是一个模块，只有一个go.mod文件
//如果分来了模块，也就是生成了多个go.mod文件的话，想要在模块之间调用（也就是有go.mod文件之间调用的话，就需要各模块有单独的版本号并在引入的时候指明）

// require github.com/wbing441282413/gotest v1.0.3 // indirect
//可以看到没有个给模块单独的设置tag版本号的话，即便我在v1.0.3版本已经删除了最外面的go.mod文件，
//但可以发现，即便是运行go get github.com/wbing441282413/gotest/case/struct/student@v1.0.3 
//download的还是github.com/wbing441282413/gotest@v1.0.3 ，这就会导致这里map的引入也出问题，因为require只到了gotest，而他下面的也有map
//这个map我们是不需要的，因为我们自己模块就是map，这就导致执行的时候还是会出现ambiguous import，不知道引入哪一个的错误


//所以最后还是遵守一个项目一个go.mod的习惯比较好，如果要分多个模块的话
// 要保证的引入的时候有办法可以引入到你想要的模块，而不要直接引入目标模块的父级，不然还是会出现ambiguous import的问题