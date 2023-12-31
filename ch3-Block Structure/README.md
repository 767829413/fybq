# 区块结构

| 字节 | 字段 | 说明 |
| :-- | :-- | :-- |
| 4 | 区块大小 | 用来表示该字段之后的区块大小 |
| 80 | 区块头 | 组成区块头的几个字段 |
| 1~9 | 交易计数器 | 该区块包含的交易数量,包含coinbase交易 |
| 动态 | 交易 | 记录在区块里的交易信息,使用原生交易信息格式,并且交易在数据流中的位置必须与Merkle(墨克)树的叶子节点顺序一致 |

* 注意: 
  * 比特币的区块大小被严格限制在1MB以内,4字节的区块大小字段不包含在此内

## 区块头

| 字节 | 字段 | 说明 |
| :-- | :-- | :-- |
| 4 | 版本 | 区块版本号,表示本区块遵守的验证规则 |
| 32 | 父区块头哈希值 | 前一区块的哈希值,使用SHA256(父区块头)计算 |
| 32 | Merkle(墨克)根 | 该区块中交易的Merkle树(墨克树)根的哈希值,同样采用SHA256()计算 |
| 4 | 时间值 | 该区块产生的近似时间,精确到秒的unix时间戳,必须严格大于前11个区块时间的中值,同时全节点也会拒绝那些超过自己2小时时间戳的区块 |
| 4 | 难度目标(难度值) | 该区块工作量证明算法的难度目标,已使用特定的算法编码 |
| 4 | Nonce | 为了找到满足难度目标所设定的随机数,为了解决32位随机数在算力飞升的的情况下不够用的问题,规定时间戳和coinbase交易信息均可更改,用来扩展Nonce位数 |

* 示例: 
  * ![1.png](https://s2.loli.net/2023/09/05/autd6fCN7mMFRjg.png)
* Merkle(墨克)树
  * [哈系树](https://zh.wikipedia.org/wiki/%E5%93%88%E5%B8%8C%E6%A0%91)
* 注意:
  * 区块不存储hash值,节点接收区块后独立计算并存储在本地.
  * ![2.png](https://s2.loli.net/2023/09/05/9CHWcPSh7pIQLdt.png)

## 区块体

* 示例:
  * ![3.png](https://s2.loli.net/2023/09/05/lA9aiCZYMqdTUg8.png)

* Coinbase交易
  * 第一条交易,挖矿奖励矿工

* 普通转账交易
  * 每笔交易包含付款方,收款方,付款金额,收款金额,手续费等

## 完整区块示意图

* ![4.png](https://s2.loli.net/2023/09/05/NxU4XspnSuMjZVa.png)

* 创世块: Block #0
* ![5.png](https://s2.loli.net/2023/09/05/bJ2anTgtLusYIwW.png)
        