
## 堆

实现泛型堆，实现后发现除了不用类型断言以外，并不是很好使用


### 堆的基本操作

#### 上滤（percolate up）

上滤一般应用于在一个已经排序好的二叉堆中插入一个新节点。通过上滤，使堆在容纳了新节点后仍能保持原来的堆序。

首先在堆末新建一个空间，称为空穴（hole），然后比较穴的值和其父节点的值。如果空穴的当前位置满足整个堆的堆序，那么就把要插入的值赋给空穴；否则，交换空穴和父节点的位置。迭代实现上面的过程，直到符合整个堆的堆序为止。从宏观上看，空穴会自下而上 地到达满足堆序的位置，称为上滤（percolate up）。


#### 下滤（percolate down）

下滤一般应用于删除了堆顶后的堆序重整过程中。通过下滤，使堆在删除堆顶后把新的堆顶放置在满足堆序的正确的位置上。同时，亦可应用于对一个无序数组的堆排序的算法中。把数组视为一个无序的堆，通过下滤，使堆重新满足一定的堆序。

下滤一般针对堆中的某一个位置进行，传入的参数一般为三个：数组的引用&v、下滤的位置i、数组的大小n（也可以重载为4个参数，添加一个比较器来确定下滤为最大堆还是最小堆。）。假设下滤为最小堆，首先寻找该节点的左右两个孩子中最小的孩子，然后比较该节点和最小的孩子，如果孩子比较小，就交换二者的位置。迭代进行这个过程，直到该节点为叶子节点为止。从宏观上看，节点会在以该节点为根的堆中，自上而下地到达其满足整个堆序的正确位置，称为下滤（percolate down）

在删除堆顶后的堆序重整过程中，堆末的元素会放入堆顶，而其他元素位置不变，从而保证了所有真子堆满足堆序，所以只需对堆顶执行下滤即可。


### 建堆

#### 自顶向下

1. 插入堆，
2. 上滤

#### 自下而上

1. 对每个父节点进行下滤







