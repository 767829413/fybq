# BTC交易

## 转账方式

1. 传统银行

    * **传统银行的每个账户都会有一个数据库表来存储用户的信息，包括姓名，卡号，余额等基本信息，每产生一笔交易后，最终都会更新这个余额字段，这个数据表就是这个账户的存储结构。**

    * ![1](https://pic.imgdb.cn/item/64fa98ec661c6c8e5491ec97.png)

    * 比特币的数据库中只有交易，没有这个用来集中保存用户基本信息的数据表，也就是没有地方存储账户，没有地方存储余额，那么比特币系统如何维护我们的钱呢?

2. 比特币转账

    * **每个交易要以以前的交易为基础，而不是以余额为基础，从而确定有足够的金额发起新的交易**

    * Lily -> Tom, Tom -> Divid, Divid -> John
        * ![2](https://pic.imgdb.cn/item/64fa9b10661c6c8e54925c11.png)

    * 钱零散的分布在不同交易中
  
3. 比特币找零机制
    * 情况一: 
        * 上面最初Lily的钱来自挖矿,Tom的钱来自于Lily转账
            * 那么假设Lily有200,转给Tom的钱是100,那么如何扣除Lily的钱呢?
            * 能修改上笔的交易中的数据吗?
            * 比特币是防止篡改的,那怎么办呢?
        * 解决方案:
            * ![3](https://pic.imgdb.cn/item/64faa778661c6c8e5494d816.png)
            * 把上笔交易中Lily的钱全部花掉,扣除转给Tom的钱,余额再转给Lily,最终就是剩余的钱自己转给自己
            * 类似于生活中的找零机制
            
    * 情况二:
        * Tom的钱来自Lily,Jim也可以转账给Tom,这样Tom有两笔钱
            * ![4](https://pic.imgdb.cn/item/64faa8e9661c6c8e549501f7.png)
            * 当Tom给Divid进行转账的时候,可能需要两笔钱加起来才能凑齐
        * 解决方案:
            * Tom对Divid的转账就要同时花掉这两笔钱，如果有余额，Tom同样要转回给自己.

4. 总结
    * 每一笔能够支配的钱来自上一次交易的输出(普通交易)
    * 每一笔花费的输出都要一次性花完,如果有剩余再转给自己

## 比特币交易形式

* 比特币中没有付款人和收款人,只有输入(input)和输出(output),每个输入都对应之前别人给你转账时产生的某个输出

1. 普通交易(类似找零)
    * ![5](https://pic.imgdb.cn/item/64faabe0661c6c8e54958d14.png)

2. 多对一(类似凑零钱付账)
    * ![6](https://pic.imgdb.cn/item/64faac15661c6c8e54959274.png)

3. 一对多(类似代发工资)
    * ![7](https://pic.imgdb.cn/item/64faac4a661c6c8e5495986b.png)

4. 多对多(类似大额支付+找零钱)
    * ![8](https://pic.imgdb.cn/item/64faac8c661c6c8e5495b081.png)

## 比特币交易结构

1. 如何同时转账并且找零
    * 一笔交易中可以有多个输入和多个输出,给自己找零就是给自己生成一个输出
    * 上面的Tom给Divid转帐中,对于output而言,Divid和Tom具有完全相同的地位
    * 每一笔交易的输入都来源于上一个交易的输出
    * ![9](https://pic.imgdb.cn/item/64faae87661c6c8e5496c24b.png)

2. 交易输出如何产生
    * 输出产生流程
        * Lily转账给Tom,比特币系统会生成一个output
        * output包含两个东西
            * Lily给Tom转账的金额,比如1BTC
            * 一个锁定脚本,使用Tom的公钥哈希对转账金额1BTC进行锁定
        * 注意:
            * 是公钥哈希不是Tom的地址,地址可以推出来公钥哈希
            * 不用关心锁定脚本是什么,理解为这个钱被Tom的公钥加密了,只有Tom能解开支配
    * 真实的锁定脚本
        * [参考](https://zhuanlan.zhihu.com/p/33157713)
        * P2PKH的举例
            * 在一笔交易中，Lily给Tom支付了0.15BTC
            * 由于在比特币中并没有账户的概念，这一笔交易的输出并没有写上Tom的名字，也没有写上Tom的公钥，而是写上了Tom公钥的哈希值。
                * 进一步保证了用户的隐私。
                * Tom想要花费这0.15个BTC，应该如何证明自己拥有这个UTXO，并且，其他人无法假冒Tom来花费这个UTXO呢？
                    * 答案是比特币的交易创建的输出其实并非一个简单的公钥地址，而是一个脚本。Lily给Tom支付0.15个BTC的这个交易中，Lily创建的输出脚本类似:
                        * `OP_DUP OP_HASH160 <Tom Public Key Hash> OP_EQUAL OP_CHECKSIG`
                        * 谁能够提供一个签名和一个公钥，让这个脚本运行通过，谁就能花费这笔交易的0.15个BTC。由于创建签名只能使用Tom的私钥，非Tom的私钥创建的签名将无法通过这个脚本的验证，所以，其他人无法假冒Tom来花费这笔交易

3. 交易输入如何产生
    * 输入产生流程
        * 与output对应的是input结构,每一个input都源自一个output,在Tom对Divid进行转账时,系统会创建input
        * 为了定位这笔钱的来源,input结构包含以下内容:
            * 在哪一笔交易中，即需要Lily->Tom这笔转账的交易ID(hash)
            * 所引用交易的那个output，所以需要一个output的索引(int)
            * 定位到了这个output，如何证明能支配呢，所以需要一个Tom的签名。 (解锁脚本，包括签名和自己的公钥)
        * 注意:
            * 不用关心这个解锁脚本原理，只需记得这个能解开用我公钥加密的比特币即可
            * .挖矿奖励没有输入
        * 说明:
            * ![10](https://pic.imgdb.cn/item/64fab60e661c6c8e54981f20.png)
            * 上述是input的全部内容，那么由于Tom引用了两个output，所以他这笔交易中包含两个input
            * 上面的场景中，Tom引用Lily转给他的1btc以及Jim转给他的2btc，完成对Divid转账2.5btc，找零0.5btc给自己
            * 这笔交易中，有两个input和两个output。
    * 真实的解锁脚本
        * [参考](https://zhuanlan.zhihu.com/p/33157713)
        * P2PKH的举例
            * 对应上面的输出流程
            * `<Tom Signature> <Tom Public Key>`
            * 只有当解锁版脚本与锁定版脚本的设定条件相匹配时，执行组合有效脚本时才会显示结果为真（Ture）
            * 只有当解锁脚本得到了Tom的有效签名，交易执行结果才会被通过（结果为真），该有效签名是从与公钥哈希相匹配的Tom的私钥中所获取的

4. 完整脚本流程
    * [参考](https://zhuanlan.zhihu.com/p/33157713)
    * ![11](https://pic.imgdb.cn/item/64fab823661c6c8e5498c902.png)
        * 矿工接收到交易之后会进行校验:
            1. 将输入的解锁脚本解析出来,放到栈里(内存中)
            2. 将输出的锁定脚本解析出来,放到栈里(内存中)
    * [示例](https://explorer.btc.com/btc/transaction/4c3c177d08ec79e292aa169e14355a81924a5d995e18e686d87ebacc07dce252)

5. 未消费输出(UTXO)

    ```text
    我们看到当Tom给Divid转账的时候，已经花掉了Lily和Jim转给他的钱，当完成给Divid的转账后，他还有找零得到的0.5btc，那这里就会涉及到output的消费问题，我们把尚未使用的output有个专用的名字，叫做未消费输出(unspent transaction outputUTXO)
    ```

    * 总结:
        1. UTX0:unspent transaction output，是比特币交易中最小的支付单元，不可分割，每一个UTXO必须一次性消耗完，然后生成新的UTXO，存放在比特币网络的UTXO池中。
        2. UTXO是不能再分割、被所有者锁住或记录于区块链中的并被整个网络识别成货币单位的一定量的比特币货币
        3. 比特币网络监测着以百万为单位的所有可用的(未花费的)UTXO。当一个用户接收比特币时，金额被当作UTXO记录到区块链里。这样，一个用户的比特币会被当作UTXO分散到数百个交易和数百个区块中。
        4. 实际上，并不存在储存比特币地址或账户余额的地点，只有被所有者锁住的、分散的UTXO。
        5. “一个用户的比特币余额”，这个概念是一个通过比特币钱包应用创建的派生之物。比特币钱包通过扫描区块链并聚合所有属于该用户的UTXO来计算该用户的余额。
        6. UTXO被每一个全节点比特币客户端在一个储存于内存中的数据库所追踪，该数据库也被称为“UTXO集"或者“UTXO池”。新的交易从UTXO集中消耗(支付)一个或多个输出。

6. 交易结构

    | 字段 | 大小 | 说明 |
    | :-- | :-- | :-- |
    | 4字节 | 版本 | 明确这笔交易参照的规则 |
    | 1-9字节 | 输入数量 | 被包含的输入的数量 |
    | 不定 | 输入 | 一个或多个交易输入 |
    | 1-9字节 | 输出数量 | 被包含的输出的数量 |
    | 不定 | 输出 | 一个或多个交易输出 |
    | 4字节 | 时钟时间 | 个UNIX时间戳或区块号 |

    * 交易输入(TXInput)
        * 指明交易发起人可支付资金的来源:
            * 引用的utxo所在交易id
            * 所消费的utxo在output中的索引
            * 解锁脚本
                * 签名,公钥
    * 交易输出(TXOutput)
        * 包含资金接收方的相关信息:
            * 接收金额
            * 锁定脚本
                * 对方公钥哈希
                    * 可以通过地址来运算出来,转账只要知道地址即可
    * 交易ID
        * 一般是交易结构的哈希
            * 参考block哈希
    
7. UTXO模拟流程图
    * ![1.png](https://s2.loli.net/2023/09/11/KNih5e9gcGYu7IM.png)