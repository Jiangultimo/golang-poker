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
