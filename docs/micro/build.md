## Go编译注入信息
Go微服务的编译是微服务的第一步，也是比较重要的一个环节。我们可以在编译的时候注入很多编译信息，例如应用名称、应用版本号、框架版本号、编译所在机器、编译时间，我们可以直接注入到二进制里。编译完成后，我们可以使用`./micro --version` ，查看该服务的基本情况，如下图所示。
![image.png](https://cdn.nlark.com/yuque/0/2020/png/497518/1594221263763-b8d01a7f-74fb-4140-ac33-7c98d1d495e4.png#align=left&display=inline&height=182&margin=%5Bobject%20Object%5D&name=image.png&originHeight=156&originWidth=638&size=20314&status=done&style=none&width=746)
我们还可以在微服务启动后，将这些编译信息写入prometheus或者etcd中。当线上出现什么问题的时候，我们能够快速知道微服务在线上使用的哪个版本、编译在什么时间，提升我们排查微服务问题的速度。


接下来我们就来看下如何在Go微服务里编译这些信息
## Go编译含义
我们可以使用指令 `go help build`查看go build的具体用法。
```bash
usage: go build [-o output] [-i] [build flags] [packages]

Build compiles the packages named by the import paths,
along with their dependencies, but it does not install the results.

If the arguments to build are a list of .go files from a single directory,
build treats them as a list of source files specifying a single package.

When compiling packages, build ignores files that end in '_test.go'.

When compiling a single main package, build writes
the resulting executable to an output file named after
the first source file ('go build ed.go rx.go' writes 'ed' or 'ed.exe')
or the source code directory ('go build unix/sam' writes 'sam' or 'sam.exe').
The '.exe' suffix is added when writing a Windows executable.

When compiling multiple packages or a single non-main package,
build compiles the packages but discards the resulting object,
serving only as a check that the packages can be built.

The -o flag forces build to write the resulting executable or object
to the named output file or directory, instead of the default behavior described
in the last two paragraphs. If the named output is a directory that exists,
then any resulting executables will be written to that directory.

The -i flag installs the packages that are dependencies of the target.

The build flags are shared by the build, clean, get, install, list, run,
and test commands:

	-a
		force rebuilding of packages that are already up-to-date.
	-n
		print the commands but do not run them.
	-p n
		the number of programs, such as build commands or
		test binaries, that can be run in parallel.
		The default is the number of CPUs available.
	-race
		enable data race detection.
		Supported only on linux/amd64, freebsd/amd64, darwin/amd64, windows/amd64,
		linux/ppc64le and linux/arm64 (only for 48-bit VMA).
	-msan
		enable interoperation with memory sanitizer.
		Supported only on linux/amd64, linux/arm64
		and only with Clang/LLVM as the host C compiler.
		On linux/arm64, pie build mode will be used.
	-v
		print the names of packages as they are compiled.
	-work
		print the name of the temporary work directory and
		do not delete it when exiting.
	-x
		print the commands.

	-asmflags '[pattern=]arg list'
		arguments to pass on each go tool asm invocation.
	-buildmode mode
		build mode to use. See 'go help buildmode' for more.
	-compiler name
		name of compiler to use, as in runtime.Compiler (gccgo or gc).
	-gccgoflags '[pattern=]arg list'
		arguments to pass on each gccgo compiler/linker invocation.
	-gcflags '[pattern=]arg list'
		arguments to pass on each go tool compile invocation.
	-installsuffix suffix
		a suffix to use in the name of the package installation directory,
		in order to keep output separate from default builds.
		If using the -race flag, the install suffix is automatically set to race
		or, if set explicitly, has _race appended to it. Likewise for the -msan
		flag. Using a -buildmode option that requires non-default compile flags
		has a similar effect.
	-ldflags '[pattern=]arg list'
		arguments to pass on each go tool link invocation.
	-linkshared
		build code that will be linked against shared libraries previously
		created with -buildmode=shared.
	-mod mode
		module download mode to use: readonly, vendor, or mod.
		See 'go help modules' for more.
	-modcacherw
		leave newly-created directories in the module cache read-write
		instead of making them read-only.
	-modfile file
		in module aware mode, read (and possibly write) an alternate go.mod
		file instead of the one in the module root directory. A file named
		"go.mod" must still be present in order to determine the module root
		directory, but it is not accessed. When -modfile is specified, an
		alternate go.sum file is also used: its path is derived from the
		-modfile flag by trimming the ".mod" extension and appending ".sum".
	-pkgdir dir
		install and load all packages from dir instead of the usual locations.
		For example, when building with a non-standard configuration,
		use -pkgdir to keep generated packages in a separate location.
	-tags tag,list
		a comma-separated list of build tags to consider satisfied during the
		build. For more information about build tags, see the description of
		build constraints in the documentation for the go/build package.
		(Earlier versions of Go used a space-separated list, and that form
		is deprecated but still recognized.)
	-trimpath
		remove all file system paths from the resulting executable.
		Instead of absolute file system paths, the recorded file names
		will begin with either "go" (for the standard library),
		or a module path@version (when using modules),
		or a plain import path (when using GOPATH).
	-toolexec 'cmd args'
		a program to use to invoke toolchain programs like vet and asm.
		For example, instead of running asm, the go command will run
		'cmd args /path/to/asm <arguments for asm>'.
```
`go build` 指令比较多。我们把微服务里常用的命令展示在下表：



| 参数 | 备  注 |
| --- | --- |
| -o | 目标地址 |
| -race | 开启竞态检测 |
| -ldflags | 传递参数 |
| -n | 打印编译时会用到的所有命令，但不真正执行 |
| -x | 打印编译时会用到的所有命令 |
| -tag | 根据tag版本编译 |

## -o
编译到指定地址
```bash
go build -o micro
```
## -race
开启竞态检查编译。通过这个编译方式。你的程序可以在运行的时候崩溃
```bash
go build -o micro -race
curl http://127.0.0.1:8080/race
```
我们开启race编译后，访问该地址，就可以看到代码中出现race的报错

## -ldflags

- -w 去掉DWARF调试信息，得到的程序就不能用gdb调试了
-  -s 去掉符号表,panic时候的stack trace就没有任何文件名/行号信息了，这个等价于普通C/C++程序被strip的效果
- -X 设置包中的变量值
```bash
 go build -o micro -ldflags "-X main.buildName=micro\
 -X main.buildGitRevision=f8c315083e7b739f0f055ee46a747c8e109d7539-dirty\
 -X main.buildStatus=Modified -X main.buildUser=`whoami` \
 -X main.buildHost=`hostname -f` -X main.buildTime=`date +%Y-%m-%d--%T`"

```


```go
package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"runtime"
)

var (
	buildName        = "unknown"
	buildGitRevision = "unknown"
	buildUser        = "unknown"
	buildHost        = "unknown"
	buildStatus      = "unknown"
	buildTime        = "unknown"
)

var (
	version bool
	run     bool
)

func init() {
	pflag.BoolVarP(&version, "version", "v", false, `查看版本号`)
	pflag.BoolVarP(&run, "run", "r", false, `运行程序`)
	pflag.Parse()
}

func main() {
	if version == true {
		fmt.Println(LongForm())
	}
	if run == true {
		fmt.Println("go to micro")
	}
}

func LongForm() string {
	return fmt.Sprintf(`Name: %v
GitRevision: %v
User: %v@%v
GolangVersion: %v
BuildStatus: %v
BuildTime: %v
`,
		buildName,
		buildGitRevision,
		buildUser,
		buildHost,
		runtime.Version(),
		buildStatus,
		buildTime)
}
```




![image.png](https://cdn.nlark.com/yuque/0/2020/png/497518/1594134300491-c153add2-9d50-4b3d-8809-db9a80ae7ad9.png#align=left&display=inline&height=156&margin=%5Bobject%20Object%5D&name=image.png&originHeight=156&originWidth=638&size=20314&status=done&style=none&width=638)
## -tag
用于编译打tag，灰度测试代码使用。例如
```bash
go build -o micro -tag="build1"
```
