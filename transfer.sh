#!/bin/bash

# ./mbc printChain -r 张三

# ./mbc send -a 10 -d "张三转李四10" -f 张三 -r 班长 -t 李四
# ./mbc send -a 20 -d "张三转王五20" -f 张三 -r 班长 -t 王五

# ./mbc send -a 2 -d "王五转李四2" -f 王五 -r 班长 -t 李四
# ./mbc send -a 3 -d "王五转李四3" -f 王五 -r 班长 -t 李四
# ./mbc send -a 5 -d "王五转张三5" -f 王五 -r 班长 -t 张三

# ./mbc send -a 14 -d "李四转赵六14" -f 李四 -r 班长 -t 赵六

./mbc getBalance -r "1HUYzcHwe3WAGUmjpjNNK9MGjH9z9ThC7Y"

./mbc send -a 10 -d "张三转李四10" -f "1HUYzcHwe3WAGUmjpjNNK9MGjH9z9ThC7Y" -r "1397ihxFZ6YHzGcUVnB7Qz9TR78VRMgb5P" -t "1Hpb6ncLcqf2xhwhCkdro23svsKf75b61q"
./mbc send -a 20 -d "张三转王五20" -f "1HUYzcHwe3WAGUmjpjNNK9MGjH9z9ThC7Y" -r "1397ihxFZ6YHzGcUVnB7Qz9TR78VRMgb5P" -t "1HwuPwuxcg6p5U66EMouoNscD7jnJCh6wv"

./mbc send -a 2 -d "王五转李四2" -f "1HwuPwuxcg6p5U66EMouoNscD7jnJCh6wv" -r "1397ihxFZ6YHzGcUVnB7Qz9TR78VRMgb5P" -t "1Hpb6ncLcqf2xhwhCkdro23svsKf75b61q"
./mbc send -a 3 -d "王五转李四3" -f "1HwuPwuxcg6p5U66EMouoNscD7jnJCh6wv" -r "1397ihxFZ6YHzGcUVnB7Qz9TR78VRMgb5P" -t "1Hpb6ncLcqf2xhwhCkdro23svsKf75b61q"
./mbc send -a 5 -d "王五转张三5" -f "1HwuPwuxcg6p5U66EMouoNscD7jnJCh6wv" -r "1397ihxFZ6YHzGcUVnB7Qz9TR78VRMgb5P" -t "1HUYzcHwe3WAGUmjpjNNK9MGjH9z9ThC7Y"

./mbc send -a 14 -d "李四转赵六14" -f "1Hpb6ncLcqf2xhwhCkdro23svsKf75b61q" -r "1397ihxFZ6YHzGcUVnB7Qz9TR78VRMgb5P" -t "1ExZndpNLryEBbaebRn8Qhk4YJYfVjY3tQ"