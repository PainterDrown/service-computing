# selpg

Golang 实现 Linux 的 selpg 命令，接受 -s, -e, -l 等参数。

## Usage

1. cd $GOPATH

2. go install service-computing/selpg

3. bin/selpg ...

## Arguments

1. bin/selpg -s=1 -e=1 -l=2 input.txt

  注意，需要提前在当前的目录下创建 input.txt 文件并输入两行文本。

2. bin/selpg -s=1 -e=1 -l=2 < input.txt

3. cat input.txt | bin/selpg -s=1 -e=1 -l=2

4. bin/selpg -s=1 -e=1 -l=2 input.txt > output.txt

  注意，需要提前在当前的目录下创建 output.txt 文件。

5. bin/selpg -s=1 -e=-1 input.txt 2> error.txt

  注意，需要提前在当前的目录下创建 error.txt 文件。

6. bin/selpg -s=1 -e=2 -l=200 input.txt > output.txt > error.txt

7. bin/selpg -s=1 -e=2 -l=200 input.txt | cat

8. bin/selpg -s=1 -e=2 -l=200 input.txt > error.txt | cat

9. bin/selpg -s=1 -e=2 -f input.txt

10. bin/selpg -s1=1 -e=2 -l=1 -d=zz input.txt
