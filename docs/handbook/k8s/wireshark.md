
TCP Out_of_Order的原因分析：
一般来说是网络拥塞，导致顺序包抵达时间不同，延时太长，或者包丢失，需要重新组合数据单元，因为他们可能是由不同的路径到达你的电脑上面。
TCP Retransmission原因分析：
很明显是上面的超时引发的数据重传。
TCP dup ack XXX#X原因分析：
就是重复应答#前的表示报文到哪个序号丢失，#后面的是表示第几次丢失。
tcp previous segment not captured原因分析
意思就是报文没有捕捉到，出现报文的丢失。


https://blog.csdn.net/chenfengdejuanlian/article/details/53761004