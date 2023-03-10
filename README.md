# 名可名

一款测名、取名工具，基于中国人姓名的评价标准，包括字音、字形、字意、生肖喜忌、三才五格、八字五行喜用等。

## 八字五行喜用

八字命局分析的核心在于判断日主旺衰，找出用神，借以补救过旺或过衰的命局。

这里我所述用神是一种五行，它基于扶抑原理来平衡日主的旺衰。
我认为用神不是必须在原八字中找的，即使原八字中没有该五行也可取为用神。

找出用神，首先要判断日主是旺还是衰，这也是各家各有方法。 这里我融合主流的得令、得地、得生、得助以及五行生克、地支相冲的理论，建立了一套计算日主旺衰得分的数学模型。
我利用该模型，对所有518400个八字进行计算，算出日主旺衰的得分。日主旺衰的得分可以认为近似地符合期望为282，方差为642的正态分布。
这意味着，八字中日主旺衰的分布是对称的，大概就是一半的八字旺，一半的八字衰；同时注意到方差很大，说明八字之间差异还是比较明显的，有极端旺的八字，也有极端衰的八字。
这样看，这套数学模型基于了主流的八字理论，也得到了尚且合理的数据分布，总体看是可以接受的。

![正态分布拟合图](https://raw.githubusercontent.com/wortin/mkm/master/nchart/bzws.jpg)

我将日主旺衰的情况分为 极衰、过衰、衰、稍衰、平衡、稍旺、旺、过旺、极旺 这9类。

其中日主是极衰、平衡、极旺的八字，都应该是极少数的，应该在所有八字中占比极低的，所以我将这三类八字出现的概率设定为1%。

对于日主旺的八字来说，稍旺、旺、过旺的八字，出现的概率可能符合橄榄球分布，即20%、60%、20%，所以他们在所有八字中出现的概率就是9.7%、29.1%、9.7%。

对于日主衰的八字，同日主旺的八字一样划分概率。

这样，对这9类在正态分布上划分相应的概率区间，就得到它们对应的得分区间为：

+ 极旺 (1776,+
+ 过旺 (1080,1776]
+ 旺 (448,1080]
+ 稍旺 (290,448]
+ 平衡 [274,290]
+ 稍衰 [116,274)
+ 衰 [-516,116)
+ 过衰 [-1212,-516)
+ 极衰 -,-1212)

如果日主平衡，那么无需找用神。

如果日主旺，那么选择克日主、被日主生、被日主克的五行做用神，但究竟选这3个中的哪一个？
我从五行扶抑的理论角度出发，认为八字干支中哪种五行对日主扶助最大，就要在这3个中选出最能抑制该五行的五行为用神。

如果日主衰，同样面临选择生日主还是同日主的五行做用神的问题，也是要先找出八字干支中对日主抑制作用最大的五行，就在这2个中选出最能抑制该五行的五行为用神。

如何找出八字干支中哪种五行对日主扶助或抑制的作用最大？同样建立了一个类似与计算日主旺衰的数学模型即可得到。

## 术语表

姓名 full name (name) 指中国人传统形式的姓名，由姓和名组成。

姓 family name (xing) 指中国人传统形式的姓，由1个汉字或2个汉字组成。

名 first name (ming) 指中国人传统形式的名，由1个汉字或2个汉字组成。

八字 eight characters (bz) 即生辰八字，表示一个人出生的时间，由年柱、月柱、日柱、时柱4部分组成。

四柱 four columns (zhu4) 由年柱、月柱、日柱、时柱这4个柱组成。

年柱 year column (nz) 由年的天干和地支组成。

年干 year heavenly stem (ng) 年的天干。

年支 year earthly branch (nz) 年的地支。

月柱 month column (yz)

日柱 day column (rz)

时柱 hour column (sz)

天干 heavenly stem (tg)

地支 earthly branch (dz)

阴阳 yin and yang (yy)

五行 the five elements (wx)

生 generation (sheng)

克 restriction (ke)

冲 conflict (chong)

合 harmony (he)

化 conversion (hua)

长生十二神 the twelve immortals (zs12s)

天干五合 five-harmonies of heavenly stems (tg5h)

地支六合 six-harmonies of earthly branches (dz6h)

地支半合 semi-harmonies of earthly branches (dzbh)

地支三合 three-harmonies of earthly branches (dz3h)

地支三会 three-meetings of earthly branches (dz3m)

命盘 natal chart (nchart) 指按照人的出生时间和相关理论形成的一种能体现、解释、预测人的性格或命运的信息组织形式

用神 helpful element (yshen) 指对八字可以起到补偏救弊作用的五行中的某一行

断用神 find helpful element (find yshen) 指通过分析八字中日元的旺衰来找出对日元起到或扶或抑作用的五行

十神 ten spirits (shen10) 指根据日干所对应五行，与其他干支所对应五行的生克关系而定的名词

神煞 good spirit and evil spirit (shea)

日主 self (self) 指八字中的日干，代表八字对应的本人，也成命主

发音 pronunciation (fyin) 指基于汉字拼音的发音

声母 initial consonant (shm)

韵母 finals (yum)

复姓 compound surname (fux)

复名 tautonomy (fum)

三才 the three powers (cai3)

五格 the five squares (ge5)

三才五格 the three powers and five squares (c3g5)



## 姓名评价标准









