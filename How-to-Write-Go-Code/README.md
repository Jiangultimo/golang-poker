### 工作空间
- Go代码必须放在工作空间内，该目录包含三个子目录：
  1. `src`目录包含Go的源文件，他们被组织成包(每个目录都对应一个包)
  2. `pkg`目录包含对象
  3. `bin`目录包含可执行命令
  
### GOPATH环境变量
`GOPATH`环境变量制定了工作空间位置。首先创建一个工作空间目录，并设置相应的GOPATH
```bash
mkdir $HOME/work
export GOPATH=$HOME/work
```
作为约定，将工作空间的`bin`子目录添加到你的PATH中:
```bash
export PATH=$PATH:$GOPATH/bin
```

### 包路径

### 包名
```bash
package 名称
```
- 一个包中的所有文件都必须使用使用想用的`名称`
- GO约定包名为导入路径的最后一个元素：作为"crypto/rot13"导入的包应命名为rot13

### 测试
Go拥有一个轻量级的测试框架，它由go test命令和testing包构成。

### 远程包
- 若在包的导入路径中包含了代码仓库的URL，go get 就会自动地获取、构建并安装：
  ```bash
  $ go get github.com/golang/example/hello
  $ $GOPATH/bin/hello
  Hello, Go examples!
  ```
- 若指定的包不在工作空间中，`go get`就会将它放到GOPATH指定的第一个工作间内。（若该包已经存在，`go get`就会跳过远程获取，其行为与`go install`相同）