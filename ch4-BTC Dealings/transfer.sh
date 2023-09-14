#!/bin/bash

./mbc printChain -r 张三

./mbc send -a 10 -d "张三转李四10" -f 张三 -m 班长 -t 李四
./mbc send -a 20 -d "张三转王五20" -f 张三 -m 班长 -t 王五

./mbc send -a 2 -d "王五转李四2" -f 王五 -m 班长 -t 李四
./mbc send -a 3 -d "王五转李四3" -f 王五 -m 班长 -t 李四
./mbc send -a 5 -d "王五转张三5" -f 王五 -m 班长 -t 张三

./mbc send -a 14 -d "李四转赵六14" -f 李四 -m 班长 -t 赵六