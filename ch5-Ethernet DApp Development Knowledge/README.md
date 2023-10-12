# 以太坊 DApp 开发知识

## 基础知识介绍

1. 什么是区块链
    * 从不同的技术角度来剖析，可以这样来看待区块链
    * 分布式数据库(本质)
        * 每个用户都可以通过合法手段进行读写，不存储于某一两个特定的服务器或安全节点上，而是分布式地存放于网络上所有的完整节点上，每个节点保留一个备份。
    * 网络底层协议(抽象)
        * 它是一种共识协议，基于这种协议，可以在其上开发出各种应用，这些应用在每一时刻都保存一条最长的、最具权威的、共同认可的数据记录，并遵循共同认可的机制进行无需中间权威仲裁的、直接的、点对点的交互信息。
        * TCP/IP: 数据层，网络层，传输层，应用层
        * 区块链: 数据层，网络层，共识层，激励层(Token,通证)，合约层，应用层

2. 区块链特点
    * 去中心化
        * 所有参与其中的网络节点共同维护,无需中心节点调控
    * 不可篡改
        * 整个链条按照时间顺序和哈希指针链接起来,环环相扣
    * 匿名性
        * 私钥和地址是使用网络的所有条件,无需身份验证
    * 可溯源
        * 素有写入区块的数据需要多方验证,公开透明
    * 不可能三角
        * [安全性、去中心化和可扩展性](https://medium.com/@NervosCN/%E5%8C%BA%E5%9D%97%E9%93%BE%E4%B8%8D%E5%8F%AF%E8%83%BD%E4%B8%89%E8%A7%92-%E7%BB%88%E6%9E%81%E6%8C%87%E5%8D%97-85c069f21adc)

3. 比特币和区块链的关系
    * [一文读懂比特币与区块链的关系](https://www.okx.com/cn/learn/relationship-between-bitcoin-and-blockchain-cn)

4. 区块链的发展
    * [什么是区块链技术](https://aws.amazon.com/cn/what-is/blockchain/?aws-products-all.sort-by=item.additionalFields.productNameLowercase&aws-products-all.sort-order=asc)

5. 区块链协议层
    * 网络层（Network Layer）：
        * 介绍: 
            * 网络层负责处理节点之间的通信和连接。它定义了节点之间的传输协议、节点发现和连接的方式，以及数据传输的加密和验证机制。
    * 共识层（Consensus Layer）：
        * 介绍: 
            * 共识层是区块链协议的核心层，决定了如何达成共识并验证交易和区块。它定义了共识算法、区块生成和确认规则，以确保所有节点在区块链上达成一致。
        * 举例: 
            * 工作量证明（Proof of Work，PoW）：
                * 说明
                    * 是比特币和以太坊等早期区块链系统所采用的共识机制。它通过节点通过完成一定的计算任务（即挖矿）来竞争获得记账权，最终获得共识。
                * 特点
                    * 算一道很难的谜题，系统给予挖矿奖励
                    * 多劳多得的社会主义
                * 优点
                    * 所有节点均可参与，记账权公平的分派到每个节点，去中心化
                    * 多劳多得，矿工积极性高
                    * 安全性高，欺诈成本高，如果能够欺诈成功，那么做诚实节点收益更大
                * 缺点
                    * 主流矿池垄断严重，存在51%算力攻击风险
                    * 浪费资源严重 (2018年底消耗全球0.5%电量)
                    * 持币人没有话语权，算力决定一切
                    * 网络性能低，共识时间长
                * 项目
                    * 比特币、以太坊(pow|pos)、比原链等
            * 权益证明（Proof of Stake，PoS）：
                * 说明: 
                    * PoS机制根据节点在系统中拥有的加密货币数量来确定记账权。拥有更多加密货币的节点在共识过程中更有可能被选中，这样可以减少能源消耗和计算资源的浪费。
                * 特点
                    * 不挖矿，依靠币龄也叫币天(币持有数量*持有天数)，币龄越大，获得记账几率越大，利息即为奖励，记账后币龄清零。
                    * 按钱分配，钱生钱的资本主义
                * 优点
                    * 在一定程度上缩短了共识达成的时间
                    * 节约资源
                    * 防作弊，币龄越大，获得记账权几率越大、避免51%攻击， 因为攻击会使自己权益受损
                * 缺点
                    * 数字货币过于集中化，富者越来越富有，散户参与积极性低
                * 项目
                    * 以太坊(pow|pos)、ADA等
            * 权威共识（Delegated Proof of Stake，DPoS）：
                * 说明: 
                    * DPoS是一种基于PoS的共识机制，通过选举一组受信任的代表来验证交易和生成新的区块。这些代表负责维护网络的安全性和一致性。
                * 特点
                    * 不挖矿，每年按比例增发代币，奖励超级节点
                * 优点
                    * 高效、扩展性强
                * 缺点
                    * 节点太少，非去中心化，而是多中心化
                * 项目
                    * EOS等
            * 实用拜占庭容错（Practical Byzantine Fault Tolerance，PBFT）：
                * 说明: 
                    * PBFT是一种拜占庭容错算法，通过节点之间的消息交换和多轮投票来达成共识。它可以容忍一部分节点的错误或恶意行为。
        * [主流区块链共识机制的简介与比较](https://medium.com/@tokenroll/%E4%B8%BB%E6%B5%81%E5%8C%BA%E5%9D%97%E9%93%BE%E5%85%B1%E8%AF%86%E6%9C%BA%E5%88%B6%E7%9A%84%E7%AE%80%E4%BB%8B%E4%B8%8E%E6%AF%94%E8%BE%83-%E5%8C%BA%E5%9D%97%E9%93%BE%E6%8A%80%E6%9C%AF%E5%BC%95%E5%8D%B7%E4%B9%8B%E4%B8%89-e8d1995554b5)
    * 交易层（Transaction Layer）：
        * 介绍: 
            * 交易层定义了区块链上的交易格式和交易验证规则。它负责处理用户发起的交易，并将其打包成区块进行广播和验证。
    * 数据层（Data Layer）：
        * 介绍: 
            * 数据层是区块链存储和管理数据的层。它包括区块的存储和索引、交易的存储和检索，以及其他数据结构的管理，确保数据的安全和完整性。
    * 智能合约层（Smart Contract Layer）：
        * 介绍: 
            * 智能合约层允许开发者编写和部署智能合约，这些合约可以在区块链上执行和管理自动化的业务逻辑。它提供了编程语言、虚拟机和执行环境，使得智能合约可以在区块链上运行。
        * 可编程性：
            * 智能合约可以使用编程语言来定义和实现各种逻辑和功能。这使得开发者可以根据具体需求创建定制化的智能合约。
        * 自动执行：
            * 智能合约在预定条件满足时自动执行，无需第三方介入或人工干预。这提供了高度的可信度和可靠性。
        * 不可篡改性：
            * 智能合约一旦部署在区块链上，就无法修改或删除。这确保了合约规则的不可篡改性和透明性。
        * 去中心化：
            * 智能合约层在整个区块链网络中分布式执行，并由网络中的多个节点验证和维护。这增加了系统的安全性和去中心化程度
    * 用户界面层（User Interface Layer）：
        * 介绍: 
            * 用户界面层是与用户交互的界面，可以是Web应用、移动应用或其他形式的界面。它提供了用户操作和管理区块链的功能，例如创建钱包、发送交易、查询账户余额等。

6. 区块链按应用场景分类
    * 公有链
        * 特点
            * 所有人都可以随时自由的加入和退出，每个节点平等，都有权交易和记账，属于开放式
        * 代表
            * 比特币、以太坊、EOS、NEO、量子链、比原链、并通链
    * 联盟链
        * 特点
            * 仅部分人参与，加入和退出需要授权，选定某些节点为记账人，其他人可以交易，但无记账权，属于半封闭式
        * 代表
            * R3CEV，全球40多个银行成立的联盟组织，2017年7月成立，共享区块链技术
            * IBM farbric项目
            * Linux基金会发起的超级账本 (HyperLedger) 项目,2015年成立,farbric 是子项目
    * 私有链
        * 特点
            * 公司内部使用，可实现更好的权限控制: 管理和审计，属于封闭式
        * 代表
            * 以太坊可以定制自己的私有链、商用区块链链定制

7. 介绍几个概念
    * 分叉
        * 说明: 
            * 代码升级时不同社区意见发生分歧时的结果、重大bug修复是会分叉
        1. 软分叉
            * 旧节点接受新协议产生的区块，毫无感知，新老协议共同维护一条链
        2. 硬分叉
            * 硬分叉是指区块格式或交易格式发生改变时，未升级的节点拒绝验证已经升级的节点生产出的区块，不过已经升级的节点可以验证未升级节点生产出的区块，然后大家各自延续自己认为正确的链，所以分成两条链
            * 旧节点拒绝接收新节点创造的区块，从此分裂为两条独立的链
            * 案例:
                * 以太坊分叉，分为ETC (以太经典)，ETH (以太坊v神)
        3. 叔块
            * 在同一时间出现两个矿工同时挖出矿的情况，此时出现临时的分叉，区块链会同时保留两条链，并等待新生成的区块，新区块选择链接的链就是最长链，即主链，那么另外一个区块就被称为叔块 (以太坊，有奖励)/孤块(比特币，无奖励)
    * BIP和EIP
        * 说明:
            * BIP和EIP是两种不同的提案标准，用于描述和规范比特币（Bitcoin）和以太坊（Ethereum）这两个区块链网络中的改进提案
        1. BIP（Bitcoin Improvement Proposal）：
            * BIP是比特币社区中用于提出和讨论比特币改进的标准化提案。BIP提案由比特币社区成员编写和提交，经过开放的讨论和审查后，最终可以被采纳并应用于比特币协议的升级。BIP提案的目标是通过改进比特币的功能、性能、安全性或协议规范来促进比特币网络的发展和进步
        2. EIP（Ethereum Improvement Proposal）：
            * EIP是以太坊社区中用于提出和讨论以太坊改进的标准化提案。类似于BIP，EIP提案由以太坊社区成员编写和提交，经过开放的讨论和审查后，最终可以被采纳并应用于以太坊协议的升级。EIP提案的目标是改进以太坊的功能、性能、安全性或协议规范，推动以太坊网络的发展和创新

## 关于DApp的一些说明

1. DApp介绍
    * DApp（去中心化应用）是构建在区块链技术之上的应用程序，具有去中心化的特点。与传统的中心化应用程序不同，DApp的数据存储和处理是通过区块链网络上的分布式节点进行的，无需信任单个中心化实体。

2. 特点:
    * 去中心化：
        * DApp的关键特征是去中心化，它不依赖于单一的中央服务器或实体进行操作和管理。相反，DApp使用区块链网络上的多个节点来存储和处理数据，确保透明性和安全性。
    * 自主性：
        * DApp的操作和决策是由智能合约或协议定义的，而不是由中心化实体控制。这使得DApp能够以一种自治的方式运行，无需任何人工干预或信任。
    * 透明性：
        * 由于DApp的数据和交易记录存储在区块链上，所有参与者都可以查看和验证这些数据。这增加了透明度和可信度，减少了潜在的欺诈或篡改风险。
    * 安全性：
        * DApp基于区块链的密码学和共识机制，具有很高的安全性。由于数据存储在多个节点上，并且使用密码学技术进行保护，DApp能够防止数据篡改和未经授权的访问。
    * 去信任化交易：
        * DApp可以实现点对点的交易，无需信任第三方中介。这可以降低交易成本和延迟，并提高交易的效率。

3. DApp应用举例:
    * Cryptokitties：
        * 一个基于以太坊的DApp，允许用户购买、繁殖和交易虚拟猫咪。它在2017年引起了广泛的关注，甚至导致了以太坊网络拥堵。
    * MakerDAO：
        * 一个去中心化的稳定币平台，使用智能合约和抵押来创建和管理稳定价值的数字货币。MakerDAO发行的稳定币DAI在加密货币社区中非常有名。
    * Augur：
        * 一个去中心化的预测市场平台，允许用户创建和交易事件的预测合约。Augur的目标是通过集体智慧来预测事件的结果。
    * Golem：
        * 一个去中心化的计算机网络，允许用户出租自己的计算机资源并获得代币作为报酬。这个平台旨在创建一个全球的计算机资源市场。
    * Brave：
        * 一个基于区块链的隐私浏览器，用户可以选择接受广告并获得代币奖励。Brave的目标是改变互联网广告模式，并提供更好的隐私保护。
    * Metamask: 
        * 一个常用的以太坊钱包和浏览器插件，它允许用户在浏览器中与以太坊区块链进行交互。用户可以使用Metamask来管理以太坊账户、发送和接收以太币

## 以太坊的说明

1. [建议看维基百科](https://zh.wikipedia.org/zh-hans/%E4%BB%A5%E5%A4%AA%E5%9D%8A)